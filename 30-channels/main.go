package main

import "time"

func main() { //mg1
	ch := make(chan int)
	//sig := make(chan bool)
	sig := make(chan struct{})
	go sender(ch, 10)
	go receiver(ch, sig, 10)
	<-sig // main would block until it receives the value from the sig
	//<-ch

	// var e struct{}
	// var e2 empty
}

// type empty struct {

// }

func sender(ch chan<- int, num int) {
	for i := 1; i <= num; i++ {
		time.Sleep(time.Millisecond * 200)
		ch <- i * i
	}
}

func receiver(ch <-chan int, sig chan<- struct{}, num int) {
	for i := 1; i <= num; i++ {
		println(<-ch)
	}
	sig <- struct{}{}
}
