package main

import (
	"fmt"
)

//SumArray :
// Sum List
func SumArray(data []int) int {    // the parameter's type is slice
	size := len(data)
	total := 0
	for index := 0; index < size; index++ {
		total = total + data[index]
	}

	return total
}

//SequentialSearch :
// Sequential Search, which will search a list for some given value
func SequentialSearch(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}

	return false
}

//BinarySearch :
// Binary search in a sorted list
func BinarySearch(data []int, value int) bool {
	size := len(data)
	var mid int
	low := 0
	high := size - 1

	for low <= high {
		mid = low + (high - low) / 2
		if data[mid] == value {
			return true
		} else {
			if data[mid] < value {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}

//RotateArray :
// Rotating a list by K positions
func RotateArray(data []int, k int) {
	n := len(data)
	ReverseArray(data, 0, k-1)
	ReverseArray(data, k, n-1)
	ReverseArray(data, 0, n-1)

}

//ReverseArray :
func ReverseArray(data []int, start int, end int) {
	i := start
	j := end
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

//MaxSubArraySum :
// Find a contiguous subarray whose sum(sum of elements) is maximum
func MaxSubArraySum(data []int) int {
	size := len(data)
	maxSoFar := 0
	maxEndingHere := 0

	for i := 0; i < size; i++ {
		maxEndingHere = maxEndingHere + data[i]
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}

	return maxSoFar
}

func main() {
	// An array is a collection of variables of the same data
	var arr [10]int
	fmt.Println(arr)

	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)

	count := len(arr)
	fmt.Println("Length of array:: ", count)

	fmt.Println("Sum of array:: ", SumArray(arr[:]))

	fmt.Println("4 in the array?:: ", SequentialSearch(arr[:], 4))

	fmt.Println("6 in the array?:: ", BinarySearch(arr[:], 6))

	rArr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rArr)
	RotateArray(rArr, 3)
	fmt.Println(rArr) 

	mArr := []int{1, -2, 3, 4, -4, 6, -14, 8, 2}
	fmt.Println(MaxSubArraySum(mArr))
}
