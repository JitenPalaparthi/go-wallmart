package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	defer fmt.Println("end of main")
	fmt.Println("start of main")
	go greet()
	go func() {
		c := add(10, 20)
		go fmt.Println("c:", c)
	}()
	//os.Exit(1)
	//time.Sleep(time.Millisecond * 1) // this is not a solution

	num := 0

out:
	num++
	time.Sleep(time.Second * 1)
	if num <= 3 {
		goto out
	} else {
		goto errors
	}
errors:
	err := errors.New("timedout")
	if err != nil {
		println(err.Error())
		runtime.Goexit() // not a solution
	}
	//log.Fatal()

	// if err := http.ListenAndServe(":8081", nil); err != nil {
	// 	runtime.Goexit()
	// }
}

func greet() {
	defer println("end of greet")
	println("Hello Walmart minds!")
	runtime.Goexit()
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
