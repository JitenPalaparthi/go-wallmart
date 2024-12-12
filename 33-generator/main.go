package main

import (
	"fmt"
	"time"
)

func main() { //mg1

	sig1 := make(chan struct{})
	//sig2 := make(chan struct{})

	ch1 := generate(time.Second*1, time.Millisecond*100, "gen1")
	ch2 := generate(time.Second*2, time.Millisecond*200, "gen2")

	go receiver(ch1, ch2, sig1)
	//go receiver(ch2, sig2)

	<-sig1
	//<-sig2
}

func generate(d time.Duration, sd time.Duration, name string) chan string {
	ch := make(chan string)
	timeOut := time.After(d)
	num := 1
	go func() {
	out:
		for {
			select {
			case ch <- fmt.Sprint("sending from ", name, "-->", num):
				time.Sleep(sd)
				num++
			case <-timeOut:
				close(ch)
				break out
			}

		}
	}()

	return ch
}

func receiver(ch1 <-chan string, ch2 <-chan string, sig chan<- struct{}) {
	done1, done2 := false, false
	for {
		if done1 && done2 {
			sig <- struct{}{}
			break
		}
		select {
		case v, ok := <-ch1:
			if ok {
				println(v)
			} else {
				done1 = true
			}

		case v, ok := <-ch2:
			if ok {
				println(v)
			} else {
				done2 = true
			}
		}
	}
}
