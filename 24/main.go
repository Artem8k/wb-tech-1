package main

import "fmt"

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (p *Point) findDistance() int {
	if p.X > p.Y {
		return p.X - p.Y
	} else {
		return p.Y - p.X
	}
}

func main() {
	point := NewPoint(-10, 12)

	fmt.Println(point.findDistance())
}
