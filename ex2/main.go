package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Point struct {
	x, y float64
}

type Game struct {
	a, b, c, g    Point
	img           *ebiten.Image
	width, height int
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
	}
}

func (g *Game) Update() error {
	return nil
}
func (g *Game) DrawPoints(a, b, c Point) {
	x, y := rand.Intn(screenWidth), rand.Intn(screenHeight)
	v := Point{b.x - a.x, b.y - a.y}
	u := Point{c.x - b.x, c.y - b.y}
	z := Point{a.x - c.x, a.y - c.y}
	q1 := (float64(x)-a.x)*v.y - (float64(y)-a.y)*v.x
	q2 := (float64(x)-b.x)*u.y - (float64(y)-b.y)*u.x
	q3 := (float64(x)-c.x)*z.y - (float64(y)-c.y)*z.x
	if q1 < 0 && q2 < 0 && q3 < 0 {
		g.g = Point{x: float64(x), y: float64(y)}
	}
	switch rand.Intn(3) {
	case 0:
		g.g.x += (g.a.x - g.g.x) / 2
		g.g.y += (g.a.y - g.g.y) / 2
	case 1:
		g.g.x += (g.b.x - g.g.x) / 2
		g.g.y += (g.b.y - g.g.y) / 2
	case 2:
		g.g.x += (g.c.x - g.g.x) / 2
		g.g.y += (g.c.y - g.g.y) / 2
	}
	g.img.Set(int(g.g.x), int(g.g.y), color.RGBA{R: 245, G: 114, B: 227, A: 255})
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawPoints(g.a, g.b, g.c)
	screen.DrawImage(g.img, nil)
	ebitenutil.DrawLine(g.img, float64(g.a.x), float64(g.a.y), float64(g.b.x), float64(g.b.y), color.RGBA{R: 245, G: 114, B: 227, A: 255})
	ebitenutil.DrawLine(g.img, float64(g.b.x), float64(g.b.y), float64(g.c.x), float64(g.c.y), color.RGBA{R: 195, G: 114, B: 245, A: 255})
	ebitenutil.DrawLine(g.img, float64(g.c.x), float64(g.c.y), float64(g.a.x), float64(g.a.y), color.RGBA{R: 195, G: 114, B: 245, A: 255})
}
func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	a1 := Point{float64(rand.Intn(520)), float64(rand.Intn(360))}
	rand.Seed(time.Now().UnixNano())
	if err := ebiten.RunGame(&Game{
		a:   Point{a1.x, a1.y},
		b:   Point{a1.x + 100, a1.y},
		c:   Point{a1.x + 50, a1.y + 100},
		img: ebiten.NewImage(screenWidth, screenHeight)}); err != nil {
		log.Fatal(err)
	}
}
