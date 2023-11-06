package main

import (
	"fmt"
	"log"
	"playground/lesson2/shapes"
)

func main() {
	// You don't need the two arguments(x, y float64) in shapes.D3 interface.
	// Interfaces do not know anything about their real implementations.
	// of their behaviour. If you define an argument for one of your interface's methods
	// you should obviously declare the name of the input arguments. For example, what is the
	// meaning of the x, y characters? Are they refer to a length, a width, a height, or a circle?
	// In fact, you didn't use x, y argument in real implementation of the Cone's methods, e.g. GetArea, and GetVolume.
	c := shapes.Cone{Height: 3.5, Radius: 7.8}
	fmt.Println("====================", c)
	// You can call the cone's methods either directly, as following
	area := c.Area()
	volume := c.Volume()
	fmt.Println(">>> Calling Cone Directly")
	fmt.Printf("\tArea: %.3f\n", area)
	fmt.Printf("\tVolume:%.3f\n", volume)
	// or indirectly, from interface definition: Two different approach, with same results!
	areaOfCone := shapes.D3.Area(c)
	volumeOfCone := shapes.D3.Volume(c)
	fmt.Println(">>> Using shapes.D3 interface")
	fmt.Printf("\tArea: %.3f\n", areaOfCone)
	fmt.Printf("\tVolume:%.3f\n", volumeOfCone)

	// @TODO: Review the following codes and try to get the points.
	rectangle := shapes.NewRectangle(2.5, 3.5)
	fmt.Println("====================", rectangle)
	fmt.Printf("\tArea: %.3f\n", rectangle.Area())
	fmt.Printf("\tPerimeter: %.3f\n", rectangle.Perimeter())

	fmt.Println(">>> Changing width to 5")
	fmt.Printf("\tArea: %.3f\n", rectangle.SetWidth(5).Area())
	fmt.Printf("\tPerimeter: %.3f\n", rectangle.Perimeter())
	fmt.Println(">>> Changing length to 5")
	fmt.Printf("\tArea: %.3f\n", rectangle.SetLength(5).Area())
	fmt.Printf("\tPerimeter: %.3f\n", rectangle.Perimeter())

	circle, err := shapes.NewCircle(7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("====================", circle)
	fmt.Printf("\tArea: %.3f\n", circle.Area())
	fmt.Printf("\tArea: %.3f\n", circle.Perimeter())
	fmt.Println(">>> Changing circle to 3")
	fmt.Printf("\tArea: %.3f\n", circle.SetRadius(3).Area())
	fmt.Printf("\tArea: %.3f\n", circle.Perimeter())

	pointedTriangle, err := shapes.NewTriangleWithPoints([3]shapes.Point{
		{X: 3, Y: 0, Z: 0},
		{X: 0, Y: 3, Z: 0},
		{X: 0, Y: 0, Z: 3},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("====================", pointedTriangle)
	fmt.Printf("\tArea: %.3f\n", pointedTriangle.Area())
	fmt.Printf("\tPerimeter: %.3f\n", pointedTriangle.Perimeter())

	wingedTriangle, err := shapes.NewWingedTriangle([3]float64{3, 4, 5})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("====================", wingedTriangle)
	fmt.Printf("\tArea: %.3f\n", wingedTriangle.Area())
	fmt.Printf("\tPerimeter: %.3f\n", wingedTriangle.Perimeter())

	invalidWingedTriangle, err := shapes.NewWingedTriangle([3]float64{1, 2, 3})
	if err != nil {
		log.Println("invalidWingedTriangle variable value: ", invalidWingedTriangle)
		log.Fatal(err)
	}
}
