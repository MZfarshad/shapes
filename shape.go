package shapes

// D3 represents three-dimensional (3D) shapes.
type D3 interface {
	Area() float64
	Volume() float64
}

// D2 represents two-dimensional (2D) shapes.
type D2 interface {
	Area() float64
	Perimeter() float64
}
