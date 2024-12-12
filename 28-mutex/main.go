package main

import "sync"

var num = 0

func main() {

	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go increment(wg, mu)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go decrement(wg, mu)
		}
		wg.Done()
	}()

	wg.Wait()
	println(num)
}

func increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	num++
	mu.Unlock()
	wg.Done()
}

func decrement(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	num--
	mu.Unlock()
	wg.Done()
}

type SomeDataTSafe struct {
	num int
	mu  sync.Mutex
}

func (s *SomeDataTSafe) increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	s.mu.Lock()
	num++
	s.mu.Unlock()
	wg.Done()
}

func (s *SomeDataTSafe) decrement(wg *sync.WaitGroup, mu *sync.Mutex) {
	s.mu.Lock()
	num--
	s.mu.Unlock()
	wg.Done()
}

func (s *SomeDataTSafe) print() {
	println(num)
}

type SomeDataNotTSafe struct {
	num int
}

func (s *SomeDataNotTSafe) increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	num++
	wg.Done()
}

func (s *SomeDataNotTSafe) decrement(wg *sync.WaitGroup, mu *sync.Mutex) {
	num--
	wg.Done()
}

func (s *SomeDataNotTSafe) print() {
	println(num)
}
