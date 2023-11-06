package shapes

import (
	"fmt"
	"math"
)

type Cone struct {
	Height float64
	Radius float64
}

func (c Cone) Volume() float64 {
	return math.Pi * c.Height * math.Pow(c.Radius, 2) / 3
}

func (c Cone) Area() float64 {
	s := math.Sqrt(math.Pow(c.Height, 2) + math.Pow(c.Radius, 2))
	return math.Pi * c.Radius * (s + c.Radius)
}

func (c Cone) String() string {
	return fmt.Sprintf("Cone{Height: %.3f, Radius: %.3f}", c.Height, c.Radius)
}
