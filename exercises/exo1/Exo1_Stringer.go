//Define a Rectangle type that implements fmt.Stringer
//Define a Circle type that implements fmt.Stringer
//Define a Shape interface, implemented by Circle and Rectangle such as this code:
//Bonus point if Shape forces to implement Stringer

package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Shape interface {
	String() string
}

type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle of width %v  and length %v \n\n", r.Width, r.Length)
}

type Circle struct {
	Radius float64
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle of radius %v \n\n", c.Radius)
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
	}
}

//func main() {
//	a := Rectangle{12, 42}
//	b := Circle{2.90}
//	fmt.Println(a, b)
//}
