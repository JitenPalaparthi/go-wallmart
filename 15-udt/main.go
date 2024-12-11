package main

import (
	"fmt"
	"reflect"
)

type integer = int // is not a new type

type myint int // myint is a new type
// by creating a new type , can attach methods

func (mi myint) Sq() myint {
	return mi * mi
}

type myint2 myint

func (m myint2) Cube() myint2 {
	return m * m * m
}

type myint3 int

func (mi myint3) ToString() string {
	//return strconv.Itoa(int(mi))
	return fmt.Sprint(mi)
}

func main() {
	var num1 int = 749

	var nummyint myint = 234

	var nummyint2 myint2 = 9

	var nummyint3 myint3 = 4

	fmt.Println("myint variable")
	sq1 := nummyint.Sq()
	str1 := myint3(nummyint).ToString()
	cb1 := myint2(nummyint).Cube()

	fmt.Println("Square of nummyint(myint):", sq1)
	fmt.Println("Cube of nummyint(myint):", cb1)
	fmt.Println("ToString of nummyint(myint):", str1, "Type:", reflect.TypeOf(str1))
	fmt.Println("-------------\n")
	fmt.Println("int variable")

	sq2 := myint(num1).Sq()
	cb2 := myint2(num1).Cube()
	str2 := myint3(num1).ToString()

	fmt.Println("Square of num1(int):", sq2)
	fmt.Println("Cube of num1(int):", cb2)
	fmt.Println("ToString of num1(int):", str2, "Type:", reflect.TypeOf(str2))

	fmt.Println("-------------\n")
	fmt.Println("nummyint2 variable")
	sq3 := myint(nummyint2).Sq()
	cb3 := nummyint2.Cube()
	str3 := myint3(nummyint2).ToString()

	fmt.Println("Square of nummyint2(nummyint2):", sq3)
	fmt.Println("Cube of nummyint2(nummyint2):", cb3)
	fmt.Println("ToString of nummyint2(nummyint2):", str3, "Type:", reflect.TypeOf(str3))

	fmt.Println("-------------\n")
	fmt.Println("nummyint3 variable")
	sq4 := myint(nummyint3).Sq()
	cb4 := myint2(nummyint2).Cube()
	str4 := nummyint3.ToString()

	fmt.Println("Square of nummyint2(nummyint3):", sq4)
	fmt.Println("Cube of nummyint2(nummyint3):", cb4)
	fmt.Println("ToString of nummyint2(nummyint3):", str4, "Type:", reflect.TypeOf(str4))

	fmt.Println("-------------\n")
	fmt.Println("float32 variable")
	var float1 float32 = 12.34

	sq5 := myint(float1).Sq()
	cb5 := myint2(float1).Cube()
	str5 := myint3(float1).ToString()

	fmt.Println("Square of nummyint2(float32):", sq5)
	fmt.Println("Cube of nummyint2(float32):", cb5)
	fmt.Println("ToString of nummyint2(float32):", str5, "Type:", reflect.TypeOf(str5))

	//var ok1 bool = true
	//myint3(ok1).ToString // does not work since bool cannot be casted to int so it also cannot be casted to myint3

}
