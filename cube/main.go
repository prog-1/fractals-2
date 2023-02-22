package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	winTitle     = "Cube"
	screenWidth  = 1000
	screenHeight = 1000
)

var c = color.RGBA{R: 255, G: 255, B: 255, A: 255}

type (
	point struct {
		x, y, z float64
	}
	game struct {
		p [8]point
	}
)

func (g *game) rotateX() {
	for i, v := range g.p {
		g.p[i].x = v.x*math.Cos(0.0174533) - v.y*math.Sin(0.0174533)
		g.p[i].y = v.x*math.Sin(0.0174533) + v.y*math.Cos(0.0174533)
	}
}
func (g *game) rotateY() {
	for i, v := range g.p {
		g.p[i].x = v.x*math.Cos(0.0174533) - v.z*math.Sin(0.0174533)
		g.p[i].z = v.x*math.Sin(0.0174533) + v.z*math.Cos(0.0174533)
	}
}
func (g *game) rotateZ() {
	for i, v := range g.p {
		g.p[i].y = v.y*math.Cos(-0.0174533) - v.z*math.Sin(-0.0174533)
		g.p[i].z = v.y*math.Sin(-0.0174533) + v.z*math.Cos(-0.0174533)
	}
}

func NewGame() *game {
	return &game{
		[8]point{
			{-300, -300, -300},
			{-300, 300, -300},
			{300, 300, -300},
			{300, -300, -300},
			{-300, -300, 300},
			{-300, 300, 300},
			{300, 300, 300},
			{300, -300, 300},
		},
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }

func (g *game) Update() error {
	g.rotateX()
	g.rotateY()
	g.rotateZ()
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, g.p[0].x+float64(screenWidth/2), g.p[0].y+float64(screenHeight/2), g.p[1].x+float64(screenWidth/2), g.p[1].y+float64(screenHeight/2), c)
	ebitenutil.DrawLine(screen, g.p[1].x+float64(screenWidth/2), g.p[1].y+float64(screenHeight/2), g.p[2].x+float64(screenWidth/2), g.p[2].y+float64(screenHeight/2), c)
	ebitenutil.DrawLine(screen, g.p[2].x+float64(screenWidth/2), g.p[2].y+float64(screenHeight/2), g.p[3].x+float64(screenWidth/2), g.p[3].y+float64(screenHeight/2), c)
	ebitenutil.DrawLine(screen, g.p[3].x+float64(screenWidth/2), g.p[3].y+float64(screenHeight/2), g.p[0].x+float64(screenWidth/2), g.p[0].y+float64(screenHeight/2), c)
	ebitenutil.DrawLine(screen, g.p[4].x+float64(screenWidth/2), g.p[4].y+float64(screenHeight/2), g.p[5].x+float64(screenWidth/2), g.p[5].y+float64(screenHeight/2), c)
	ebitenutil.DrawLine(screen, g.p[5].x+float64(screenWidth/2), g.p[5].y+float64(screenHeight/2), g.p[6].x+float64(screenWidth/2), g.p[6].y+float64(screenHeight/2), c)
	ebitenutil.DrawLine(screen, g.p[6].x+float64(screenWidth/2), g.p[6].y+float64(screenHeight/2), g.p[7].x+float64(screenWidth/2), g.p[7].y+float64(screenHeight/2), c)
	ebitenutil.DrawLine(screen, g.p[7].x+float64(screenWidth/2), g.p[7].y+float64(screenHeight/2), g.p[4].x+float64(screenWidth/2), g.p[4].y+float64(screenHeight/2), c)
	ebitenutil.DrawLine(screen, g.p[0].x+float64(screenWidth/2), g.p[0].y+float64(screenHeight/2), g.p[4].x+float64(screenWidth/2), g.p[4].y+float64(screenHeight/2), c)
	ebitenutil.DrawLine(screen, g.p[1].x+float64(screenWidth/2), g.p[1].y+float64(screenHeight/2), g.p[5].x+float64(screenWidth/2), g.p[5].y+float64(screenHeight/2), c)
	ebitenutil.DrawLine(screen, g.p[2].x+float64(screenWidth/2), g.p[2].y+float64(screenHeight/2), g.p[6].x+float64(screenWidth/2), g.p[6].y+float64(screenHeight/2), c)
	ebitenutil.DrawLine(screen, g.p[3].x+float64(screenWidth/2), g.p[3].y+float64(screenHeight/2), g.p[7].x+float64(screenWidth/2), g.p[7].y+float64(screenHeight/2), c)
}

func main() {
	ebiten.SetWindowTitle(winTitle)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizable(true)
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
