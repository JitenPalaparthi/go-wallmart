package main

import (
	"fmt"
	"time"
)

func main() { //mg1
	ch1 := make(chan string)
	sig1 := make(chan struct{})
	//timeOutSig := make(chan struct{})

	//t := time.After(time.Second * 1)

	//go timeOut(timeOutSig, time.Second*1)
	//go sender1(ch1, timeOutSig)
	//go sender2(ch1, time.Second*2)
	go sender3(ch1, time.Second*2)
	go receiver(ch1, sig1)

	<-sig1
}

func sender1(ch chan<- string, timeOutSig <-chan struct{}) {
	num := 1
out:
	for {
		select {
		case ch <- fmt.Sprint("sender -->", num*num):
			time.Sleep(time.Millisecond * 200)
			num++
		case <-timeOutSig:
			close(ch)
			break out
		}
	}
}

func sender2(ch chan<- string, d time.Duration) {
	num := 1
	timeOutch := time.After(d)
out:
	for {
		select {
		case ch <- fmt.Sprint("sender -->", num*num):
			time.Sleep(time.Millisecond * 200)
			num++
		case <-timeOutch:
			close(ch)
			break out
		}
	}
}

func sender3(ch chan<- string, d time.Duration) {
	num := 1
	timeOutch := timeOutR(d)
out:
	for {
		select {
		case ch <- fmt.Sprint("sender -->", num*num):
			time.Sleep(time.Millisecond * 200)
			num++
		case <-timeOutch:
			close(ch)
			break out
		}
	}
}

func receiver(ch <-chan string, sig chan<- struct{}) {
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

func timeOut(sig chan<- struct{}, d time.Duration) {
	time.Sleep(d)
	sig <- struct{}{}
}

func timeOutR(d time.Duration) chan struct{} {
	timeOutSig := make(chan struct{})
	go func() {
		time.Sleep(d)
		timeOutSig <- struct{}{}
		close(timeOutSig)
	}()
	return timeOutSig
}
