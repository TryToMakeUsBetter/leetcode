package design

import "math"

type shape interface {
	Draw()
}

type Circle struct {
	radius float64
}

func (c *Circle) Draw() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	length float64
	height float64
}

func (r *Rectangle) Draw() float64 {
	return r.length * r.height
}
