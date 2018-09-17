package gif

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"io"
	"math/rand"

	"github.com/bcspragu/threebody"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func GIF(out io.Writer) error {
	const (
		iterations = 1000
		size       = 500
		delay      = 2
	)

	bodies := make([]threebody.Body, 100)
	for i, b := range bodies {
		b.Mass = 1
		b.Vector = threebody.Vector{
			Pos: threebody.Point{
				X: (rand.Float64() * (size * 4 / 5)) + size*1/10,
				Y: (rand.Float64() * (size * 4 / 5)) + size*1/10,
			},
			Vel: threebody.Point{
				X: rand.Float64()*20 - 10,
				Y: rand.Float64()*20 - 10,
			},
		}
		bodies[i] = b
	}

	bodies = append(bodies, threebody.Body{
		Mass:   1000,
		Vector: threebody.Vector{Pos: threebody.Point{X: size / 2, Y: size / 2}},
	})

	r := threebody.New(bodies)

	anim := gif.GIF{LoopCount: -1}
	for i := 0; i < iterations; i++ {
		rect := image.Rect(0, 0, size, size)
		img := image.NewPaletted(rect, palette)
		for _, b := range r.Bodies {
			drawCircle(img, int(b.Vector.Pos.X+0.5), int(b.Vector.Pos.Y+0.5), 5, color.Black)
		}

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

		r.Update()
	}

	return gif.EncodeAll(out, &anim)
}

func drawCircle(img draw.Image, x0, y0, r int, c color.Color) {
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)

	for x > y {
		img.Set(x0+x, y0+y, c)
		img.Set(x0+y, y0+x, c)
		img.Set(x0-y, y0+x, c)
		img.Set(x0-x, y0+y, c)
		img.Set(x0-x, y0-y, c)
		img.Set(x0-y, y0-x, c)
		img.Set(x0+y, y0-x, c)
		img.Set(x0+x, y0-y, c)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
}
