package main

import (
	"fmt"
)

func main() {
	var num IEmpty = 100
	var n int = num.(int)
	fmt.Println(n)

	c, err := NewCalc(10)
	if err != nil {
		fmt.Println(err)
	} else {
		v := c.Add(10).Sub(10).Add(5).Mul(5).Div(5).Get()
		fmt.Println("Data in c:", v)
	}

}

type ICalc interface {
	Add(num int) ICalc
	Sub(num int) ICalc
	Mul(num int) ICalc
	Div(num int) ICalc
	Get() int
}

type IEmpty interface{} /// every variable or object satisfies this

// ICalc, IEmpty and any

// What happen when an inteface is created?-> IDT (Interface Definition Table--> similar to VTables)
// What is the size of an interface? -> 16 bytes (ptr and type-ptr)
// Is interface static or dynamic dispatch? -> dynamic dispatch

func NewCalc(d int) (*Calc, error) {
	if d == 0 {
		//return nil, errors.New("invalid calc")

		return nil, &CalcError{data: d}
	}
	return &Calc{data: d}, nil
}

type CalcError struct {
	data int
}

func (ce *CalcError) Error() string {
	return fmt.Sprintf("invalid data %d is given", ce.data)
}

type Calc struct {
	data int
}

func (c *Calc) Add(num int) ICalc {
	c.data += num
	return c
}

func (c *Calc) Sub(num int) ICalc {
	c.data -= num
	return c
}

func (c *Calc) Mul(num int) ICalc {
	c.data *= num
	return c
}

func (c *Calc) Div(num int) ICalc {
	c.data /= num
	return c
}

func (c *Calc) Get() int {
	return c.data
}
