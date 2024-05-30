package main

import (
	"fmt"
	"os"

	gps "github.com/kevincobain2000/go-progress-svg"
)

func main() {
	battery, err := gps.NewBattery(func(o *gps.BatteryOptions) error {
		o.Progress = 70
		o.ProgressCaption = "1%"
		o.Width = 200
		o.Height = 50
		o.ProgressColor = "#76e5b1"
		o.TextColor = "black"
		o.TextSize = 20
		o.Caption = ""
		o.CaptionSize = 16
		o.CaptionColor = "#000000"
		o.BackgroundColor = "#e0e0e0"
		o.CornerRadius = 10
		return nil
	})
	if err != nil {
		fmt.Println("Error creating Circular:", err)
		return
	}

	content := battery.SVG()

	err = os.WriteFile("output.svg", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing SVG to file:", err)
	} else {
		fmt.Println("SVG written to output.svg successfully!")
	}
}
