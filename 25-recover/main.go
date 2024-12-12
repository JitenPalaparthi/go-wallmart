package main

import (
	"errors"
	"fmt"
)

func main() { // func0
	defer println("end of main")
	func() { // func1
		//num = 100
		defer println("end func1")
		func() { // func1.func2
			defer println("end func1.func2")
			defer recoverme()
			f1 := func(a, b int) error {
				if a == 0 || b == 0 {
					return errors.New("a or b is zero")
				}
				println("anonymous function", a/b)
				return nil
			}

			if err := f1(20, 0); err != nil { // give a or b  zero to check panic
				panic(err)
			}
		}()

		num := 10

		for i := num; ; i-- {
			if i%2 == 1 {
				continue
			}
			if i != 0 { // comment and uncomment based on panics, to check
				println(100 / i)
			}
			if i == 0 {
				break
			}
		}
	}()
	println("This is executed directly in main")
}

func recoverme() {
	rec := recover()
	if rec != nil {
		fmt.Println("This recovers from the cascading panic", rec)
	}
}
