package main

import (
	"errors"
	"fmt"
)

func main() {

	shapes := make([]IShape, 7)
	shapes[0] = NewRect(10.23, 14.13)
	shapes[1] = NewRect(14.2, 14.23)
	shapes[2] = NewSquare(12.24)
	shapes[3] = NewSquare(123.123)
	shapes[4] = NewCuboid(12.12, 13.45, 13.45)
	shapes[5] = NewCuboid(12.49, 13.123, 3.45)
	shapes[6] = nil

	for _, v := range shapes {
		if err := Shape(v); err != nil {
			fmt.Println(err)
		}
	}

}

func Shape(ishape IShape) error {
	if ishape != nil {
		fmt.Println("-----------------")
		fmt.Printf("Area of %s: %.2f\n", ishape.What(), ishape.Area())
		fmt.Printf("Perimeter of %s: %.2f\n", ishape.What(), ishape.Perimeter())
		fmt.Println("-----------------")
	} else {
		return errors.New("nil shape")
	}
	return nil
}

type IShape interface {
	Area() float64
	Perimeter() float64
	IWhat
}

type IWhat interface {
	What() string
}

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b} // go compiler writes it as 1.0
}

type Rect struct {
	L, B float32
	A, P float64
}

func (r *Rect) Area() float64 {
	r.A = float64(r.L) * float64(r.B)
	return r.A
}

func (re *Rect) Perimeter() float64 {
	re.P = 2 * (float64(re.L) + float64(re.B))
	return re.P
}

func (r *Rect) What() string {
	return "Rect"
}

func NewSquare(s float32) Square {
	return Square(s)
}

type Square float32

func (r Square) Area() float64 {
	return float64(r * r)
}

func (re Square) Perimeter() float64 {
	return float64(4 * re)
}

func (r Square) What() string {
	return "Square"
}

func NewCuboid(l, b, h float32) *Cuboid {
	return &Cuboid{L: l, B: b, H: h} // go compiler writes it as 1.0
}

type Cuboid struct {
	L, B, H float32
	A, P    float64
}

func (r *Cuboid) Area() float64 {
	r.A = float64(r.L) * float64(r.B) * float64(r.H)
	return r.A
}

func (re *Cuboid) Perimeter() float64 {
	re.P = 2 * (float64(re.L) + float64(re.B) + float64(re.H))
	return re.P
}

func (r *Cuboid) What() string {
	return "Cuboid"
}
