package main

import (
	"fmt"

	mycircle "github.com/jitenpalaparthi/walmart-shapes/circle"
	cbd "github.com/jitenpalaparthi/walmart-shapes/cuboid"
	"github.com/jitenpalaparthi/walmart-shapes/rect"
	"github.com/jitenpalaparthi/walmart-shapes/shape"
	"github.com/jitenpalaparthi/walmart-shapes/square"
)

func init() {
	fmt.Println("Init is called before main")
}

func init() {
	fmt.Println("Init is called when the pacakge is imported")
}

func init() {
	fmt.Println("cannot order init directly from differnet pacakges")
}

var num int = 10

var Global int = func(n int) int {
	return n * n
}(num)

var map1 map[string]any // this gets stored in data segment, this is similar to static variable

func init() {
	map1 = make(map[string]any) // what you do with it is differnet but it gets intantiated in init function
	fmt.Println("main is not yet called , still the function is evaluated as -->", Global)
}

func main() {
	Greet()

	shapes := make([]shape.IShape, 9)

	shapes[0] = rect.New(10.23, 14.13)
	shapes[1] = rect.New(14.2, 14.23)
	shapes[2] = square.New(12.24)
	shapes[3] = square.New(123.123)
	shapes[4] = cbd.New(12.12, 13.45, 13.45)
	shapes[5] = cbd.New(12.49, 13.123, 3.45)
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

// binary application
// pkg or lib
// what is the difference? binary conatins main func from main pacakge
