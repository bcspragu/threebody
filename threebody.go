package threebody

import "math"

const (
	Scaling   = 0.1
	MinRadius = 3
)

type Runner struct {
	Bodies []Body
}

type Body struct {
	Mass   float64
	Vector Vector
}

type Vector struct {
	Pos   Point
	Vel   Point
	Accel Point
}

type Point struct {
	X, Y float64
}

func (p Point) Add(p2 Point) Point {
	return Point{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}

func (p Point) Sub(p2 Point) Point {
	return Point{
		X: p.X - p2.X,
		Y: p.Y - p2.Y,
	}
}

func (p Point) Div(d float64) Point {
	return Point{
		X: p.X / d,
		Y: p.Y / d,
	}
}

func (p Point) Reset() {
	p.X, p.Y = 0, 0
}

func (p Point) Dist(p2 Point) float64 {
	dx, dy := p.X-p2.X, p.Y-p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func New(bs []Body) *Runner {
	return &Runner{Bodies: bs}
}

func (r *Runner) Update() {
	for i, b := range r.Bodies {
		v := b.Vector

		v.Accel.Reset()
		v.Pos = v.Pos.Add(v.Vel)

		r.Bodies[i].Vector = v
	}

	for i, b1 := range r.Bodies[:len(r.Bodies)-1] {
		for j, b2 := range r.Bodies[i+1:] {
			p := calcForce(b1, b2)
			b1.Vector.Accel = b1.Vector.Accel.Add(p.Div(b1.Mass))
			b2.Vector.Accel = b2.Vector.Accel.Sub(p.Div(b2.Mass))

			r.Bodies[i].Vector = b1.Vector
			r.Bodies[j+i+1].Vector = b2.Vector
		}
	}

	for i, b := range r.Bodies {
		r.Bodies[i].Vector.Vel = b.Vector.Vel.Add(b.Vector.Accel)
	}
}

func calcForce(b1, b2 Body) Point {
	p1, p2 := b1.Vector.Pos, b2.Vector.Pos
	d := p1.Dist(p2)
	if d < MinRadius {
		d = MinRadius
	}
	f := Scaling * (b1.Mass * b2.Mass) / (d * d)

	theta := math.Atan2(p2.Y-p1.Y, p2.X-p1.X)
	return Point{
		X: math.Cos(theta) * f,
		Y: math.Sin(theta) * f,
	}
}
