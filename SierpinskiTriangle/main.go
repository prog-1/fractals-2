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
	width, height int
	a, b, c, p    Point
	img           *ebiten.Image
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		a:      Point{x: 320, y: 100},
		b:      Point{x: 600, y: 400},
		c:      Point{x: 100, y: 400},
		img:    ebiten.NewImage(width, height),
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	for (g.a.y-g.b.y)*g.p.x+(g.b.x-g.a.x)*g.p.y+(g.a.x*g.b.y-g.a.y*g.b.x) > 0 &&
		(g.c.y-g.b.y)*g.p.x+(g.b.x-g.c.x)*g.p.y+(g.c.x*g.b.y-g.c.y*g.b.x) > 0 &&
		(g.c.y-g.a.y)*g.p.x+(g.a.x-g.c.x)*g.p.y+(g.c.x*g.a.y-g.c.y*g.a.x) > 0 {
		g.p = Point{x: float64(rand.Intn(g.width)), y: float64(rand.Intn(g.height))}
	}
	random := rand.Intn(3)
	if random == 0 {
		g.p.x += (g.a.x - g.p.x) / 2
		g.p.y += (g.a.y - g.p.y) / 2
	} else if random == 1 {
		g.p.x += (g.b.x - g.p.x) / 2
		g.p.y += (g.b.y - g.p.y) / 2
	} else {
		g.p.x += (g.c.x - g.p.x) / 2
		g.p.y += (g.c.y - g.p.y) / 2
	}
	g.img.Set(int(g.p.x), int(g.p.y), color.White)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.img, nil)
	ebitenutil.DrawLine(screen, g.a.x, g.a.y, g.b.x, g.b.y, color.White)
	ebitenutil.DrawLine(screen, g.b.x, g.b.y, g.c.x, g.c.y, color.White)
	ebitenutil.DrawLine(screen, g.a.x, g.a.y, g.c.x, g.c.y, color.White)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
