package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//---------------------------Declaration--------------------------------

const (
	sW = 640
	sH = 480
)

type Game struct {
	//global variables
	a, b, c point    //triangle
	pixels  []*point //random pixels
}

type point struct {
	x, y float64
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	//all logic on update
	g.NewPixel() //add new pixel
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, g.a.x, g.a.y, g.b.x, g.b.y, color.RGBA{100, 100, 100, 255})
	ebitenutil.DrawLine(screen, g.b.x, g.b.y, g.c.x, g.c.y, color.RGBA{100, 100, 100, 255})
	ebitenutil.DrawLine(screen, g.c.x, g.c.y, g.a.x, g.a.y, color.RGBA{100, 100, 100, 255})
	for i := range g.pixels {
		screen.Set(int(g.pixels[i].x), int(g.pixels[i].y), color.White)
	}
}

//-------------------------Functions----------------------------------

//adds new pixel if it appeared in range of triangle
func (g *Game) NewPixel() {

	tp := rand.Intn(3) //triangle point
	var np point
	last := g.pixels[len(g.pixels)-1]
	switch tp {
	case 0:
		np.x, np.y = (last.x+g.a.x)/2, (last.y+g.a.y)/2
	case 1:
		np.x, np.y = (last.x+g.b.x)/2, (last.y+g.b.y)/2
	case 2:
		np.x, np.y = (last.x+g.c.x)/2, (last.y+g.c.y)/2
	}
	g.pixels = append(g.pixels, &np)
}

//returns result of the normal lN of the line
func lN(p, a, b point) float64 {
	dy := b.y - a.y
	dx := b.x - a.x
	A := -dy
	B := dx
	C := dy*a.x - dx*a.y
	return A*p.x + B*p.y + C
}

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return sW, sH
}

func main() {

	//Window
	ebiten.SetWindowSize(sW, sH)
	ebiten.SetWindowTitle("Sierpiński triangle")

	//Game instance
	g := NewGame()                            //creating game instance
	if err := ebiten.RunGame(g); err != nil { //running game
		log.Fatal(err)
	}
}

//New game instance function
func NewGame() *Game {

	//triangle
	var a, b, c point
	a.x, a.y = sW/2, 50
	b.x, b.y = (sW/3)-100, sH-50
	c.x, c.y = ((sW/3)*2)+100, sH-50

	//random point

	var p point

	for lN(p, a, b) < 0 && lN(p, b, c) < 0 && lN(p, c, a) < 0 {
		p.x = float64(rand.Intn(sW))
		p.y = float64(rand.Intn(sH))
	}

	var pixels []*point
	pixels = append(pixels, &p)

	return &Game{a, b, c, pixels}
}
