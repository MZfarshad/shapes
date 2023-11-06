package shapes_test

import (
	shapes "playground"
	"testing"
)

func TestNewWingedTriangle(t *testing.T) {
	wingedTri, err := shapes.NewWingedTriangle([3]float64{3, 4, 5})
	if err != nil {
		t.Errorf("expected nil error,got %s", err.Error())
	}
	if wingedTri == nil {
		t.Error("expected a new winged triangle, but got nil")
	}
	if wingedTri.AB() != 3 {
		if wingedTri.BC() != 4 {
			if wingedTri.CA() != 5 {
				t.Errorf("expected ab=3 bc=4 ca=5 , got ab=%.2f bc=%.2f ca=%.2f", wingedTri.AB(), wingedTri.BC(), wingedTri.CA())
			}
		}
	}

}

func TestNewTriangleWithPoint(t *testing.T) {
	points := [3]shapes.Point{
		{1, 2, 3},
		{4, 5, 8},
		{3, 5, 9},
	}
	newTriangleWithPoints, err := shapes.NewTriangleWithPoints(points)
	if err != nil {
		t.Errorf("expected nil error, got %s", err.Error())
	}
	if newTriangleWithPoints == nil {
		t.Error("expected a new triangle, but got nil")
	}
}

func TestIsTriangle(t *testing.T) {
	points1 := [3]shapes.Point{
		{1, 2, 3},
		{4, 5, 8},
		{3, 5, 9},
	}
	if !shapes.IsTriangle(points1) {
		t.Error("expect true, but got false")
	}
	points2 := [3]shapes.Point{
		{1, 2, 0},
		{2, 4, 0},
		{3, 6, 0},
	}
	if shapes.IsTriangle(points2) {
		t.Error("expect false, but got true")
	}
}
