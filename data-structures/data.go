package dataStructures

import (
	"errors"
	"math"
)

type Rectangle struct {
	length float64
	height float64
}

type Circle struct {
	radius float64
}

type Dict map[string]string

func Perimeter(rect Rectangle) (perim float64) {
	return 2 * (rect.length + rect.height)
}

func (r Rectangle) Area() (area float64) {
	return r.length * r.height
}

func (c Circle) Area() (area float64) {
	return math.Pi * c.radius * c.radius
}

func (dict Dict) Search(word string) (string, error) {
	searchResult, exists := dict[word]
	
	if !exists{
		return "", errors.New("Word not found")
	}

	return searchResult, nil
}