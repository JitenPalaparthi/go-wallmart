package mycircle // can give a different name from dir name but generally we dont give like that

const PI = 3.14

var pi = 3.14

type Circle float32

func New(r float32) Circle {
	return Circle(r)
}

func (c Circle) Area() float64 {
	return float64(PI * c * c)
}

func (c Circle) Perimeter() float64 {
	return float64(2 * pi * float64(c))
}

func (c Circle) What() string {
	return "Circle"
}
