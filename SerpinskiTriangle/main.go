package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type vector struct {
	x, y float64
}

func add(a, b vector) vector {
	return vector{a.x + b.x, a.y + b.y}
}

func sub(a, b vector) vector {
	return vector{a.x - b.x, a.y - b.y}
}

func divide(v vector, a float64) vector {
	return vector{v.x / a, v.y / a}
}

func (a *vector) mod() float64 {
	return math.Sqrt(math.Pow(a.x, 2) + math.Pow(a.y, 2))
}

type game struct {
	a, b, c      vector
	screenBuffer *ebiten.Image
}

func drawLine(screen *ebiten.Image, a, b vector) {
	ebitenutil.DrawLine(screen, a.x, a.y, b.x, b.y, color.White)
}

func NewGame() *game {
	return &game{
		vector{screenWidth / 3, screenHeight / 3 * 2}, vector{screenWidth / 2, screenHeight / 3}, vector{screenWidth / 3 * 2, screenHeight / 3 * 2},
		ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.screenBuffer, &ebiten.DrawImageOptions{})
}

func serpinski(screen *ebiten.Image, a, b, c vector) {
	drawLine(screen, a, b)
	drawLine(screen, b, c)
	drawLine(screen, c, a)

	if s := sub(b, a); s.mod() <= 5 {
		return
	}

	d := add(divide(sub(b, a), 2), a)
	e := add(divide(sub(c, b), 2), b)
	f := add(divide(sub(c, a), 2), a)

	serpinski(screen, a, d, f)
	serpinski(screen, d, b, e)
	serpinski(screen, f, e, c)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame()
	serpinski(g.screenBuffer, g.a, g.b, g.c)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
