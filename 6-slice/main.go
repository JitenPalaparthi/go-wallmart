package main

import (
	"fmt"
	v2 "math/rand/v2"
)

func main() {
	// var str1 string //&str1 ptr=0x11 len=0
	// var any1 any    // &any1 ptr=0x00 type=0x00 // 0x00 is nil

	//str1 = "Hello world"
	// any1 = "hello world"
	// if any1 == nil {

	// }
	// if slice1 == nil {

	// }

	slice1 := make([]int, 5, 10)

	// for i, _ := range slice1 {
	// 	slice1[i] = v2.IntN(100)
	// }
	for i := range slice1 {
		slice1[i] = v2.IntN(100)
	}

	PrintSliceInfo(slice1)
	slice2 := slice1 // headers are copied, kind of a shallow copy but not a deep copy
	PrintSliceInfo(slice2)
	slice2[0] = 99999
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	slice2 = append(slice2, 10000)
	slice2[1] = 111111
	fmt.Println("after append")
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	PrintSliceInfo(slice1)
	PrintSliceInfo(slice2)
	slice2[5] = 88888
	//Sq(slice1)
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	slice1 = SqElem(slice1, 11, 12, 13, 14, 15, 16, 17, 19) // coping the header
	slice2 = slice1                                         // coping header
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2} // evaluated at compile time
	//slice3 := []int{1, 2, 3, 4} // header is created ptr, len 4 cap 4

	sum1 := SumOf(arr2[:])
	sum2 := SumOf(arr1[:])

	fmt.Println("sum1:", sum1)
	fmt.Println("sum2:", sum2)

}

// string header
// ptr -> address of the 0th element
// len -> how many elements

// any header
// ptr -> address of the 0th element
// type -> actual type (pointer of the type)

// slice header
// ptr -> address of the 0th element
// len -> length of the elements
// cap -> cap of the slice

func PrintSliceInfo(slice []int) {
	fmt.Println("-------------------")
	fmt.Printf("Address of the slice:%p\n", &slice)
	if slice != nil {
		if len(slice) > 0 {
			fmt.Printf("Ptr of the slice:%p\n", &slice[0])
		}
		fmt.Printf("Len of the slice:%d\n", len(slice))
		fmt.Printf("Cap of the slice:%d\n", cap(slice))
		fmt.Println("slice:", slice)

	} else {
		fmt.Println("nil slice")
	}
	fmt.Println("-------------------")
}

func Sq(slice []int) {
	for i, v := range slice {
		slice[i] = v * v
	}
}

func SumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

// variadic must be the last paramer in a func
// variadic can only be used in funcs and methods
// can use range loop to iterate variadia variable
func SqElem(slice []int, nums ...int) []int {
	slice = append(slice, nums...)
	for i, v := range slice {
		slice[i] = v * v
	}
	return slice
}
