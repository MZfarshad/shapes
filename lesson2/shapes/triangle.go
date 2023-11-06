package shapes

import (
	"errors"
	"fmt"
	"math"
)

type Triangle interface {
	D2
	AB() float64
	BC() float64
	CA() float64
	//A() Point
	//B() Point
	//C() Point
}
type wingedTriangle struct {
	wingAB float64
	wingBC float64
	wingCA float64
}

func (t *wingedTriangle) String() string {
	return fmt.Sprintf("points of triangle {%.3f,%.3f,%.3f}", t.wingAB, t.wingBC, t.wingCA)
}

func (t *wingedTriangle) AB() float64 {
	return t.wingAB
}

func (t *wingedTriangle) BC() float64 {
	return t.wingBC
}

func (t *wingedTriangle) CA() float64 {
	return t.wingCA
}

func (t *wingedTriangle) Validate() error {
	if (t.wingAB+t.wingCA > t.wingBC) && (t.wingAB+t.wingBC > t.wingCA) && (t.wingBC+t.wingCA > t.wingAB) {
		return nil
	}
	return fmt.Errorf("there is no valid triangle with given {%0.3f,%0.3f,%0.3f} wings in real world", t.wingAB, t.wingBC, t.wingCA)
}

func NewWingedTriangle(wings [3]float64) (Triangle, error) {
	wingedTri := wingedTriangle{wingAB: wings[0], wingBC: wings[1], wingCA: wings[2]}
	if err := wingedTri.Validate(); err != nil {
		return nil, err
	}
	return &wingedTri, nil
}

// @TODO: NewTriangleWing should not accept 3 points.
//func NewTriangleWing(wing [3]Point) (Triangle, error) {
//	if wing[0][0] < 0 && wing[0][1] < 0 && wing[0][2] < 0 {
//		if IsTriangle() {
//
//		}
//	}
//	return nil, errors.New("length of line cannot be zero or sum of two wings is less than as other wing")
//}

func (t *wingedTriangle) Area() float64 {
	s := (t.wingAB + t.wingBC + t.wingCA) / 2
	return math.Sqrt(s * (s - t.wingAB) * (s - t.wingBC) * (s - t.wingCA))
}

func (t *wingedTriangle) Perimeter() float64 {
	return t.wingAB + t.wingBC + t.wingCA
}

type triangle struct {
	vertexA Point
	vertexB Point
	vertexC Point
}

func NewTriangleWithPoints(points [3]Point) (Triangle, error) {
	//if vertex[0].Is(vertex[1]) && vertex[0].Is(vertex[2]) && vertex[1].Is(vertex[2]) {
	//	if IsTriangle(vertex) {
	//		return &triangle{
	//			vertexA: vertex[0],
	//			vertexB: vertex[1],
	//			vertexC: vertex[2],
	//		}, nil
	//	}
	//}
	//return nil, errors.New("there is no such triangle with this specification")
	if IsTriangle(points) {
		return &triangle{
			vertexA: points[0],
			vertexB: points[1],
			vertexC: points[2],
		}, nil
	}
	return nil, errors.New("there is no such triangle with this specification")

}

func (t *triangle) AB() float64 {
	return t.vertexA.Distance(t.vertexB)
}

func (t *triangle) BC() float64 {
	return t.vertexB.Distance(t.vertexC)
}

func (t *triangle) CA() float64 {
	return t.vertexC.Distance(t.vertexA)
}

func (t *triangle) A() Point {
	return t.vertexA
}
func (t *triangle) B() Point {
	return t.vertexB
}
func (t *triangle) C() Point {
	return t.vertexC
}

func (t *triangle) Area() float64 {
	s := t.Perimeter() / 2
	return math.Sqrt(s * (s - t.AB()) * (s - t.BC()) * (s - t.CA()))
}
func (t *triangle) Perimeter() float64 {
	return t.AB() + t.BC() + t.CA()
}

func (t *triangle) String() string {
	return fmt.Sprintf("points of triangle {%.3f,%.3f,%.3f}", t.vertexA, t.vertexB, t.vertexC)
}
