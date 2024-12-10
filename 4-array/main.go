package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
)

func main() {
	var arr1 [5]int // zero values . What is the type of a varible?
	// the type of an array includes its length as well
	arr1[0] = 100
	arr1[1] = 101
	arr1[2] = 102
	fmt.Println(arr1, "type of ", reflect.TypeOf(arr1))
	arr2 := [8]int{10, 123, 123, 23, 23, 2, 32, 23} // shorthand declaration
	arr3 := [...]bool{true, false, true, true, true, true, false, false}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	sum1 := sumofArr1(arr1)
	fmt.Println("sum1:", sum1)
	arr2 = assignArrElem(arr2)
	sum2 := SumofArr2(arr2)
	fmt.Println("sum1:", sum2)

	arr2d := [2][2]int{{1, 2}, {3, 4}}

	//arr3d := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	for i := 0; i < len(arr2d); i++ {
		for j := 0; j < len(arr2d[i]); j++ {
			fmt.Println("i:", i, "j:", j, "-->", arr2d[i][j])
		}
	}

	//arrany := [5]any{100, 10.12, true, "hello World", arr1}

	fmt.Println("address of an array:", &arr1[0])
}

// string header
// ptr
// len

// any header
// ptr
// type_ptr

func sumofArr1(arr [5]int) int {
	sum := 0 // sum is an int, it is a shorthand declaration
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func assignArrElem(arr [8]int) [8]int {
	for i, _ := range arr {
		arr[i] = rand.IntN(1000)
	}
	return arr
}

func SumofArr2(arr [8]int) int {
	sum := 0 // sum is an int, it is a shorthand declaration
	for _, v := range arr {
		sum += v
	}
	return sum
}

// range loop on array, slice --> index and value

// length is known at the compile time
// length of an array is fixed, cannot be changed during runtime
