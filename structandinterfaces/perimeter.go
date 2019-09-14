package structandinterfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2.0)
}
