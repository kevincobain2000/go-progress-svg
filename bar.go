package gps

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/kevincobain2000/go-progress-svg/utils"
)

var barTPL = `
<svg width="{{.Width}}" height="{{.TotalHeight}}" version="1.1" xmlns="http://www.w3.org/2000/svg">
  <rect x="0" y="0" width="100%" height="{{.Height}}" fill="{{.BackgroundColor}}" rx="{{.CornerRadius}}" ry="{{.CornerRadius}}" />
  <rect x="0" y="0" width="{{.ProgressWidth}}" height="{{.Height}}" fill="{{.ProgressColor}}" rx="{{.CornerRadius}}" ry="{{.CornerRadius}}" />
  {{if .ProgressCaption}}
  <text x="50%" y="{{.HeightHalf}}" font-family="sans-serif" fill="{{.TextColor}}" font-size="{{.TextSize}}px" font-weight="bold" text-anchor="middle" alignment-baseline="middle">{{.ProgressCaption}}</text>
  {{end}}
  {{if .Caption}}
  <text x="50%" y="{{.CaptionY}}" font-family="sans-serif" fill="{{.CaptionColor}}" font-size="{{.CaptionSize}}px" text-anchor="middle">{{.Caption}}</text>
  {{end}}
</svg>
`

type Bar struct {
	options *BarOptions
	strings *utils.Strings
}

type BarOptions struct {
	Progress        int
	ProgressCaption string
	Width           int
	Height          int
	ProgressWidth   string
	ProgressColor   string
	TextColor       string
	TextSize        int
	Caption         string
	CaptionSize     int
	CaptionColor    string
	BackgroundColor string
	TotalHeight     int
	HeightHalf      int
	CaptionY        int
	CornerRadius    int
}

type BarOption func(*BarOptions) error

func NewBar(opts ...BarOption) (*Bar, error) {
	options := &BarOptions{
		Progress:        0,
		Width:           200,
		Height:          50,
		ProgressColor:   "#76e5b1",
		TextColor:       "#6bdba7",
		TextSize:        20,
		Caption:         "",
		CaptionSize:     16,
		CaptionColor:    "#000000",
		BackgroundColor: "#e0e0e0",
		CornerRadius:    10,
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}

	options.ProgressWidth = fmt.Sprintf("%d%%", options.Progress)
	options.TotalHeight = options.Height + options.CaptionSize + 10 // Additional space for caption
	options.HeightHalf = options.Height / 2
	options.CaptionY = options.Height + options.CaptionSize // Position caption below the bar

	return &Bar{
		options: options,
		strings: utils.NewStrings(),
	}, nil
}

func (b *Bar) SVG() string {
	tpl := b.strings.StripXMLWhitespace(barTPL)
	tmpl, err := template.New("svg").Parse(tpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return ""
	}

	var rendered strings.Builder
	err = tmpl.Execute(&rendered, b.options)
	if err != nil {
		fmt.Println("Error rendering template:", err)
		return ""
	}

	return rendered.String()
}
