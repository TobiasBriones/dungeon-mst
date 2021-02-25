/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"image"
	"log"
)

const (
	screenWidth  = 1280
	screenHeight = 720

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
)

type Runner struct {
	Rect  Rect
	Scale float64
	count int
	image *ebiten.Image
}

func (r *Runner) Update(dungeon *Dungeon) {
	r.count++

	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			switch k {
			case ebiten.KeyUp, ebiten.KeyW:
				r.walkUp()
			case ebiten.KeyDown, ebiten.KeyS:
				r.walkDown()
			case ebiten.KeyLeft, ebiten.KeyA:
				r.walkLeft()
			case ebiten.KeyRight, ebiten.KeyD:
				r.walkRight()
			}
		}
	}
	r.normalize(dungeon)
}

func (r *Runner) Draw(screen *ebiten.Image) {
	x := r.Rect.Left
	y := r.Rect.Top
	op := &ebiten.DrawImageOptions{}
	i := (r.count / 5) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY
	rect := image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)

	op.GeoM.Scale(r.Scale, r.Scale)
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(r.image.SubImage(rect).(*ebiten.Image), op)
}

func (r *Runner) Center() {
	x := int(-(frameWidth*r.Scale)/2) + screenWidth/2
	y := int(-(frameHeight*r.Scale)/2) + screenHeight/2
	r.setPosition(x, y)
}

func (r *Runner) normalize(dungeon *Dungeon) {
	pos := Point{r.Rect.Left, r.Rect.Top}

	// Check for screen collision
	if pos.X < 0 {
		pos.X = 0
	}
	if pos.X > screenWidth-int(frameWidth*r.Scale) {
		pos.X = screenWidth - int(frameWidth*r.Scale)
	}
	if pos.Y < 0 {
		pos.Y = 0
	}
	if pos.Y > screenHeight-int(frameHeight*r.Scale) {
		pos.Y = screenHeight - int(frameHeight*r.Scale)
	}
}

func (r *Runner) walkLeft() {
	r.setPosition(r.Rect.Left-1, r.Rect.Top)
}

func (r *Runner) walkUp() {
	r.setPosition(r.Rect.Left, r.Rect.Top-1)
}

func (r *Runner) walkRight() {
	r.setPosition(r.Rect.Left+1, r.Rect.Top)
}

func (r *Runner) walkDown() {
	r.setPosition(r.Rect.Left, r.Rect.Top+1)
}

func (r *Runner) setPosition(x int, y int) {
	r.Rect.Left = x
	r.Rect.Top = y
	r.Rect.Right = x + int(frameWidth*r.Scale)
	r.Rect.Bottom = y + int(frameHeight*r.Scale)
}

func NewRunner() Runner {
	scale := 1.0
	runner := Runner{
		Rect: Rect{
			Left:   0,
			Top:    0,
			Right:  int(frameWidth * scale),
			Bottom: int(frameHeight * scale),
		},
		Scale: scale,
		count: 0,
	}
	runner.Center()

	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))

	if err != nil {
		log.Fatal(err)
	}
	runner.image = ebiten.NewImageFromImage(img)
	return runner
}
