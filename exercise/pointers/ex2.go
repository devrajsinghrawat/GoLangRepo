package pointers

import (
	"fmt"
	"math"
)

type square struct {
	side float32
}

type circle struct {
	radius float32
}

type shape interface {
	area() float32
}

func (s square) area() float32 {
	return (s.side * s.side)
}

func (c *circle) area() float32 {
	return math.SqrtPi * c.radius * c.radius
}

func info(s shape) {
	fmt.Printf("Area of shape. is %f \n", s.area())
}

// Ex2 -Exported function to main - interface method
func Ex2() {

	s := square{side: 4.1}

	s1 := s.area()

	fmt.Printf("Area of Square is %f \n", s1)

	c := circle{radius: 4}

	c1 := c.area()
	fmt.Printf("Area of circle is %f \n", c1)

	info(s)

	// c2 := &c
	// fmt.Printf("Type of c2 - %T\n", c2)
	// info(c2)
	info(&c) // only accepts pointer if pointer receiver
}
