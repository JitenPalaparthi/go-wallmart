package main

import "fmt"

func main() {

	//r1 := &Rect{L: 12.2, B: 12.43}

	// r1 := new(Rect) // pointer of Rect
	// r1.L = 12.23
	// r1.B = 12.45

	r1 := New(12.3, 14.3)

	// a1 := Area(r1)
	// p1 := Perimeter(r1)
	a1 := r1.Area()
	p1 := r1.Perimeter()
	fmt.Printf("Area of Rect r1:%.2f\n", a1)
	fmt.Printf("Perimeter of Rect r1:%.2f\n", p1)
	fmt.Println("Area of r1 inside of area A", r1.A)
	fmt.Println("Perimeter of r1 inside of perimeter P", r1.P)

}

func NewDefault() *Rect {
	r := new(Rect) // pointer of Rect
	r.L = 1
	r.B = 1
	return r
}

func New(l, b float32) *Rect {
	return &Rect{L: 1, B: 1} // go compiler writes it as 1.0
}

type Rect struct {
	L, B float32
	// L float32
	// B float32
	// A float32
	// P float32
	A, P float64
}

// LastModifed   -> pascal case
// lastModified  -> camel case
// last_modified -> snake case // dont use in go

func Area(r *Rect) float64 {
	r.A = float64(r.L) * float64(r.B)
	return r.A
}

func Perimeter(r *Rect) float64 {
	r.P = 2 * (float64(r.L) + float64(r.B))
	return r.P
}

func (r *Rect) Area() float64 {
	r.A = float64(r.L) * float64(r.B)
	return r.A
}

func (re *Rect) Perimeter() float64 {
	re.P = 2 * (float64(re.L) + float64(re.B))
	return re.P
}

// non pointer receiver is call/pass by value
// pointer receiver is call/pass by reference
// methods are attached to the type using reeiver
// receiver name can be any valid identifer but but idomatic approach is use single letter as and where possible

type EmptyShape struct {
}

func (e *EmptyShape) Area(l, b float32) float64 {
	return float64(l) * float64(l)
}

func (e EmptyShape) Perimeter(l, b float32) float64 {
	return float64(2 * (l + b))
}
