  
package main

import (
	"fmt"
)

type Shape interface {
	fmt.Stringer
}

type Circle struct {
	radius int
}

type Rectangle struct {
	width  int
	lenght int
}

func (rec Rectangle) String() string {
	return fmt.Sprintf("Square of width %d and length %d", rec.width, rec.lenght)
}

func (cir Circle) String() string {
	return fmt.Sprintf("Circle of radius %d", cir.radius)
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}
	shapes := []Shape{r, c}
	for _, s := range shapes {
		fmt.Println(s)
	}
}