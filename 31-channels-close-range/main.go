package main

import (
	"fmt"
	"time"
)

func main() { //mg1
	ch1, ch2 := make(chan string), make(chan string)
	//sig := make(chan bool)
	sig1, sig2 := make(chan struct{}), make(chan struct{})
	go sender1(ch1, 10)
	go receiver1(ch1, sig1)

	go sender2(ch2, 10)
	go receiver2(ch2, sig2)
	<-sig1 // main would block until it receives the value from the sig
	<-sig2
}

// type empty struct {

// }

func sender1(ch chan<- string, num int) {
	for i := 1; i <= num; i++ {
		time.Sleep(time.Millisecond * 200)
		ch <- fmt.Sprint("sender-1 -->", i*i)
	}
	close(ch) // closing channel is declaring that the channel finished sending data
	// it does not make the ch nil
}

func sender2(ch chan<- string, num int) {
	for i := 1; i <= num; i++ {
		time.Sleep(time.Millisecond * 200)
		ch <- fmt.Sprint("sender-2 -->", i*i)
	}
	close(ch) // closing channel is declaring that the channel finished sending data
	// it does not make the ch nil
}

func receiver1(ch <-chan string, sig chan<- struct{}) {
	for {
		v, ok := <-ch
		if ok {
			println(v)
		} else {
			sig <- struct{}{}
			//runtime.Goexit()
			//return
			break
		}
	}
}

func receiver2(ch <-chan string, sig chan<- struct{}) {
	for v := range ch { // range iterates until the channel is closed
		println(v)
	}
	sig <- struct{}{}
}
