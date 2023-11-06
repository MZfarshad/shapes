package shapes

import (
	"fmt"
)

type Rectangle interface {
	D2
	Width() float64
	SetWidth(w float64) Rectangle
	Length() float64
	SetLength(l float64) Rectangle
}

type rectangle struct {
	width  float64
	length float64
}

func NewRectangle(width, length float64) Rectangle {
	return &rectangle{width: width, length: length}
}

func (r *rectangle) Width() float64 {
	if r == nil {
		return 0
	}
	return r.width
}

func (r *rectangle) Length() float64 {
	if r == nil {
		return 0
	}
	return r.length
}

func (r *rectangle) Area() float64 {
	return r.width * r.length
}

func (r *rectangle) Perimeter() float64 {
	return 2 * (r.width + r.length)
}

func (r *rectangle) String() string {
	return fmt.Sprintf("rectangle{width: %.3f, lentgh: %.3f}", r.width, r.length)
}

// SetWidth changes the receiver's width property to the given input, returning nil
// rectangle in case of zero.
func (r *rectangle) SetWidth(w float64) Rectangle {
	if w <= 0 {
		return nil
	}
	r.width = w
	return r
}

// SetLength changes the receiver's length property to the given input, returning nil
// rectangle in case of zero.
func (r *rectangle) SetLength(l float64) Rectangle {
	if r == nil {
		return NewRectangle(0, l)
	}
	r.length = l
	return r
}
