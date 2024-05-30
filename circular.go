package gps

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/kevincobain2000/go-progress-svg/utils"
)

var circularTPL = `
<svg width="{{.Size}}" height="{{.Hsize}}" viewBox="0 0 200 200" version="1.1" xmlns="http://www.w3.org/2000/svg">
{{if .BackgroundColor}}
  <circle r="90" cx="100" cy="100" fill="{{.BackgroundColor}}" />
{{end}}
  <circle r="90" cx="100" cy="100" fill="transparent" stroke="{{.CircleColor}}" stroke-width="{{.CircleWidth}}px" stroke-dasharray="{{.CircleDashArray}}" stroke-dashoffset="0"></circle>
  <circle r="90" cx="100" cy="100" stroke="{{.ProgressColor}}" stroke-width="{{.ProgressWidth}}px" stroke-linecap="round" stroke-dasharray="{{.ProgressDashArray}}" stroke-dashoffset="{{.Offset}}px" fill="transparent" style="transform:rotate(-90deg); transform-origin: 50% 50%;"></circle>
  {{if .ShowPercentage}}
  <text x="100" y="100" font-family="sans-serif" fill="{{.TextColor}}" font-size="{{.TextSize}}px" font-weight="bold" text-anchor="middle" alignment-baseline="middle">{{.Progress}}%</text>
  {{else}}
  <text x="100" y="100" font-family="sans-serif" fill="{{.TextColor}}" font-size="{{.TextSize}}px" font-weight="bold" text-anchor="middle" alignment-baseline="middle">{{.Progress}}</text>
  {{end}}
  {{if .Caption}}
  <text x="100" y="210" font-family="sans-serif" fill="{{.CaptionColor}}" font-size="{{.CaptionSize}}px" text-anchor="middle">{{.Caption}}</text>
  {{end}}
</svg>
`

type Circular struct {
	options *CircularOptions
	strings *utils.Strings
}

type CircularOptions struct {
	Progress          int
	Size              int
	Hsize             int
	CircleWidth       int
	ProgressWidth     int
	CircleColor       string
	ProgressColor     string
	TextColor         string
	TextSize          int
	ShowPercentage    bool
	BackgroundColor   string
	Caption           string
	CaptionSize       int
	CaptionColor      string
	Offset            float64
	SegmentCount      int
	SegmentGap        float64
	CircleDashArray   string
	ProgressDashArray string
}

type Option func(*CircularOptions) error

func NewCircular(opts ...Option) (*Circular, error) {
	options := &CircularOptions{
		Progress:        0,
		Size:            200,
		Hsize:           230,
		CircleWidth:     16,
		ProgressWidth:   16,
		CircleColor:     "#e0e0e0",
		ProgressColor:   "#76e5b1",
		TextColor:       "#6bdba7",
		TextSize:        52,
		ShowPercentage:  true,
		BackgroundColor: "",
		Caption:         "",
		CaptionSize:     20,
		CaptionColor:    "#000000",
		SegmentCount:    10,
		SegmentGap:      2,
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	if options.Caption != "" {
		options.Hsize = options.Size + 30
	} else {
		options.Hsize = options.Size
	}

	totalCircumference := 565.48
	segmentLength := (totalCircumference - options.SegmentGap*float64(options.SegmentCount)) / float64(options.SegmentCount)
	options.CircleDashArray = fmt.Sprintf("%f %f", segmentLength, options.SegmentGap)

	progressSegments := float64(options.SegmentCount) * float64(options.Progress) / 100
	progressSegmentLength := progressSegments*(segmentLength+options.SegmentGap) - options.SegmentGap
	options.ProgressDashArray = fmt.Sprintf("%f %f", progressSegmentLength, totalCircumference-progressSegmentLength)
	options.Offset = 0

	return &Circular{
		options: options,
		strings: utils.NewStrings(),
	}, nil
}

func (c *Circular) SVG() string {
	tpl := c.strings.StripXMLWhitespace(circularTPL)

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
