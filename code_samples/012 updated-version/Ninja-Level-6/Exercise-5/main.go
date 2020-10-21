package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return (c.radius * c.radius * math.Pi)
}

func (s square) area() float64 {
	return (s.side * s.side)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	circ := circle{
		radius: 3,
	}
	sq := square{
		side: 3,
	}
	info(sq)
	info(circ)
}
