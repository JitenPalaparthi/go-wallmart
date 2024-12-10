package main

import "unsafe"

func main() {
	var num int = 100
	var ptr *int = &num

	var ptr1 *int = new(int) // allocates memory for int, assign value 0 initially and returns the address
	*ptr1 = 300

	var ptr4 **int = &ptr1

	var ptr5 ***int = &ptr4
	***ptr5 = 9999

	println("ptr1", *ptr1)

	*ptr = 200 // dereference
	println("num:", num)

	*&num = 300
	println("num:", num)

	num = 400
	println("num:", num)
	Incr(num)
	println("num:", num) // 400

	IncrP(&num)
	println("num:", num) // 400

	IncrP(ptr)
	println("num:", num) // 400

	ptr2 := Sq1(num)
	println("ptr2:", *ptr2)

	//ptr++ // not allowed , becase it is pointer arthemetic

	arr1 := [5]int{99, 88, 77, 66, 55}

	// var uintptr1 uintptr = uintptr(unsafe.Pointer(&arr1[0]))
	// uintptr1 += unsafe.Sizeof(arr1[0])
	// v1 := *(*int)(unsafe.Pointer(uintptr1))
	// println(v1)
	var uintptr1 uintptr = uintptr(unsafe.Pointer(&arr1[0]))
	for i := 0; i < len(arr1); i++ {
		v1 := (*int)(unsafe.Pointer(uintptr1))
		println(*v1)
		uintptr1 += unsafe.Sizeof(arr1[i])
	}
}

/*
&01 09 f2
Sliceheader{
ptr
len
cap
}*/

func Incr(n int) { // pass by value
	n++
}

func IncrP(p *int) { // pass by ref
	*p++
}

func Sq1(n int) *int { // dangling pointer
	n = n * n
	return &n
}

func Sq2(n int) *int { // dangling pointer
	ptr := new(int) // new allocates memory for int
	*ptr = n * n    // the zero value of ptr is not nil and *ptr zero value is zero
	return ptr
}

func Sq3(n int) *int { // dangling pointer
	var ptr *int
	ptr = &n
	*ptr = n * n
	return ptr
}

// check whether the ptr is nil
