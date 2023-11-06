package shapes

import (
	"math"
)

type Point struct {
	X float64
	Y float64
	Z float64
}

func (p Point) Distance(p1 Point) float64 {
	return math.Sqrt((math.Pow(p.X-p1.X, 2)) + (math.Pow(p.Y-p1.Y, 2)) + (math.Pow(p.Z-p1.Z, 2)))
}

func (p Point) Is(p1 Point) bool {
	if p.X == p1.X {
		if p.Y == p1.Y {
			if p.Z == p1.Z {
				return true
			}
		}
	}
	return false
}

//type points struct {
//	vertexa Point
//	vertexb Point
//	vertexc Point
//}

// func IsTriangle(p1 [3]Point) bool { >>> Bad naming in function argument

func IsTriangle(points [3]Point) bool {

	//ab := p.vertexa.Distance(p.vertexb)
	//ac := p.vertexa.Distance(p.vertexc)
	//bc := p.vertexb.Distance(p.vertexc)

	ab := points[0].Distance(points[1])
	ac := points[1].Distance(points[2])
	bc := points[2].Distance(points[0])

	if (ab+ac > bc) && (ab+bc > ac) && (bc+ac > ab) {
		return true
	}
	return false
}
