package shapes_test

import (
	"playground/lesson2/shapes"
	"testing"
)

func TestIs(t *testing.T) {
	p0 := shapes.Point{X: 3, Y: 6, Z: 9}
	p1 := p0
	if !p0.Is(p1) {
		t.Error("expect true,got false")
	}
}
