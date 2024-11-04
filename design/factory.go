package design

import (
	"errors"
	"math"
)

type ShapeFactory interface {
	CreateShape(name string) Shape
}

type ShapeFactoryStruct struct {
	GuestShape *Shape
}

func (sf *ShapeFactoryStruct) CreateShape(name string) (Shape, error) {
	switch name {
	case "Circle":
		return new(Circle), nil
	case "Rectangle":
		return new(Rectangle), nil
	default:
		return nil, errors.New("不符合需求的图形")
	}
}

type Shape interface {
	Draw() float64
	Init(args ...float64)
}

type Circle struct {
	radius float64
}

func (c Circle) Draw() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Init(args ...float64) {
	c.radius = args[0]
}

type Rectangle struct {
	height float64
	length float64
}

func (rc Rectangle) Draw() float64 {
	return rc.height * rc.length
}

func (rc *Rectangle) Init(args ...float64) {
	rc.height = args[0]
	rc.length = args[1]
}
