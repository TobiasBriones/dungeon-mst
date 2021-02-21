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
	Pos   image.Point
	Scale float64
	count int
	image *ebiten.Image
}

func (r *Runner) Center() {
	r.Pos.X = int(-(frameWidth * r.Scale) / 2)
	r.Pos.Y = int(-(frameHeight * r.Scale) / 2)
	r.Pos.X += screenWidth / 2
	r.Pos.Y += screenHeight / 2
}

func (r *Runner) Normalize() {
	pos := &r.Pos

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

func (r *Runner) Update() {
	r.count++

	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			switch k {
			case ebiten.KeyUp, ebiten.KeyW:
				r.Pos.Y--
			case ebiten.KeyDown, ebiten.KeyS:
				r.Pos.Y++
			case ebiten.KeyLeft, ebiten.KeyA:
				r.Pos.X--
			case ebiten.KeyRight, ebiten.KeyD:
				r.Pos.X++
			}
		}
	}
	r.Normalize()
}

func (r *Runner) Draw(screen *ebiten.Image) {
	pos := r.Pos
	op := &ebiten.DrawImageOptions{}
	i := (r.count / 5) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY
	rect := image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)

	op.GeoM.Scale(r.Scale, r.Scale)
	op.GeoM.Translate(float64(pos.X), float64(pos.Y))
	screen.DrawImage(r.image.SubImage(rect).(*ebiten.Image), op)
}

func NewRunner() Runner {
	runner := Runner{
		Pos:   image.Point{},
		Scale: 1,
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
