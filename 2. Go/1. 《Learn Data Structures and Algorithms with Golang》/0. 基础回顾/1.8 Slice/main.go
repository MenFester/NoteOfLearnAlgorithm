package main

import (
	"fmt"
)

//PrintSlice :
func PrintSlice(data []int) {
	fmt.Printf("%v :: len = %d cap = %d \n", data, len(data), cap(data))
}

func main() {
	var s []int
	for i := 1; i < 17; i++ {
		s = append(s, i)
		PrintSlice(s)
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	PrintSlice(s2)

	s3 := make([]int, 10)
	PrintSlice(s3)

	s4 := make([]int, 0, 10)
	PrintSlice(s4)

	s5 := s2[0:4]
	PrintSlice(s5)

	s6 := s5[2:5]    // 底层对应的是同一个数组
	PrintSlice(s6)
}
