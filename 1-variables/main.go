package main

import (
	"fmt"
	"unsafe"
)

var global int = 60 // stored in data segment

const (
	PI       = 3.14
	MAX  int = 9999
	MIN      = 60
	SEC      = 60
	HOUR     = 24 * MIN * SEC       // this is an expression
	SOME     = 24*7/HOUR + MIN + PI //24 * global * global
)

func main() {

	var num1 int = 100
	println("num:", num1)

	var num2 = 12331 // go infers the type based on value

	var num3 uint8 // zero value is given based on the type

	var ok1 bool // zero value is false

	var str1 string // zero value is ""

	float1 := 12312.123 // shorthand declaration, float64

	age1 := 39 // int which is 8 bytes on 64bit machine,int (8 bytes)

	var age2 uint8 = 39

	str2 := "Hello World"

	char3 := 'D'

	println(num1, num2, num3, str1, str2, ok1, age1, age2, float1)

	var char1 uint8 = 'A' // single quote,based on the unicode value
	var char2 rune = 'B'  // rune is 4 bytes it is alias of int32
	println(char1, char2, char3)

	var (
		num4 uint8 = 123
		num5       = 12312
		ok2  bool
		ok3  = true
	)

	println(num4, num5, ok2, ok3)

	var any1 any // interface{} --> empty interface
	fmt.Println(any1)

	if any1 == nil {
		println("no value is assigned to any1")
	}

	var str3 string = "Hello World"
	var str4 string = "Hello Wallmart world!"

	var str5 string // "", internally it is not nil

	fmt.Println("Size of any1:", unsafe.Sizeof(any1))
	fmt.Println("Size of str3:", unsafe.Sizeof(str3))
	fmt.Println("Size of str4:", unsafe.Sizeof(str4))
	fmt.Println("Size of str5:", unsafe.Sizeof(str5))

	any1 = true // boolean

	any1 = 12312.12312 // float64

	any1 = "Hello World" // string

	var any2 interface{} = any1
	fmt.Println("Size of any2:", unsafe.Sizeof(any2))

}

// uint,int, int8-in64,uint8-uint64,flaot32,flaot64, rune, byte
// type inference
// zero value (kind of a default value)

// code/text segment -> constant
// data segment -> global variables
// stack memory
// heap memory

// zero value of any is nil

// Tony Hoare -> csp , erlang, golang
// Tony Hoare -> Null pointer for c++

// string header
// ptr -> 8 bytes
// len -> 8 bytes

// any header
// ptr -> 8 bytes // pointer to the actual address it is stored
// type -> 8 bytes (pointer) // the type of the value that is assigned
