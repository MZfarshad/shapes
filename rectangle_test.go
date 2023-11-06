package shapes_test

import (
	"playground"
	"testing"
)

var (
	width  = 4.5
	length = 6.5
)

func TestNewRectangle(t *testing.T) {
	newRectangle := shapes.NewRectangle(width, length)
	if newRectangle == nil {
		t.Error("expected new rectangle, got nil")
	}
}

func TestWidth(t *testing.T) {
	newRectangle := shapes.NewRectangle(width, length)
	if newRectangle.Width() != width {
		t.Errorf("expected %0.2f, got %0.2f", width, newRectangle.Width())
	}
}

func TestLength(t *testing.T) {
	newRectangle := shapes.NewRectangle(width, length)
	if newRectangle.Length() != length {
		t.Errorf("expected %0.2f, got %0.2f", length, newRectangle.Length())
	}
}
func TestSetWidth(t *testing.T) {
	newRectangle := shapes.NewRectangle(width, length)
	newWidth := 6.7
	if newRectangle.SetWidth(newWidth) == nil {
		t.Error("expected a new rectangle, got nil")
	}
	if newRectangle.Width() != newWidth {
		t.Errorf("expected %0.2f, got %0.2f", newWidth, newRectangle.Width())
	}
}

func TestSetLength(t *testing.T) {
	newRectangle := shapes.NewRectangle(width, length)
	newLength := 7.7
	if newRectangle.SetLength(newLength) == nil {
		t.Error("expected a new rectangle, got nil")
	}
	if newRectangle.Length() != newLength {
		t.Errorf("expected %0.2f, got %0.2f", newLength, newRectangle.Length())
	}
}
