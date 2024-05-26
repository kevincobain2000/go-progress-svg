package gps

import (
	"strings"
	"testing"
)

func TestBarSVG(t *testing.T) {
	// Test rendering the SVG with default options
	bar, err := NewBar()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	svg := bar.SVG()
	expectedStart := `<svg width="200" height="76"`
	if !strings.HasPrefix(svg, expectedStart) {
		t.Errorf("Expected SVG to start with %s, got %s", expectedStart, svg)
	}

	// Test rendering the SVG with custom options
	bar, err = NewBar(
		func(o *BarOptions) error {
			o.Progress = 75
			o.ProgressColor = "#ff0000"
			o.BackgroundColor = "#0000ff"
			return nil
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	svg = bar.SVG()
	expectedStart = `<svg width="200" height="76"`
	if !strings.HasPrefix(svg, expectedStart) {
		t.Errorf("Expected SVG to start with %s, got %s", expectedStart, svg)
	}
}
