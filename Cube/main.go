package main

import (
	"image/color"
	"log"
	"math"
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
	x, y, z float64
}

type Game struct {
	width, height int
	verts         []*Point
	edges         [][2]int
}

func Rotate(a *Point, angle float64) {
	a.x = a.x*math.Cos(angle) - a.y*math.Sin(angle)
	a.y = a.x*math.Sin(angle) + a.y*math.Cos(angle)

	a.x = a.x*math.Cos(angle) - a.z*math.Sin(angle)
	a.z = a.x*math.Sin(angle) + a.z*math.Cos(angle)

	a.y = a.y*math.Cos(angle) + a.z*math.Sin(angle)
	a.z = -a.y*math.Sin(angle) + a.z*math.Cos(angle)
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		verts: []*Point{
			{100, 100, 100},
			{100, -100, 100},
			{-100, -100, 100},
			{-100, 100, 100},
			{100, 100, -100},
			{100, -100, -100},
			{-100, -100, -100},
			{-100, 100, -100},
		},
		edges: [][2]int{
			{0, 1}, {1, 2}, {2, 3}, {3, 0},
			{4, 5}, {5, 6}, {6, 7}, {7, 4},
			{0, 4}, {1, 5}, {2, 6}, {3, 7},
		},
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	for _, v := range g.verts {
		Rotate(v, math.Pi/360)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, v := range g.edges {
		ebitenutil.DrawLine(screen, g.verts[v[0]].x+float64(g.width)/2, g.verts[v[0]].y+float64(g.height)/2, g.verts[v[1]].x+float64(g.width)/2, g.verts[v[1]].y+float64(g.height)/2, color.White)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
