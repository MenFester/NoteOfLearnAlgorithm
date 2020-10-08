package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// there eight important family of primitive data types:
	// Boolean, Byte, Integer, Unsigned Integer, Float, String, Rune, Complex
	var B bool = true
	var I int = -1000 // other variation of integers are: int8, int16, int32, int64
	var j uint = 1000 // other variation of integers are: uint8, uint16, uint32, uint64
	var f float32 = 1.2345

	// Byte are alias of uint8
	// Rune are alias of int32, and are used to store Unicode characters
	fmt.Println(B, I, j, f)

	maxInt8 := math.MaxInt8
	minInt8 := math.MinInt8
	maxInt16 := math.MaxInt16
	minInt16 := math.MinInt16
	maxInt32 := math.MaxInt32
	minInt32 := math.MinInt32
	maxInt64 := math.MaxInt64
	minInt64 := math.MinInt64

	maxUint8 := math.MaxUint8   // 没有 math.MinUint8
	maxUint16 := math.MaxUint16 // 没有 math.MinUint16
	maxUint32 := math.MaxUint32 // 没有 math.MinUint32
	// math.MaxUint64 是一个无类型常量，定义变量时如果没有指定类型，会尝试转化为int
	// maxUint64 := math.MaxUint64 // 没有 math.MinUint64
	var maxUint64 uint64 = math.MaxUint64

	maxFloat32 := math.MaxFloat32 // 没有 math.MinFloat32
	maxFloat64 := math.MaxFloat64 // 没有 math.MinFloat64

	fmt.Println("Range of Int8::", minInt8, maxInt8)
	fmt.Println("Range of int16::", minInt16, maxInt16)
	fmt.Println("Range of int32::", minInt32, maxInt32)
	fmt.Println("Range of int64::", minInt64, maxInt64)
	fmt.Println("Max Uint8::", maxUint8, reflect.TypeOf(maxInt8))    // 输出：255， int；注意：自动推导的不一定是你想要的类型
	fmt.Println("Max Uint16::", maxUint16)
	fmt.Println("Max Uint32::", maxUint32)
	fmt.Println("Max Uint64::", maxUint64, reflect.TypeOf(maxUint64))    // 因为不是自动推导，typeof输出正确类型
	fmt.Println("Max Float32::", maxFloat32)
	fmt.Println("Max Float64::", maxFloat64)

	var s string = "hello, World!"
	fmt.Println(s, s[0])    // s[0] 输出是104
	fmt.Printf("s[0] is:: %c\n", s[0])
	fmt.Println("s[1:4] is:: ", s[1:4], " Type of s[1:4] is:: ", reflect.TypeOf(s[1:4]))
	r := []rune(s)
	r[0] = 'H'    // 使用 "H" 则报错，"H"是string类型
	s2 := string(r)
	fmt.Println("s2 is:: ", s2)
	fmt.Println("len of s2 is:: ", len(s2))
	println("hello" + " linguanqiang")
	println("world" == "hello")
	println("a" < "b")    // Unicode 值的大小比较
}