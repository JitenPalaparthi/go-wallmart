package main

import "fmt"

func main() {
	var slice1 []int // This is a declaration of slice, not an instantiation
	if slice1 == nil {
		fmt.Println("yes this slice is nil")
	}
	PrintSliceInfo(slice1)
	slice1 = make([]int, 5, 10)
	PrintSliceInfo(slice1)
	slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 100, 101, 102, 103, 104
	//a, b, c := 10, 20, 30
	//a, b, c = b, c, a
	// t := b
	// b = a
	// a = t
	PrintSliceInfo(slice1)

	slice1 = append(slice1, 10, 20, 30) // a header is copied
	PrintSliceInfo(slice1)

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

// slice you dont give the length
// slice can be nil
// slice header
// ptr - 8
// len - 8
// cap - 8
