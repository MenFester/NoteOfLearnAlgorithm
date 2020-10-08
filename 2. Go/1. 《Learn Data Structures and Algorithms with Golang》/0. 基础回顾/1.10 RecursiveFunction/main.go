package main

import (
	"fmt"
)

//Factorial :
// N! = N*(N-1)*...*2*1
func Factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * Factorial(i-1)
}

func main() {
	fmt.Println("factorial 5 is:: ", Factorial(5))
}