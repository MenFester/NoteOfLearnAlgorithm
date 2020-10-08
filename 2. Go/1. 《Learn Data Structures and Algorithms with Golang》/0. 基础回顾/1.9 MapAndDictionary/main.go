package main

import (
	"fmt"
)

// a map is a collection of Key-Value pairs.
// Hash-Table is used to store elements in a Map so it is unordered
// Maps have to be initialized using make() before they can be used
// var <variable> map[<key DataType>]<value DataType> = make(map[<key DataType>]<value DataType>)
// or <variabel> := make(map[<key DataType>]<value DataType>)

func main() {
	m := make(map[string]int)
	m["Apple"] = 40
	m["Banana"] = 30
	m["Mango"] = 50

	for key, val := range m {
		fmt.Println("[", key, "->", val, "]")
	}

	fmt.Println("Apple price:: ", m["Apple"])
	delete(m, "Apple")    // delete map items with key

	fmt.Println(m)

	value, ok := m["Apple"]
	fmt.Println("Apple price:: ", value, "Present:: ", ok)

	value2, ok2 := m["Banana"]
	fmt.Println("Banana price:: ", value2, "Present:: ", ok2)
}