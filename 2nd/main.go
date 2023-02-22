package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Height = 480
	Width  = 640
)

type point struct {
	x, y float64
}

func randomPointOnScreen() point {
	return point{float64(rand.Intn(Width)), float64(rand.Intn(Height))}
}

type Game struct {
	backBuffer      *ebiten.Image
	triangle        [3]point
	currentPosition point
}

func (g *Game) Update() error {
	for i := 0; i < 20; i++ {
		tmp := rand.Intn(3)
		x := (g.currentPosition.x + g.triangle[tmp].x) / 2
		y := (g.currentPosition.y + g.triangle[tmp].y) / 2
		g.currentPosition = point{x, y}
		g.backBuffer.Set(int(x), int(y), color.White)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.backBuffer, &ebiten.DrawImageOptions{})

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Hello, World!")
	g := &Game{ebiten.NewImage(Width, Height), [3]point{randomPointOnScreen(), randomPointOnScreen(), randomPointOnScreen()}, point{}}
	g.currentPosition.x = (g.triangle[0].x + g.triangle[1].x + g.triangle[2].x) / 3
	g.currentPosition.y = (g.triangle[0].y + g.triangle[1].y + g.triangle[2].y) / 3

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
