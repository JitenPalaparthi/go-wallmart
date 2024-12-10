package main

import (
	"fmt"
)

func main() {
	//slice1 := make([]int, 10) // [0 0 0 0 0 0 0 0 0 0]
	//var slice1 []int                    // declaring but not instantiated, means slice1 is nil
	//slice1 = append(slice1, 10, 20, 30) // instantiate and add new elements to the slice
	// for i := range slice1 {
	// 	slice1[i] = v2.IntN(100)
	// }
	// fmt.Println(slice1)

	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := arr1[:]

	slice2 := slice1
	slice3 := slice1[0:]
	slice4 := slice1[:5]  // but not 5 0-4 elements
	slice5 := slice1[2:8] // but not 8th
	slice6 := slice1[5:]  // 5th element to the end

	fmt.Println("arr1:  ", arr1)
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
	fmt.Println("slice4:", slice4)
	fmt.Println("slice5:", slice5)
	fmt.Println("slice6:", slice6)

	fmt.Println("Sum", sumOf())
	fmt.Println("Sum", sumOf(1, 2))
	fmt.Println("Sum", sumOf(slice1...))
	fmt.Println("Sum", sumOf(arr1[:]...))

	slice7 := append(slice1[:4], slice1[5:]...)
	fmt.Println("slice7:", slice7)

	slice8 := make([]int, 10)
	copy(slice8, slice1) // deep copy
	fmt.Println("slice8:", slice8)
	clear(slice7)
	fmt.Println("slice7:", slice7)
	fmt.Println("arr1:", arr1)
	fmt.Println("slice8:", slice8)

}

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

func sumOf(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

// write a function to delete elements from a slice and return the slice
// think of pros and cons. and also return error as well
