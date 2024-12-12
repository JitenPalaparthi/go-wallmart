package main

import (
	"fmt"
	"os"
	"sync"
)

//var wg *sync.WaitGroup = new(sync.WaitGroup)

// func init() {
// 	wg = new(sync.WaitGroup)
// }

func main() {
	wg := new(sync.WaitGroup)
	defer fmt.Println("end of main")
	fmt.Println("start of main")
	wg.Add(3)
	//wg.Add(1)
	go greet1(wg)
	//wg.Add(1)
	go func() {
		c := add(10, 20)
		//wg.Add(1)
		fmt.Println("c:", c)
		//wg.Done()
		wg.Done()
	}()

	//wg.Add(1)
	go func() {
		greet2()
		wg.Done()
	}()

	go func(wg *sync.WaitGroup) {
		greet2()
		wg.Done()
	}(wg)

	wg.Wait() // wait untile the counter becomes zero
}

func greet1(wg *sync.WaitGroup) {
	defer println("end of greet")
	println("Hello Walmart minds!-1")
	wg.Done()
	//runtime.Goexit()
}

func greet2() {
	defer println("end of greet")
	println("Hello Walmart minds!-2")
	//runtime.Goexit()
}

func add(a, b int) int {
	return a + b
}
func fatal(msg string) {
	println(msg)
	os.Exit(1)
}

// 1. main is also a go routine
// 2. no goroutine waits for other goroutine  to complete its execution
// 3. cannot guarantee the order of execution of goroutines
// 4. just use a go keyword to run a func or method as a goroutine
