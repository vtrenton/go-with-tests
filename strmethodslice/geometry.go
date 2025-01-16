package geometry

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimiter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
