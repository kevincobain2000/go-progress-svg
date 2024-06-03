package pkg

import (
	"net/http"

	gps "github.com/kevincobain2000/go-progress-svg"
	"github.com/labstack/echo/v4"
	"github.com/mcuadros/go-defaults"
)

type APIHandler struct {
}

func NewAPIHandler() *APIHandler {
	return &APIHandler{}
}

type APIRequest struct {
	Style string `json:"style"  query:"style" default:"circle" validate:"required,ascii,oneof=circle battery bar" message:"style is required"`

	// common on all styles
	Progress int `json:"progress" query:"progress" default:"1" validate:"required,min=0,max=100" message:"progress is required"`

	// only circle style
	CircleSize    int    `json:"size" query:"size" default:"100" validate:"required,min=50,max=500" message:"size is required"`
	CircleWidth   int    `json:"circle_width" query:"circle_width" default:"15" validate:"required,min=1,max=200" message:"circle_width is required"`
	ProgressWidth int    `json:"progress_width" query:"progress_width" default:"15" validate:"required,min=1,max=50" message:"progress_width is required"`
	CircleColor   string `json:"circle_color" query:"circle_color" default:"e0e0e0" validate:"required" message:"circle_color is required"`

	// circle style
	ProgressColor string `json:"progress_color" query:"progress_color" default:"76e5b1" validate:"required" message:"progress_color is required"`
	TextColor     string `json:"text_color" query:"text_color" default:"6bdba7" validate:"required" message:"text_color is required"`
	TextSize      int    `json:"text_size" query:"text_size" default:"52" validate:"required,min=10,max=100" message:"text_size is required"`

	BackgroundColor string `json:"background_color" query:"background_color" default:"e0e0e0" validate:"required" message:"background_color is required"`
	CaptionSize     int    `json:"caption_size" query:"caption_size" default:"25" validate:"required,min=2,max=100" message:"caption_size is required"`
	CaptionColor    string `json:"caption_color" query:"caption_color" default:"000000" validate:"required" message:"caption_color is required"`
	SegmentCount    int    `json:"segment_count" query:"segment_count" default:"1" validate:"min=0,max=100" message:"segment_count min 0 and max 100"`
	SegmentGap      int    `json:"segment_gap" query:"segment_gap" default:"0" validate:"min=0,max=10" message:"segment_gap min 0 and max 10"`
	Caption         string `json:"caption" query:"caption" default:"" validate:"ascii,min=0,max=20" message:"caption min 0 and max 20"`
	CornerRadius    int    `json:"corner_radius" query:"corner_radius" default:"10" validate:"required,min=1,max=50" message:"corner_radius is required"`

	// only bar style
	Width  int `json:"width" query:"width" default:"200" validate:"required,min=50,max=500" message:"width is required"`
	Height int `json:"height" query:"height" default:"20" validate:"required,min=20,max=500" message:"height is required"`
}

func (h *APIHandler) Get(c echo.Context) error {
	req := new(APIRequest)
	if err := BindRequest(c, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	defaults.SetDefaults(req)
	msgs, err := ValidateRequest(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	if req.Style == "circle" {
		circular, err := gps.NewCircular(func(o *gps.CircularOptions) error {
			o.Progress = req.Progress
			o.CircleSize = req.CircleSize
			o.CircleWidth = req.CircleWidth
			o.ProgressWidth = req.ProgressWidth
			o.CircleColor = "#" + req.CircleColor
			o.ProgressColor = "#" + req.ProgressColor
			o.TextColor = "#" + req.TextColor
			o.TextSize = req.TextSize
			o.BackgroundColor = "#" + req.BackgroundColor
			o.Caption = req.Caption
			o.CaptionSize = req.CaptionSize
			o.CaptionColor = "#" + req.CaptionColor
			o.SegmentCount = req.SegmentCount
			o.SegmentGap = float64(req.SegmentGap)

			return nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		content := circular.SVG()
		SetHeadersResponseSvg(c.Response().Header())
		return c.Blob(http.StatusOK, "image/svg+xml", []byte(content))
	}

	if req.Style == "bar" {
		bar, err := gps.NewBar(func(o *gps.BarOptions) error {
			o.Progress = req.Progress
			o.Caption = req.Caption
			o.Width = req.Width
			o.Height = req.Height
			o.ProgressColor = "#" + req.ProgressColor
			o.TextColor = "#" + req.TextColor
			o.TextSize = req.TextSize
			o.Caption = req.Caption
			o.CaptionSize = req.CaptionSize
			o.CaptionColor = "#" + req.CaptionColor
			o.BackgroundColor = "#" + req.BackgroundColor
			o.CornerRadius = req.CornerRadius
			return nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		content := bar.SVG()
		SetHeadersResponseSvg(c.Response().Header())
		return c.Blob(http.StatusOK, "image/svg+xml", []byte(content))
	}

	if req.Style == "battery" {
		battery, err := gps.NewBattery(func(o *gps.BatteryOptions) error {
			o.Progress = req.Progress
			o.Caption = req.Caption
			o.Width = req.Width
			o.Height = req.Height
			o.ProgressColor = "#" + req.ProgressColor
			o.TextColor = "#" + req.TextColor
			o.TextSize = req.TextSize
			o.Caption = req.Caption
			o.CaptionSize = req.CaptionSize
			o.CaptionColor = "#" + req.CaptionColor
			o.BackgroundColor = "#" + req.BackgroundColor
			o.CornerRadius = req.CornerRadius
			return nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		content := battery.SVG()
		SetHeadersResponseSvg(c.Response().Header())
		return c.Blob(http.StatusOK, "image/svg+xml", []byte(content))
	}
	return nil
}
