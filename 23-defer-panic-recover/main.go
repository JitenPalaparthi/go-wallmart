package main

import (
	"fmt"
	"os"
)

func main() {
	defer println("end of main-1") // f1
	defer println("end of main-2") // f2

	defer func() { //f3
		println("end of main-3")
	}()

	println("start of main")

	f, err := os.OpenFile("data.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		println(err.Error())
		return
	}
	defer f.Close()

	for i := 1; i <= 100; i++ {
		_, err = f.WriteString("Hello Walmart minds!-->" + fmt.Sprint(i) + "\n")
	}

	if err != nil {
		println(err.Error())
		return
	}

	str1 := "Hello, 世界!"
	//str2 := "你好，世界!"

	// for _, v := range str1 {

	// }

	// for i := 0; i < len(str1); i++ {
	// 	println(str1[i], "->", string(str1[i]), " ")
	// }
	// //println()

	// for i, v := range str1 {
	// 	println("index->", i, "->", string(v), " ")
	// }
	// println()

	for _, v := range str1 {
		defer println(string(v))
	}

	num := 100
	defer func(n int) {
		println("number during defer-->", n)
	}(num)

	defer func() {
		println("number during defer-->", num)
	}()

	num += 1
	println("num ->", num)
}

// defer defers the execution to the end of the caller
// all defers w.r.t the caller are LIFO
/*
f3
f2
f1
*/
