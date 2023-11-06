package shapes_test

import (
	"playground"
	"testing"
)

var newRectangle = shapes.NewRectangle(4.5, 6.5)

func TestNewRectangle(t *testing.T) {
	if newRectangle == nil {
		t.Errorf("expected width of rectangle = %.2f and lenght of rectangle = %.2f , got nil", newRectangle.Width(), newRectangle.Length())
	}
}
func TestWidth(t *testing.T) {
	if newRectangle.Width() == 0 {
		t.Error("expected width of the rectangle be a float number >= 0 , got zero")
	}
}
func TestLength(t *testing.T) {
	if newRectangle.Length() == 0 {
		t.Error("expected length of the rectangle be a float number >= 0 , got zero")
	}
}
func TestSetWidth(t *testing.T) {
	if newRectangle.SetWidth(6.7) == nil {
		t.Error("expected a new width of rectangle,got nil")
	}
}

func TestSetLength(t *testing.T) {
	if newRectangle.SetLength(7.7) == nil {
		t.Error("expected a new length of rectangle,got nil")
	}
}
