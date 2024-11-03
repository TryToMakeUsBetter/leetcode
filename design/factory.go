package design

import "math"

type Shape interface {
	Draw() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Draw() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Init(r float64) {
	c.radius = r
}

type Rectangle struct {
	length float64
	height float64
}

func (r *Rectangle) Init(l float64, h float64) {
	r.length = l
	r.height = h
}
func (r Rectangle) Draw() float64 {
	return r.length * r.height
}
