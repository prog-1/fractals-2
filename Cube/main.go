package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type vector struct {
	x, y, z float64
}

func (v *vector) rotate(rad rotator) {
	// Rotation around Z
	v.x = v.x*math.Cos(rad.z) - v.y*math.Sin(rad.z)
	v.y = v.x*math.Sin(rad.z) + v.y*math.Cos(rad.z)

	// Rotation around Y
	v.x = v.x*math.Cos(rad.y) - v.z*math.Sin(rad.y)
	v.z = v.x*math.Sin(rad.y) + v.z*math.Cos(rad.y)

	// Rotation around X
	v.y = v.y*math.Cos(rad.x) + v.z*math.Sin(rad.x)
	v.z = -v.y*math.Sin(rad.x) + v.z*math.Cos(rad.x)
}

type rotator struct {
	x, y, z float64
}

type line struct {
	a, b float64
}

type rect struct {
	p   [8]vector
	clr color.Color
}

func (r *rect) draw(screen *ebiten.Image) {
	halfWidth, halfHeight := float64(screenWidth/2), float64(screenHeight/2)
	// Drawing near plane
	ebitenutil.DrawLine(screen, r.p[0].x+halfWidth, r.p[0].y+halfHeight, r.p[1].x+halfWidth, r.p[1].y+halfHeight, r.clr)
	ebitenutil.DrawLine(screen, r.p[1].x+halfWidth, r.p[1].y+halfHeight, r.p[2].x+halfWidth, r.p[2].y+halfHeight, r.clr)
	ebitenutil.DrawLine(screen, r.p[2].x+halfWidth, r.p[2].y+halfHeight, r.p[3].x+halfWidth, r.p[3].y+halfHeight, r.clr)
	ebitenutil.DrawLine(screen, r.p[3].x+halfWidth, r.p[3].y+halfHeight, r.p[0].x+halfWidth, r.p[0].y+halfHeight, r.clr)

	// Drawing far plane
	ebitenutil.DrawLine(screen, r.p[4].x+halfWidth, r.p[4].y+halfHeight, r.p[5].x+halfWidth, r.p[5].y+halfHeight, r.clr)
	ebitenutil.DrawLine(screen, r.p[5].x+halfWidth, r.p[5].y+halfHeight, r.p[6].x+halfWidth, r.p[6].y+halfHeight, r.clr)
	ebitenutil.DrawLine(screen, r.p[6].x+halfWidth, r.p[6].y+halfHeight, r.p[7].x+halfWidth, r.p[7].y+halfHeight, r.clr)
	ebitenutil.DrawLine(screen, r.p[7].x+halfWidth, r.p[7].y+halfHeight, r.p[4].x+halfWidth, r.p[4].y+halfHeight, r.clr)

	// Drawing connections between planes
	ebitenutil.DrawLine(screen, r.p[0].x+halfWidth, r.p[0].y+halfHeight, r.p[4].x+halfWidth, r.p[4].y+halfHeight, r.clr)
	ebitenutil.DrawLine(screen, r.p[1].x+halfWidth, r.p[1].y+halfHeight, r.p[5].x+halfWidth, r.p[5].y+halfHeight, r.clr)
	ebitenutil.DrawLine(screen, r.p[2].x+halfWidth, r.p[2].y+halfHeight, r.p[6].x+halfWidth, r.p[6].y+halfHeight, r.clr)
	ebitenutil.DrawLine(screen, r.p[3].x+halfWidth, r.p[3].y+halfHeight, r.p[7].x+halfWidth, r.p[7].y+halfHeight, r.clr)
}

func (r *rect) rotate(rad rotator) {
	for i := range r.p {
		r.p[i].rotate(rad)
	}
}

type game struct {
	r            *rect
	screenBuffer *ebiten.Image
}

func NewGame() *game {
	return &game{
		&rect{
			[8]vector{
				{-200, -200, -200}, // NearBottomLeft
				{-200, 200, -200},  // NearTopLeft
				{200, 200, -200},   // NearTopRight
				{200, -200, -200},  // NearBottomRight

				{-200, -200, 200}, // FarBottomLeft
				{-200, 200, 200},  // FarTopLeft
				{200, 200, 200},   // FarTopRight
				{200, -200, 200},  // FarBottomRight
			},
			color.White,
		},
		ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.r.rotate(rotator{math.Pi / 180, math.Pi / 180, math.Pi / 180})
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	g.r.draw(screen)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
