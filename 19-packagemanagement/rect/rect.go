package rect

func New(l, b float32) *Rect {
	return &Rect{L: l, B: b} // go compiler writes it as 1.0
}

type Rect struct {
	L, B float32
	a, p float64 // restricted fields cannot be accessed outside of rect package
}

func (r *Rect) Area() float64 {
	r.a = float64(r.L) * float64(r.B)
	return r.a
}

func (re *Rect) Perimeter() float64 {
	re.p = 2 * (float64(re.L) + float64(re.B))
	return re.p
}

func (r *Rect) What() string {
	return "Rect"
}
