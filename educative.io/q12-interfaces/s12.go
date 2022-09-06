package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perim() float64
	introduce() string
}

type Rectangle struct {
	Width  float64
	Hieght float64
}

func (r Rectangle) introduce() string {
	return fmt.Sprintf(
		"A rectangle with %0.2f m of \"Width\" and %0.2f m of \"Hieght\"",
		r.Width,
		r.Hieght,
	)
}

func (r Rectangle) area() float64 {
	return r.Width * r.Hieght
}

func (r Rectangle) perim() float64 {
	return 2 * (r.Width + r.Hieght)
}

type Circle struct {
	Radius float64
}

func (c Circle) introduce() string {
	return fmt.Sprintf("A Circle with %0.2f m of \"Radius\"", c.Radius)
}
func (c Circle) area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 10, Hieght: 15},
		Circle{Radius: 10},
	}
	for i, shape := range shapes {
		fmt.Printf("shape%d: %s \n", i+1, shape.introduce())
		fmt.Println("	Area(m^2):", shape.area())
		fmt.Println("	Perim(m):", shape.perim())
	}
}
