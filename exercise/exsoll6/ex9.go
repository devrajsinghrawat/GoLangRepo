package exsoll6

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

func (c circle) area() float32 {
	return math.SqrtPi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

// Ex9 -Exported function to main - Defer exercise
func Ex9() {

	s := square{side: 4.1}

	s1 := s.area()

	fmt.Printf("Area of Square is %f \n", s1)

	c := circle{radius: 4}

	c1 := c.area()
	fmt.Printf("Area of circle is %f \n", c1)

	info(c)
	info(s)

}
