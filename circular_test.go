package gps

import (
	"strings"
	"testing"
)

func TestCircularSVG(t *testing.T) {
	// Test rendering the SVG with default options
	circular, err := NewCircular()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	svg := circular.SVG()
	expectedStart := `<svg width="200" height="200"`
	if !strings.HasPrefix(svg, expectedStart) {
		t.Errorf("Expected SVG to start with %s, got %s", expectedStart, svg)
	}

	// Test rendering the SVG with custom options
	circular, err = NewCircular(func(o *CircularOptions) error {
		o.Progress = 75
		o.Size = 300
		o.CircleColor = "#ff0000"
		o.ProgressColor = "#00ff00"
		o.TextColor = "#0000ff"
		o.ShowPercentage = false
		return nil
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	svg = circular.SVG()
	expectedContains := `stroke="#00ff00"`
	if !strings.Contains(svg, expectedContains) {
		t.Errorf("Expected SVG to contain %s, got %s", expectedContains, svg)
	}
}
