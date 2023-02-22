package main

import (
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	winTitle            = "Sierpinski"
	winWidth, winHeight = 800, 600
)

type point struct {
	x, y int
}
type Game struct {
	img *ebiten.Image
	a   [3]point
	p   point
}

func newSierpinski() *Game {
	var a [3]point
	a[0].x = rand.Intn(winWidth / 2)
	a[0].y = rand.Intn(winHeight / 2)
	a[1].x = rand.Intn(winWidth/2) + (winWidth / 2)
	a[1].y = rand.Intn(winHeight / 2)
	a[2].x = rand.Intn(winWidth)
	a[2].y = rand.Intn(winHeight/2) + winHeight/2
	return &Game{
		img: ebiten.NewImage(winWidth, winHeight),
		a:   a,
		p:   point{(a[0].x + a[1].x + a[2].x) / 3, ((a[0].y + a[1].y + a[2].y) / 3)},
	}
}

var c = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		os.Exit(0)
	}
	for i := 0; i < 50; i++ {
		g.img.Set(g.p.x, g.p.y, c)
		indx := rand.Int31n(3)
		g.p.x = (g.p.x + g.a[indx].x) / 2
		g.p.y = (g.p.y + g.a[indx].y) / 2
	}
	return nil
}

func (s *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.img, nil)
}
func (*Game) Layout(int, int) (w, h int) { return winWidth, winHeight }

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowTitle(winTitle)
	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(newSierpinski()); err != nil {
		log.Fatal(err)
	}
}
