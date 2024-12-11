package square

func (r Square) Area() float64 {
	return float64(r * r)
}

func (re Square) Perimeter() float64 {
	return float64(4 * re)
}

func (r Square) What() string {
	return "Square"
}
