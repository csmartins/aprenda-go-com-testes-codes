package dataStructures

import "math"

type Rectangle struct {
	length float64
	height float64
}

type Circle struct {
	radius float64
}

func Perimeter(rect Rectangle) (perim float64) {
	return 2 * (rect.length + rect.height)
}

func (r Rectangle) Area() (area float64) {
	return r.length * r.height
}

func (c Circle) Area() (area float64) {
	return math.Pi * c.radius * c.radius
}