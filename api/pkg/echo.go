package pkg

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	DIST_DIR     = "frontend/dist"
	FAVICON_FILE = "favicon.ico"
)

//go:embed all:frontend/dist/*
var publicDir embed.FS

func NewEcho(baseURL string, cors string) *echo.Echo {
	e := echo.New()

	e.HTTPErrorHandler = HTTPErrorHandler
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: ltsv(),
	}))
	SetupRoutes(e, baseURL)

	if cors != "" {
		SetupCors(e, cors)
	}

	return e
}

func SetupRoutes(e *echo.Echo, baseURL string) {
	e.GET(baseURL+"", NewAssetsHandler(publicDir, "index.html").GetHTML)
	e.GET(baseURL+"robots.txt", NewAssetsHandler(publicDir, "robots.txt").GetPlain)
	e.GET(baseURL+"ads.txt", NewAssetsHandler(publicDir, "ads.txt").GetPlain)

	e.GET(baseURL+FAVICON_FILE, NewAssetsHandler(publicDir, "favicon.ico").GetICO)
	e.GET(baseURL+"api", NewAPIHandler().Get)
}

func SetupCors(e *echo.Echo, cors string) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:" + cors},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
}

func StartEcho(e *echo.Echo, host string, port string) {
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", host, port)))
}

// HTTPErrorResponse is the response for HTTP errors
type HTTPErrorResponse struct {
	Error interface{} `json:"error"`
}

// HTTPErrorHandler handles HTTP errors for entire application
func HTTPErrorHandler(err error, c echo.Context) {
	SetHeadersResponseJSON(c.Response().Header())
	code := http.StatusInternalServerError
	var message interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	} else {
		message = err.Error()
	}

	if code == http.StatusInternalServerError {
		message = "Internal Server Error"
	}
	if err = c.JSON(code, &HTTPErrorResponse{Error: message}); err != nil {
		slog.Error(err.Error())
	}
}

func ltsv() string {
	timeCustom := time.Now().Format("2006-01-02 15:04:05")
	var format string
	format += fmt.Sprintf("time:%s\t", timeCustom)
	format += "host:${remote_ip}\t"
	format += "forwardedfor:${header:x-forwarded-for}\t"
	format += "req:-\t"
	format += "status:${status}\t"
	format += "method:${method}\t"
	format += "uri:${uri}\t"
	format += "size:${bytes_out}\t"
	format += "referer:${referer}\t"
	format += "ua:${user_agent}\t"
	format += "reqtime_ns:${latency}\t"
	format += "cache:-\t"
	format += "runtime:-\t"
	format += "apptime:-\t"
	format += "vhost:${host}\t"
	format += "reqtime_human:${latency_human}\t"
	format += "x-request-id:${id}\t"
	format += "host:${host}\n"
	return format
}
