package main

import (
	"fmt"
	"math"
)

func main() {
	var v1 int
	var v2 int
	v1 = 100
	fmt.Println("Value stored in variable v1::", v1)
	fmt.Println("Value stored in variable v2::", v2)

	v3 := 100
	fmt.Println("Value stored in variable v3::", v3)

	const PI = math.Pi
	fmt.Println("Value of math.Pi is ::", PI)
}
