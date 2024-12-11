package shape

import (
	"errors"
	"fmt"
)

type IShape interface {
	Area() float64
	Perimeter() float64
	iWhat
}

type iWhat interface { // unexported or restricted
	What() string
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

func why() {
	fmt.Println("This is to call various shapes")
}
