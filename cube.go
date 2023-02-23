package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//---------------------------Declaration--------------------------------

const (
	sW = 640
	sH = 480
)

type Game struct {
	width, height int //screen width and height
	//global variables
	cube [8]point
}

type point struct {
	x, y, z float64
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	//all logic on update
	//(can be divided on seperate functions for ex: "UpdateCircle")
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
	DrawCube(screen, g.cube)
}

func DrawCube(screen *ebiten.Image, cube [8]point) {
	clr := color.RGBA{255, 255, 255, 255}

	//1st plane
	ebitenutil.DrawLine(screen, cube[0].x, cube[0].y, cube[1].x, cube[1].y, clr)
	ebitenutil.DrawLine(screen, cube[1].x, cube[1].y, cube[2].x, cube[2].y, clr)
	ebitenutil.DrawLine(screen, cube[2].x, cube[2].y, cube[3].x, cube[3].y, clr)
	ebitenutil.DrawLine(screen, cube[3].x, cube[3].y, cube[0].x, cube[0].y, clr)

	//2nd plane
	ebitenutil.DrawLine(screen, cube[4].x, cube[4].y, cube[5].x, cube[5].y, clr)
	ebitenutil.DrawLine(screen, cube[5].x, cube[5].y, cube[6].x, cube[6].y, clr)
	ebitenutil.DrawLine(screen, cube[6].x, cube[6].y, cube[7].x, cube[7].y, clr)
	ebitenutil.DrawLine(screen, cube[7].x, cube[7].y, cube[4].x, cube[4].y, clr)

	//connectors
	ebitenutil.DrawLine(screen, cube[0].x, cube[0].y, cube[4].x, cube[4].y, clr)
	ebitenutil.DrawLine(screen, cube[1].x, cube[1].y, cube[5].x, cube[5].y, clr)
	ebitenutil.DrawLine(screen, cube[3].x, cube[3].y, cube[7].x, cube[7].y, clr)
	ebitenutil.DrawLine(screen, cube[2].x, cube[2].y, cube[6].x, cube[6].y, clr)
}

//-------------------------Functions----------------------------------

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {

	//Window
	ebiten.SetWindowSize(sW, sH)
	ebiten.SetWindowTitle("Cube")
	ebiten.SetWindowResizable(true) //enablening window resize

	//Game instance
	g := NewGame(sW, sH)                      //creating game instance
	if err := ebiten.RunGame(g); err != nil { //running game
		log.Fatal(err)
	}
}

//New game instance function
func NewGame(width, height int) *Game {

	//Cube
	var a, b, c, d, e, f, g, h point
	var cp point //center point
	cp.x, cp.y, cp.z = sW/2, sH/2, 0

	a.x, a.y, a.z = cp.x-100, cp.y-100, cp.z+100
	b.x, b.y, b.z = cp.x+100, cp.y-100, cp.z+100
	c.x, c.y, c.z = cp.x+100, cp.y+100, cp.z+100
	d.x, d.y, d.z = cp.x-100, cp.y+100, cp.z+100

	e.x, e.y, e.z = cp.x-100, cp.y-100, cp.z-100
	f.x, f.y, f.z = cp.x+100, cp.y-100, cp.z-100
	g.x, g.y, g.z = cp.x+100, cp.y+100, cp.z-100
	h.x, h.y, h.z = cp.x-100, cp.y+100, cp.z-100

	return &Game{width: width, height: height, cube: [8]point{a, b, c, d, e, f, g, h}}
}
