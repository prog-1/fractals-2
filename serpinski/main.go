package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 2560
	screenHeight = 1440
)

type point struct {
	x, y int
}

type Game struct {
	img *ebiten.Image
	o   [3]point
	p   point
}

var col = color.RGBA{244, 212, 124, 255}

func NewGame(width, height int) *Game {
	var o [3]point
	o[0].x = rand.Intn(screenWidth / 2)
	o[0].y = rand.Intn(screenHeight / 2)
	o[1].x = rand.Intn(screenWidth/2) + (screenWidth / 2)
	o[1].y = rand.Intn(screenHeight / 2)
	o[2].x = rand.Intn(screenWidth)
	o[2].y = rand.Intn(screenHeight/2) + screenHeight/2
	return &Game{
		o:   o,
		img: ebiten.NewImage(screenWidth, screenHeight),
		p:   point{rand.Intn(screenWidth), rand.Intn(screenHeight)},
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	for i := 0; i < 100; i++ {
		q := rand.Int31n(3)
		g.p.x = (g.p.x + g.o[q].x) / 2
		g.p.y = (g.p.y + g.o[q].y) / 2
		g.img.Set(g.p.x, g.p.y, col)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.img, nil)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
