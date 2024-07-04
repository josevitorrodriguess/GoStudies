package interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return (c.Radius * c.Radius) * math.Pi
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %f\n", s.area())

}
