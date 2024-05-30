package main

import (
	"fmt"
	"os"

	gps "github.com/kevincobain2000/go-progress-svg"
)

func main() {
	circular, err := gps.NewCircular(func(o *gps.CircularOptions) error {
		o.Progress = 60
		o.Size = 200
		o.CircleWidth = 16
		o.ProgressWidth = 16
		o.CircleColor = "#e0e0e0"
		o.ProgressColor = "#76e5b1"
		o.TextColor = "#6bdba7"
		o.TextSize = 52
		o.ShowPercentage = true
		o.BackgroundColor = ""
		o.Caption = "Coverage"
		o.CaptionSize = 20
		o.CaptionColor = "#000000"
		o.SegmentGap = 0
		return nil
	})
	if err != nil {
		fmt.Println("Error creating Circular:", err)
		return
	}

	content := circular.SVG()

	err = os.WriteFile("output.svg", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing SVG to file:", err)
	} else {
		fmt.Println("SVG written to output.svg successfully!")
	}
}
