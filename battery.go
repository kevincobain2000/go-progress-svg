package gps

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/kevincobain2000/go-progress-svg/utils"
)

var batteryTPL = `
<svg width="{{.TotalWidth}}" height="{{.TotalHeight}}" version="1.1" xmlns="http://www.w3.org/2000/svg">
  <rect x="0" y="0" width="{{.Width}}" height="{{.Height}}" fill="{{.BackgroundColor}}" rx="{{.CornerRadius}}" ry="{{.CornerRadius}}" />
  <rect x="0" y="0" width="{{.ProgressWidth}}" height="{{.Height}}" fill="{{.ProgressColor}}" rx="{{.CornerRadius}}" ry="{{.CornerRadius}}" />
  <rect x="{{.CapX}}" y="{{.CapY}}" width="{{.CapWidth}}" height="{{.CapHeight}}" fill="{{.BackgroundColor}}" rx="{{.CapCornerRadius}}" ry="{{.CapCornerRadius}}" />
  <rect x="{{.CapX}}" y="{{.CapY}}" width="{{.CapWidth}}" height="{{.CapHeight}}" fill="{{.ProgressColor}}" rx="{{.CapCornerRadius}}" ry="{{.CapCornerRadius}}" />
  {{range .Segments}}
  <rect x="{{.X}}" y="0" width="{{$.SegmentGapWidth}}" height="{{$.Height}}" fill="{{$.BackgroundColor}}" />
  {{end}}
  {{if .ProgressCaption}}
  <text x="50%" y="{{.HeightHalf}}" font-family="sans-serif" fill="{{.TextColor}}" font-size="{{.TextSize}}px" font-weight="bold" text-anchor="middle" alignment-baseline="middle">{{.ProgressCaption}}</text>
  {{end}}
  {{if .Caption}}
  <text x="50%" y="{{.CaptionY}}" font-family="sans-serif" fill="{{.CaptionColor}}" font-size="{{.CaptionSize}}px" text-anchor="middle">{{.Caption}}</text>
  {{end}}
</svg>
`

type Battery struct {
	options *BatteryOptions
	strings *utils.Strings
}

type BatteryOptions struct {
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
	TotalWidth      int
	TotalHeight     int
	HeightHalf      int
	CaptionY        int
	CornerRadius    int
	CapWidth        int
	CapHeight       int
	CapHeightHalf   int
	CapX            int
	CapY            int
	CapCornerRadius int
	SegmentCount    int
	SegmentGapWidth int
	Segments        []Segment
}

type Segment struct {
	X int
}

type BatteryOption func(*BatteryOptions) error

func NewBattery(opts ...BatteryOption) (*Battery, error) {
	options := &BatteryOptions{
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
		CapWidth:        10,
		CapHeight:       30,
		CapCornerRadius: 5,
		SegmentCount:    5,
		SegmentGapWidth: 5,
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	if options.Progress < 3 {
		options.Progress = 3
	}

	options.ProgressWidth = fmt.Sprintf("%d%%", options.Progress)
	options.TotalWidth = options.Width + options.CapWidth
	options.TotalHeight = options.Height + options.CaptionSize + 10 // Additional space for caption
	options.HeightHalf = options.Height / 2
	options.CaptionY = options.Height + options.CaptionSize // Position caption below the bar
	options.CapHeightHalf = options.CapHeight / 2
	options.CapX = options.Width
	options.CapY = options.HeightHalf - options.CapHeightHalf

	segmentWidth := (options.Width - (options.SegmentCount-1)*options.SegmentGapWidth) / options.SegmentCount
	for i := 1; i < options.SegmentCount; i++ {
		options.Segments = append(options.Segments, Segment{
			X: i * (segmentWidth + options.SegmentGapWidth),
		})
	}

	return &Battery{
		options: options,
		strings: utils.NewStrings(),
	}, nil
}

func (b *Battery) SVG() string {
	tpl := b.strings.StripXMLWhitespace(batteryTPL)
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
