package gps

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/kevincobain2000/go-progress-svg/utils"
)

var circularBottomTPL = `
<svg width="{{.Size}}" height="{{.Size}}" viewBox="-25 -25 250 250" version="1.1" xmlns="http://www.w3.org/2000/svg" style="transform:rotate(-90deg)">
{{if .BackgroundColor}}
  <circle r="90" cx="100" cy="100" fill="{{.BackgroundColor}}" />
{{end}}
  <circle r="90" cx="100" cy="100" fill="transparent" stroke="{{.CircleColor}}" stroke-width="{{.CircleWidth}}px" stroke-dasharray="565.48px" stroke-dashoffset="0"></circle>
  <circle r="90" cx="100" cy="100" stroke="{{.ProgressColor}}" stroke-width="{{.ProgressWidth}}px" stroke-linecap="round" stroke-dashoffset="{{.Offset}}px" fill="transparent" stroke-dasharray="565.48px"></circle>
  {{if .ShowPercentage}}
  <text x="100" y="100" font-family="sans-serif" fill="{{.TextColor}}" font-size="{{.TextSize}}px" font-weight="bold" text-anchor="middle" alignment-baseline="middle" transform="rotate(90 100 100)">{{.Progress}}%</text>
  {{else}}
  <text x="100" y="100" font-family="sans-serif" fill="{{.TextColor}}" font-size="{{.TextSize}}px" font-weight="bold" text-anchor="middle" alignment-baseline="middle" transform="rotate(90 100 100)">{{.Progress}}</text>
  {{end}}
  {{if .Caption}}
  <text x="-40" y="360" font-family="sans-serif" fill="{{.CaptionColor}}" font-size="{{.CaptionSize}}px" font-weight="bold" text-anchor="middle" transform="rotate(90 100 240)">{{.Caption}}</text>
  {{end}}
</svg>
`

var circularRightTPL = `
<svg width="{{.Size}}" height="{{.Size}}" viewBox="-25 -25 250 250" version="1.1" xmlns="http://www.w3.org/2000/svg" style="transform:rotate(-90deg)">
{{if .BackgroundColor}}
  <circle r="90" cx="100" cy="100" fill="{{.BackgroundColor}}" />
{{end}}
  <circle r="90" cx="100" cy="100" fill="transparent" stroke="{{.CircleColor}}" stroke-width="{{.CircleWidth}}px" stroke-dasharray="565.48px" stroke-dashoffset="0"></circle>
  <circle r="90" cx="100" cy="100" stroke="{{.ProgressColor}}" stroke-width="{{.ProgressWidth}}px" stroke-linecap="round" stroke-dashoffset="{{.Offset}}px" fill="transparent" stroke-dasharray="565.48px"></circle>
  {{if .ShowPercentage}}
  <text x="61px" y="115px" font-family="sans-serif" fill="{{.TextColor}}" font-size="{{.TextSize}}px" font-weight="bold" style="transform:rotate(90deg) translate(0px, -196px)">{{.Progress}}%</text>
  {{else}}
  <text x="71px" y="115px" font-family="sans-serif" fill="{{.TextColor}}" font-size="{{.TextSize}}px" font-weight="bold" style="transform:rotate(90deg) translate(0px, -196px)">{{.Progress}}</text>
  {{end}}
  {{if .Caption}}
  <text x="40%" y="220px" font-family="sans-serif" fill="{{.CaptionColor}}" font-size="{{.CaptionSize}}px" font-weight="bold" text-anchor="middle">{{.Caption}}</text>
  {{end}}
</svg>
`

type Circular struct {
	options *CircularOptions
	strings *utils.Strings
}

type CircularOptions struct {
	Progress        int
	Size            int
	CircleWidth     int
	ProgressWidth   int
	CircleColor     string
	ProgressColor   string
	TextColor       string
	TextSize        int
	ShowPercentage  bool
	BackgroundColor string
	Caption         string
	CaptionPos      string
	CaptionSize     int
	CaptionColor    string
	Offset          float64
}

type Option func(*CircularOptions) error

func NewCircular(opts ...Option) (*Circular, error) {
	options := &CircularOptions{
		Progress:        0,
		Size:            200,
		CircleWidth:     16,
		ProgressWidth:   16,
		CircleColor:     "#e0e0e0",
		ProgressColor:   "#76e5b1",
		TextColor:       "#6bdba7",
		TextSize:        52,
		ShowPercentage:  true,
		BackgroundColor: "",
		Caption:         "",
		CaptionPos:      "bottom",
		CaptionSize:     20,
		CaptionColor:    "#000000",
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}

	options.Offset = 565.48 * (1 - float64(options.Progress)/100)

	return &Circular{
		options: options,
		strings: utils.NewStrings(),
	}, nil
}

func (c *Circular) SVG() string {
	tpl := ""
	if c.options.CaptionPos == "right" {
		tpl = c.strings.StripXMLWhitespace(circularRightTPL)
	} else {
		tpl = c.strings.StripXMLWhitespace(circularBottomTPL)
	}

	tmpl, err := template.New("svg").Parse(tpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return ""
	}

	var rendered strings.Builder
	err = tmpl.Execute(&rendered, c.options)
	if err != nil {
		fmt.Println("Error rendering template:", err)
		return ""
	}
	return rendered.String()
}
