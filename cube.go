package main

import (
	"image/color"
	"log"
	"math"
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
	angle float64//point rotation angle 
	cp point//cube central point
	cube [8]point
}

type point struct {
	x, y, z float64
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	//all logic on update
	for i := range g.cube{
		g.cube[i].rotate(g.angle)
	}
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawCube(screen, g.cube)
}

func (g *Game)DrawCube(screen *ebiten.Image, cube [8]point) {
	clr := color.RGBA{255, 255, 255, 255}

	//1st plane
	ebitenutil.DrawLine(screen, g.cp.x+cube[0].x, g.cp.y+cube[0].y, g.cp.x+cube[1].x, g.cp.y+cube[1].y, clr)
	ebitenutil.DrawLine(screen, g.cp.x+cube[1].x, g.cp.y+cube[1].y, g.cp.x+cube[2].x, g.cp.y+cube[2].y, clr)
	ebitenutil.DrawLine(screen, g.cp.x+cube[2].x, g.cp.y+cube[2].y, g.cp.x+cube[3].x, g.cp.y+cube[3].y, clr)
	ebitenutil.DrawLine(screen, g.cp.x+cube[3].x, g.cp.y+cube[3].y, g.cp.x+cube[0].x, g.cp.y+cube[0].y, clr)

	//2nd plane
	ebitenutil.DrawLine(screen, g.cp.x+cube[4].x, g.cp.y+cube[4].y, g.cp.x+cube[5].x, g.cp.y+cube[5].y, clr)
	ebitenutil.DrawLine(screen, g.cp.x+cube[5].x, g.cp.y+cube[5].y, g.cp.x+cube[6].x, g.cp.y+cube[6].y, clr)
	ebitenutil.DrawLine(screen, g.cp.x+cube[6].x, g.cp.y+cube[6].y, g.cp.x+cube[7].x, g.cp.y+cube[7].y, clr)
	ebitenutil.DrawLine(screen, g.cp.x+cube[7].x, g.cp.y+cube[7].y, g.cp.x+cube[4].x, g.cp.y+cube[4].y, clr)

	//connectors
	ebitenutil.DrawLine(screen, g.cp.x+cube[0].x, g.cp.y+cube[0].y, g.cp.x+cube[4].x, g.cp.y+cube[4].y, clr)
	ebitenutil.DrawLine(screen, g.cp.x+cube[1].x, g.cp.y+cube[1].y, g.cp.x+cube[5].x, g.cp.y+cube[5].y, clr)
	ebitenutil.DrawLine(screen, g.cp.x+cube[3].x, g.cp.y+cube[3].y, g.cp.x+cube[7].x, g.cp.y+cube[7].y, clr)
	ebitenutil.DrawLine(screen, g.cp.x+cube[2].x, g.cp.y+cube[2].y, g.cp.x+cube[6].x, g.cp.y+cube[6].y, clr)
}

//-------------------------Functions----------------------------------

//rotates the point on given angle on all axis
func (p *point) rotate(angle float64){

	//X plane
	p.y = p.y*math.Cos(angle) + p.z*math.Sin(angle)
	p.z = -p.y*math.Sin(angle) + p.z*math.Cos(angle)

	//Y plane
	p.x = p.x * math.Cos(angle) - p.z * math.Sin(angle)
	p.z = p.x * math.Sin(angle) + p.z * math.Cos(angle)

	//Z plane
	p.x = p.x * math.Cos(angle) - p.y * math.Sin(angle)
	p.y = p.x * math.Sin(angle) + p.y * math.Cos(angle)

}

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
	var cp point //center point
	cp.x, cp.y, cp.z = sW/2, sH/2, 0
	
	var a, b, c, d, e, f, g, h point

	a.x, a.y, a.z = -100, -100, 100
	b.x, b.y, b.z = 100, -100, 100
	c.x, c.y, c.z = 100, 100, 100
	d.x, d.y, d.z = -100, 100, 100

	e.x, e.y, e.z = -100, -100, -100
	f.x, f.y, f.z = 100, -100, -100
	g.x, g.y, g.z = 100, 100, -100
	h.x, h.y, h.z = -100, 100, -100

	return &Game{width: width, height: height, angle: math.Pi/360, cp: cp, cube: [8]point{a, b, c, d, e, f, g, h}}
}
