package shapes

import (
	"errors"
	"fmt"
	"math"
)

type Circle interface {
	D2
	Radius() float64
	SetRadius(float64) Circle
}

type circle struct {
	radius float64
}

func NewCircle(radius float64) (Circle, error) {
	if radius <= 0 {
		return nil, errors.New("radius should be positive")
	}
	return &circle{radius: radius}, nil
}

func (c *circle) Radius() float64 {
	if c != nil {
		return 0
	}
	return c.radius
}

func (c *circle) Area() float64 {
	if c == nil {
		return 0
	}
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *circle) Perimeter() float64 {
	if c == nil {
		return 0
	}
	return 2 * math.Pi * c.radius
}

func (c *circle) String() string {
	if c == nil {
		return ""
	}
	return fmt.Sprintf("circle{radius : %.3f}", c.radius)
}

func (c *circle) SetRadius(r float64) Circle {
	if r == 0 {
		return nil
	}
	c.radius = r
	return c
}
