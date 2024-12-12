package main

import (
	"fmt"
	"time"
)

func main() { //mg1
	var ch chan int     // dclaring a channel
	ch = make(chan int) // instantiating a  unbuffered channel can hold only one int data

	//ch <- 100 this blocks the main
	go func() { //g1
		ch <- 100
	}() // sending data to the channel
	v := <-ch // receiving data from the channel
	fmt.Println(v)

	go func() {
		time.Sleep(time.Second * 2)
		v := <-ch // receiving data from the channel
		fmt.Println(v)
	}()

	ch <- 200
	fmt.Println("Done?")

	time.Sleep(time.Second * 3)

}
