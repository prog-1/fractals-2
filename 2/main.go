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
	screenWidth  = 960
	screenHeight = 640
)

type Point struct {
	x, y float64
}

type Game struct {
	width, height  int
	p1, p2, p3, p4 Point
	screen         *ebiten.Image
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	g.SierpińskiTriangle(g.screen)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	c := color.White
	screen.DrawImage(g.screen, nil)
	ebitenutil.DrawLine(screen, g.p1.x, g.p1.y, g.p2.x, g.p2.y, c)
	ebitenutil.DrawLine(screen, g.p2.x, g.p2.y, g.p3.x, g.p3.y, c)
	ebitenutil.DrawLine(screen, g.p3.x, g.p3.y, g.p1.x, g.p1.y, c)
}

func (g *Game) SierpińskiTriangle(screen *ebiten.Image) {
	c := color.White
	for (g.p1.y-g.p2.y)*g.p4.x+(g.p2.x-g.p1.x)*g.p4.y+(g.p1.x*g.p2.y-g.p1.y*g.p2.x) > 0 &&
		(g.p3.y-g.p2.y)*g.p4.x+(g.p2.x-g.p3.x)*g.p4.y+(g.p3.x*g.p2.y-g.p3.y*g.p2.x) > 0 &&
		(g.p3.y-g.p1.y)*g.p4.x+(g.p1.x-g.p3.x)*g.p4.y+(g.p3.x*g.p1.y-g.p3.y*g.p1.x) > 0 {
		g.p4 = Point{x: float64(rand.Intn(g.width)), y: float64(rand.Intn(g.height))}
	}
	switch rand.Intn(3) {
	case 0:
		g.p4.x += (g.p1.x - g.p4.x) / 2
		g.p4.y += (g.p1.y - g.p4.y) / 2
	case 1:
		g.p4.x += (g.p2.x - g.p4.x) / 2
		g.p4.y += (g.p2.y - g.p4.y) / 2
	case 2:
		g.p4.x += (g.p3.x - g.p4.x) / 2
		g.p4.y += (g.p3.y - g.p4.y) / 2
	}
	screen.Set(int(g.p4.x), int(g.p4.y), c)
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		p1:     Point{float64(rand.Intn(width)), float64(rand.Intn(height))},
		p2:     Point{float64(rand.Intn(width)), float64(rand.Intn(height))},
		p3:     Point{float64(rand.Intn(width)), float64(rand.Intn(height))},
		screen: ebiten.NewImage(width, height),
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
