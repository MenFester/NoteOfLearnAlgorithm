package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//IncrementPassByValue :
// parameter passing, call by value
func IncrementPassByValue(x int) {
	x++
}

//IncrementPassByPointer
// parameter passing, call by Pointer/Reference
func IncrementPassByPointer(ptr *int) {
	(*ptr)++
}

func main() {
	fmt.Println(max(10, 20))

	i := 10
	fmt.Println("Value of i before increment is:: ", i)
	IncrementPassByValue(i)
	fmt.Println("Value of i after increment is:: ", i)

	data := 10
	ptr := &data
	fmt.Println("Value stored at variable data is:: ", data)
	fmt.Println("Value stored at variable *ptr is:: ", *ptr)
	fmt.Println("The address of variable data is:: ", &data)
	fmt.Println("The address of variable stored in ptr is:: ", ptr)

	j := 10
	fmt.Println("Value of i before increment is:: ", j)
	IncrementPassByPointer(&j)
	fmt.Println("Value of i after increment is:: ", j)

}