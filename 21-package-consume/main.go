package main

import (
	"fmt"

	"github.com/jitenpalaparthi/walmart-shapes/shape"
	mycircle "github.com/jitenpalaparthi/walmart-shapes/shape/circle"
	"github.com/jitenpalaparthi/walmart-shapes/shape/cuboid"
	"github.com/jitenpalaparthi/walmart-shapes/shape/rect"
	"github.com/jitenpalaparthi/walmart-shapes/shape/square"
)

func main() {

	shapes := make([]shape.IShape, 9)

	shapes[0] = rect.New(10.23, 14.13)
	shapes[1] = rect.New(14.2, 14.23)
	shapes[2] = square.New(12.24)
	shapes[3] = square.New(123.123)
	shapes[4] = cuboid.New(12.12, 13.45, 13.45)
	shapes[5] = cuboid.New(12.49, 13.123, 3.45)
	shapes[6] = nil
	shapes[7] = mycircle.New(10.12)
	shapes[8] = mycircle.New(23.12)
	fmt.Println("PI constant value:", mycircle.PI)
	//fmt.Println("pi variable value:", mycircle.pi) // cant access due to restricted natured

	for _, v := range shapes {
		if err := shape.Shape(v); err != nil {
			fmt.Println(err)
		}
	}

}
