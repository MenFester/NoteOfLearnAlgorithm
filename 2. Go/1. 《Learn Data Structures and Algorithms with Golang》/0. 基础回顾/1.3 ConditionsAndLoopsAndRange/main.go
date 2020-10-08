package main

import (
	"fmt"
	"reflect"
)

func more(x, y int) bool {
	if x > y { // 不带else的if语句
		return true
	}
	return false
}

func max(x, y int) int {
	var max int
	if x > y { // 带else的if语句
		max = x
	} else {
		max = y
	}
	return max
}

func maxAreaCheck(length, width, limit int) bool {
	if area := length * width; area < limit {
		return true
	} else {
		return false
	}
}

func isEven(value int) {
	switch {
	case value%2 == 0:
		fmt.Println("I is even.")
	default:
		fmt.Println("I is odd.")
	}
}

func main() {
	fmt.Println("1 > 2:: ", more(1, 2))
	fmt.Println("max of 20 and 10 is:: ", max(20, 10))
	fmt.Println("area of 23*34 is lower than limit 330::", maxAreaCheck(23, 34, 330))

	/*
	       // switch has three forms:

	   	switch <initialization >; <condition >{    // switch后跟着是表达式的， case后就必须跟结果值
	   		case <value1>:
	   			<statements>
	   		case <value2>:
	   			<statements>
	   		default:    // optional
	   		    <statements>
	   	}

	   	switch {    // switch后没有表达式的，case后可以放表达式
	   	    case <condition>:
	   	    	<statements>
	   	    case <condition>:
	   	    	<statements>
	   	    default:
	   	    	<statements>
	   	}

	   	switch var.(type) {    // v, ok := var.(type) 是类型断言语法
	   	    case <type>:
	   	    	<statements>
	   	    case <type>:
	   	    	<statements>
	   	    default:
	   	    	<statements>
	   	}
	*/

	i := 2
	switch i {
	case 1, 2, 3:
		fmt.Println("One, Two or Three")
	default:
		fmt.Println("Something else")
	}

	isEven(2344)

	/*
	    // for loop has four forms:
	    for <initialization>;  <condition>; <increment/decrement> {
	    	<statements>
		}

		for <condition> {    // like a while loop
			<statements>
		}

		for {    // an infinite while loop
			<statements>
		}

		for v := range objects {
			<statements>
		}
	*/
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(reflect.TypeOf(numbers))

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println("Sum is:: ", sum)

	i = 0    // for的初始化语句不造成块级作用域
	n := len(numbers)
	sum = 0

	for {
		sum += numbers[i]
		i++
		if i >= n {
			break
		}
	}
	fmt.Println("Sum is:: ", sum)

	sum = 0
	i = 0
	for i < n {
		sum += numbers[i]
		i++
	}
	fmt.Println("Sum is:: ", sum)

	sum = 0
	for index, val := range numbers {
		sum += val
		fmt.Print("[", index, ",", val, "]")
	}
	fmt.Println("\nSum is:: ", sum)
	kvs := map[int]string{1:"apple", 2:"banana"}
	for k, v := range kvs {
		fmt.Println(k, "->", v)
	}
	str := "Hello, World!"
	for index, c := range str {
		fmt.Print("[", index, ",", string(c), "]")
	}
}
