package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/rand"
	"os"
	"time"
)

// Uses noise and a simple rule set to generate a Sierpinski triangle
// Usage (from within go/src/.../chaosGames): go run main.go > imgName.png | firefox img.png
// Code inspired by Processing PDE:
// Daniel Shiffman - The Chaos Game
// https://github.com/CodingTrain/website/blob/master/CodingChallenges/CC_123_ChaosGame_1/Processing/CC_123_ChaosGame_1/CC_123_ChaosGame_1.pde
// https://thecodingtrain.com/CodingChallenges/123.1-chaos-game.html

const (
	size = 100
	blackIndex = 1
	limit = 10000 // number of points to draw
)

func main() {
	sierpinskiTriangle(os.Stdout)

}

func sierpinskiTriangle(out io.Writer) {
	rand.Seed(time.Now().UnixNano())
	var palette = []color.Color{color.White, color.Black}

	rect := image.Rect(0, 0, size, size)
	img := image.NewPaletted(rect, palette)

	// triangle vertices
	var v_a = image.Point{size/2, 0}
	var v_b = image.Point{0, size}
	var v_c = image.Point{size, size}

	x, y := rand.Intn(size), rand.Intn(size)

	newPoint := image.Point{x, y}
	// img.SetColorIndex(newPoint.X, newPoint.Y, blackIndex) don't draw this one!

	for i := 0; i < limit; i++ {
		var r = rand.Intn(3) // number of vertices

		if r == 0 {
			newPoint = image.Point{int(lerp(float64(newPoint.X), float64(v_a.X), 0.5)), int(lerp(float64(newPoint.Y), float64(v_a.Y), 0.5))}
			img.SetColorIndex(newPoint.X, newPoint.Y, blackIndex)
		}
		if r == 1 {
			newPoint = image.Point{int(lerp(float64(newPoint.X), float64(v_b.X), 0.5)), int(lerp(float64(newPoint.Y), float64(v_b.Y), 0.5))}
			img.SetColorIndex(newPoint.X, newPoint.Y, blackIndex)
		}
		if r == 2 {
			newPoint = image.Point{int(lerp(float64(newPoint.X), float64(v_c.X), 0.5)), int(lerp(float64(newPoint.Y), float64(v_c.Y), 0.5))}
			img.SetColorIndex(newPoint.X, newPoint.Y, blackIndex)
		}
	}
	png.Encode(out, img)
}

// generates random noise in a .png image
func makeChaos(out io.Writer) {
	rand.Seed(time.Now().UnixNano())
	var palette = []color.Color{color.White, color.Black}
	rect := image.Rect(0, 0, size, size)
	startPoint := image.Point{rand.Intn(size), rand.Intn(size)}
	img := image.NewPaletted(rect, palette)
	img.SetColorIndex(startPoint.X, startPoint.Y, blackIndex)
	prevPoint := startPoint
	for i:=0; i < limit; i++ {
		newPoint := image.Point{rand.Intn(size), rand.Intn(size)}
		if newPoint != prevPoint {
			img.SetColorIndex(newPoint.X, newPoint.Y, blackIndex)
		}
		prevPoint = newPoint
	}
	png.Encode(out, img)

}

func lerp(start float64, stop float64, amt float64) (vl float64) {
	vl = start + (stop-start) * amt
	return
}