package shapes_test

import (
	shapes2 "playground"
	"testing"
)

func TestNewWingedTriangle(t *testing.T) {
	wingedtri, err := shapes2.NewWingedTriangle([3]float64{3, 4, 5})
	if err != nil {
		t.Errorf("expected nil error,got %s", err.Error())
	}
	if wingedtri == nil {
		t.Error("expected a newwingedtriangle,but got nil !!??")
	}
	if wingedtri.AB() != 3 {
		if wingedtri.BC() != 4 {
			if wingedtri.CA() != 5 {
				t.Errorf("expected ab=3 bc=4 ca=5 , got ab=%.2f bc=%.2f ca=%.2f", wingedtri.AB(), wingedtri.BC(), wingedtri.CA())
			}
		}
	}

}

func TestNewTriangleWithPoint(t *testing.T) {
	points := [3]shapes2.Point{
		{1, 2, 3},
		{4, 5, 8},
		{3, 5, 9},
	}
	newTriangleWithPoints, err := shapes2.NewTriangleWithPoints(points)

	if err != nil {
		t.Errorf("expected nil error,got %s", err.Error())
	}
	if newTriangleWithPoints == nil {
		t.Error("expected a newtriangle,but got nil !!??")
	}
	if !shapes2.IsTriangle(points) {
		t.Error("expect,points are equal and can not make a triangle,but got a triangle")
	}
}
