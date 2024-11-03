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

func (c *Circle) init(r float64) {
	c.radius = r
}

type Rectangle struct {
	length float64
	height float64
}

func (r *Rectangle) init(l float64, h float64) {
	r.length = l
	r.height = h
}
func (r *Rectangle) Draw() float64 {
	return r.length * r.height
}
