package main

import (
	"fmt"
	"reflect"
)

func main() {

	func() {
		fmt.Println("Hello WalMart")
	}() // executor

	c1 := func(a, b int) int {
		return a + b
	}(10, 20)

	fmt.Println("c1:", c1)
	f := func(a, b int) int {
		return a - b
	}
	c2 := f(10, 20)
	fmt.Println("c2:", c2)

	num := 100

	func(n int) {
		n++
	}(num)

	fmt.Println(num)
	func(n *int) {
		*n++
	}(&num)

	fmt.Println(num)

	func() {
		num++
	}()

	fmt.Println(num)

	var f1 func(int, int) int = mul

	c3 := f1(100, 5)
	fmt.Println("c3:", c3)

	f1 = func(i1, i2 int) int {
		return i1 / i2
	}

	c4 := f1(100, 5)
	fmt.Println("c4:", c4)

	map1 := make(map[string]any)

	map1["add"] = func(a, b int) int {
		return a + b
	}

	map1["sub"] = func(a, b int) int {
		return a - b
	}

	map1["mul"] = mul

	map1["div"] = f1

	map1["greet"] = func() {
		fmt.Println("Hello WalMart")
	}

	map1["something"] = 100

	a, b := 10, 20
	for key, val := range map1 {

		// switch val.(type) {
		// case func(int, int) int:
		// 	r := val.(func(int, int) int)(a, b)
		// 	fmt.Println("key:", key, "Result:", r)
		// case func():
		// 	val.(func())()
		// }

		switch f1 := val.(type) {
		case func(int, int) int:
			r := f1(a, b)
			fmt.Println("key:", key, "Result:", r)
		case func():
			f1()
		default:
			fmt.Println("type:", reflect.TypeOf(f1), "is not implemented")
		}

	}

	slice1 := make([]any, 5)

	slice1[0] = func(a, b int) int {
		return a + b
	}
	slice1[1] = func(a, b int) int {
		return a - b
	}
	slice1[2] = mul

	slice1[3] = f1
	slice1[4] = 100

	for _, v := range slice1 {
		switch f1 := v.(type) {
		case func(int, int) int:
			r := f1(a, b)
			fmt.Println("Result:", r)
		case func():
			f1()
		default:
			fmt.Println("type:", reflect.TypeOf(f1), "is not implemented")
		}
	}

	r1 := calc(10, 20, slice1[0].(func(int, int) int))
	fmt.Println("r1:", r1)
	r2 := calc(10, 20, func(i1, i2 int) int {
		return i1 - i2
	})
	fmt.Println("r2:", r2)

	f3 := calc2(10, 20)
	r3 := f3()
	fmt.Println("r3:", r3)
}

func mul(a, b int) int {
	return a * b
}

func calc[T func(int, int) int](a, b int, f1 T) int {
	return f1(a, b)
}

func calc2[T F](a, b int) T {
	return func() int {
		return a + b
	}
}

type F interface {
	func() int
}
