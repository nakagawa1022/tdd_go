package structure

import "math"

func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	ans := c.Radius * c.Radius * math.Pi
	return ans
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	width  float64
	height float64
}

func (t Triangle) Area() float64 {
	return t.width * t.height / 2
}
