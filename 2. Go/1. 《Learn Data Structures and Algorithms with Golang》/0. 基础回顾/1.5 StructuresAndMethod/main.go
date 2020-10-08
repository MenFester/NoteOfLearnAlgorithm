package main

import (
	"fmt"
)

type student struct {
	rollNo int
	name string
}

//Rect struct :
type Rect struct {
	width float64
	height float64
}

//Area :
// method of Rect
func (r Rect) Area() float64 {
	return r.width * r.height
}

//Perimeter :
// method of Rect
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// syntax of accessor function:
// func (r <Receiver Data Type>) <Funtion Name>(<Parameter List>) (<Return List>)

// syntax of modifier function:
// func (r *<Receiver Data Type>) <Function Name)(<Parameter List>) (<Return List>)


//MyInt :
type MyInt int

func (data MyInt) increment1() {
	data = data + 1
}

func (data *MyInt) increment2() {
	*data = *data + 1
}

func main() {
	stud := student{1, "Johnny"}
	fmt.Println(stud)
	fmt.Println("Student name:: ", stud.name)
	pstud := &stud
	fmt.Println("Student name:: ", pstud.name)    // 与 (*pstud).name 等效
	fmt.Println("Student name:: ", (*pstud).name)
	fmt.Println(student{rollNo: 2, name: "Ann"})
	fmt.Println(student{name: "Ann", rollNo: 2})
	fmt.Println(student{name: "Alice"})

	r := Rect{width: 10, height: 10}
	fmt.Println("Area:: ", r.Area())
	fmt.Println("Perimeter:: ", r.Perimeter())

	ptr := &Rect{width:10, height:5}
	fmt.Println("Area:: ", ptr.Area())    // 与 (*prt).Area() 等效
	fmt.Println("Perimeter:: ", (*ptr).Perimeter())

	var data MyInt = 1
	fmt.Println("Value before increment1() call:: ", data)
	data.increment1()
	fmt.Println("Value after increment1() call:: ", data)
	data.increment2()
	fmt.Println("Value after increment2() call:: ", data)
}