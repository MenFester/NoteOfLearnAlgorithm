package main

import (
	"fmt"
	"math"
)


// sytax of interface:
// Type <Interfce name> interface {
//     <Method name>(<Parameter Type List>) <Return type List>
// }
// When an object implement a particular interfac, its object can be
// assigned to an interface type variable

//Shape :
// user define an interface 
type Shape interface {
	Area() float64
	Perimeter() float64
}

//Rect :
type Rect struct {
	width float64
	height float64
}

//Circle :
type Circle struct {
	radius float64
}

//Area :
// designer implement the interface Shapes
func (r Rect) Area() float64 {
	return r.width * r.height
}


//Perimeter :
// designer implement the interface Shapes
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}


//Area :
// designer implement the interface Shapes
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//Perimeter :
// designer implement the interface Shapes
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

//TotalArea :
// user use the interface, acept the parameter whicc the designer had defined
func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}

	return area
}

//TotalPerimeter :
// user use the interface, acept the parameter whicc the designer had defined
func TotalPerimeter(shapes ...Shape) float64 {
	var peri float64
	for _, s := range shapes {
		peri += s.Perimeter()
	}

	return peri
}

func main() {
	r := Rect{width: 10, height: 10}   // user use the struct
	c := Circle{radius: 10}

	fmt.Println("Total Area:: ", TotalArea(r, c))    // user use the interface
	fmt.Println("Total Perimeter:: ", TotalPerimeter(r, c))
}