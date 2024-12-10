package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var num1 uint8 = 123

	var num2 uint64 = uint64(num1)

	//num3:=(uint64)num1
	println(num2)

	var ok1 bool = true

	//var num3 uint8 = uint8(ok1)

	var num3 = ToInt(ok1) // conversion
	println(num3)

	var str1 string = "45"

	// var (
	// 	num4 int
	// 	err  error
	// )
	// num4, err = strconv.Atoi(str1)
	num4, err := strconv.Atoi(str1) // string to int
	if err != nil {
		fmt.Println("unable to convert:", err.Error())
	} else {
		fmt.Println("converted from str1 to num4:", num4)
	}

	var str2 = "123.123"

	float1, err := strconv.ParseFloat(str2, 64) // convert from string to float64

	if err != nil {
		fmt.Println("unable to convert:", err.Error())
	} else {
		fmt.Printf("converted from str2 to float1:%.2f\n", float1)
	}

	//var float3 float32 = 123
	var any1 any = 123.123 // float64 value

	//var float2 float64 = float64(any1) // cannot type cast any variable

	var ok2 bool

	ok2, done := any1.(bool)
	if done {
		fmt.Println("ok2 to bool is successfully asserted", ok2)
	} else {
		fmt.Println("could not be asserted")
	}

	fmt.Println("ok2", ok2)

	var float2 float64 = any1.(float64) // type assertion

	fmt.Println("float2", float2)

	// add(true, false)
	// add("Hello", 100)
	// add(false, "Hello World")

	sum1, err := add(10, 20)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("add:", sum1)
	}

	sum1, err = add(true, false)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("add:", sum1)
	}

}

func ToInt(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

// true + false --> ?

//true || false -->

//num1 >= 100 && true -->

func add(a any, b any) (any, error) {
	//return a + b
	//println(a, b)
	if a == nil || b == nil {
		return nil, errors.New("a or b is nil")
	}
	a1, done1 := a.(int)
	b1, done2 := b.(int)

	if done1 && done2 {
		return a1 + b1, nil
	}

	a2, done1 := a.(float64)
	b2, done2 := b.(float64)

	if done1 && done2 {
		return a2 + b2, nil
	}

	// if a1, done1 := a.(uint); !done1 {
	// 	return nil, errors.ErrUnsupported
	// } else {
	// 	return a1, fmt.Errorf("unimpleted type")
	// }

	return nil, fmt.Errorf("unimpleted type")
}

// r1:=add<int>(10,20)

// r1:=add_int(10,20)
// r2:=add<float32>(10,20)

// add_int(a int, b int){
// 	return a+b
// }

// add_flaot32(a float32, b float32){
// 	return a+b
// }

// func T add<T>(a T,a T){
// 	return a+b
// }

// error -> interface
// interface header
// ptr
// type
