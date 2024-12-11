package cuboid

var K int

func New(l, b, h float32) *Cuboid {
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
