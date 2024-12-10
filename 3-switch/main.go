package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	var day uint8 = 4

	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Fridayday")
	case 7:
		println("Saturday")
	default:
		println("no day")
	}

	var num int = 1232

	switch { // empty switch
	case num%2 == 0:
		println(num, "is even")
	case num%2 == 1:
		println(num, "is odd")
	}

	char := '你'

	switch char {
	case 'A', 'E', 'I', 'O', 'U':
		println(string(char), "is capital vovel")
	case 'a', 'e', 'i', 'o', 'u':
		println(string(char), "is lower vovel")
	default:
		println(string(char), "is either a consonent or a special unicode char")
	}

	num = 48

	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

	println("example on false negative if order of cases are written in different way")
	num = 2

	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	}

	println("type switch")

	var any1 any = "hello world"

	switch any1.(type) {
	case int, uint, uint8, uint16, uint32, uint64, float32, float64, int8, int32, int64:
		fmt.Println("number type.any1 value", any1, "type:", reflect.TypeOf(any1))
	case nil:
		fmt.Println("nil type")
	default:
		fmt.Printf("any1 does not hold a number.It holds a type of %T\n", any1)
	}
	//str1 := "你好，世界"

	sum1, err := add(10, 20)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum1:", sum1)
	}
	sum2, err := add(10.123, 20.123)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("sum2:", sum2)
	}

	sum3, err := add(true, true)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("sum3:", sum3)
	}

	var a1, b1 uint8 = 12, 13
	sum4, err := add(a1, b1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("sum4:", sum4)
	}

	sum5, err := add("hello ", "World")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("sum5:", sum5)
	}

	sum6, err := add(10, 10.12)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("sum6:", sum6)
	}

	sum7 := addAG(12, 13)
	fmt.Println("sum7:", sum7)
	sum8 := addAG(12.123, 13.123)
	fmt.Println("sum8:", sum8)
	sum9 := addAG(12, 13.123)
	fmt.Println("sum9:", sum9)

	// sum10 := addAG(12, true)
	// fmt.Println("sum10:", sum10)

}

func add(a, b any) (any, error) {
	if a == nil || b == nil {
		return nil, errors.New("a or b is nil")
	} else if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return nil, errors.ErrUnsupported
	}
	var sum float64
	switch a.(type) {
	case int:
		sum = float64(a.(int) + b.(int))
	case uint:
		sum = float64(a.(uint) + b.(uint))
	case uint8:
		sum = float64(a.(uint8) + b.(uint8))
	case uint32:
		sum = float64(a.(uint32) + b.(uint32))
	case uint64:
		sum = float64(a.(uint64) + b.(uint64))
	case int8:
		sum = float64(a.(int8) + b.(int8))
	case int16:
		sum = float64(a.(int16) + b.(int16))
	case int32:
		sum = float64(a.(int32) + b.(int32))
	case int64:
		sum = float64(a.(int64) + b.(int64))
	case float32:
		sum = float64(a.(float32) + b.(float32))
	case float64:
		sum = a.(float64) + b.(float64)
	default:
		return nil, errors.ErrUnsupported
	}
	return sum, nil
}

func addG[T int | uint | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64](a, b T) T { // generics
	return a + b
}

type ArthTypes interface { // this interface is only for types
	int | uint | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

func addAG[T ArthTypes](a, b T) T { // generics
	return a + b
}

// go source -->
// --> compiled(constant evalues, dead code elimation, optimizatons,generic code monomorphization,validation)
// --> SSA (Static Single Assignment)
// --> IR code (not emitted) [in java it emits the java bytecode]
// --> Based on GOARCH GOOS (env variables)
// --> machine code is generatged by assembler->.o files
// --> the linker takes this .o file and links all required object files, including rutime stuff
// --> linker generates the binary file
