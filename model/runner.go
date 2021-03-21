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
	InputTypeKeyboard = 0
	InputTypeCustom   = 1
	screenWidth       = 1280
	screenHeight      = 720

	frameOX          = 0
	frameOY          = 32
	frameWidth       = 32
	frameHeight      = 32
	frameNum         = 8
	movementLengthPx = 1
)

type MotionListener func(int)

type Runner struct {
	Rect           Rect
	Scale          float64
	CustomInput    int
	MotionListener MotionListener
	inputType      int
	count          int
	image          *ebiten.Image
	currentDungeon *Dungeon
	currentPaths   []*Path
}

func (r *Runner) IsOutSide() bool {
	return !r.isInsideDungeon() && len(r.currentPaths) == 0
}

func (r *Runner) SetInputType(inputType int) {
	r.inputType = inputType
}

func (r *Runner) SetCurrentDungeon(value *Dungeon) {
	r.currentDungeon = value
}

func (r *Runner) SetCurrentPaths(value []*Path) {
	r.currentPaths = value
}

func (r *Runner) SetDungeon(value *Dungeon) {
	x := value.Cx() - frameWidth/2
	y := value.Cy() - frameHeight/2
	r.setPosition(x, y)
	r.SetCurrentDungeon(value)
}

func (r *Runner) Update() {
	r.count++

	if r.inputType == InputTypeKeyboard {
		r.readKeyboardInput()
	} else {
		r.readCustomInput()
	}
}

func (r *Runner) Draw(screen *ebiten.Image) {
	x := r.Rect.Left()
	y := r.Rect.Top()
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

func (r *Runner) readKeyboardInput() {
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			switch k {
			case ebiten.KeyUp, ebiten.KeyW:
				r.move(MoveDirTop)
			case ebiten.KeyDown, ebiten.KeyS:
				r.move(MoveDirBottom)
			case ebiten.KeyLeft, ebiten.KeyA:
				r.move(MoveDirLeft)
			case ebiten.KeyRight, ebiten.KeyD:
				r.move(MoveDirRight)
			}
		}
	}
}

func (r *Runner) readCustomInput() {
	r.move(r.CustomInput)
	r.CustomInput = -1
}

func (r *Runner) move(direction int) {
	movement := Movement{direction, 1}
	canMoveInsideDungeon := r.canMoveInsideDungeonTowards(movement)
	canMoveInsidePaths := r.canMoveInsidePathsTowards(movement)

	if !canMoveInsideDungeon && !canMoveInsidePaths {
		return
	}

	if direction == MoveDirLeft {
		r.walkLeft()
	} else if direction == MoveDirTop {
		r.walkUp()
	} else if direction == MoveDirRight {
		r.walkRight()
	} else if direction == MoveDirBottom {
		r.walkDown()
	}

	if r.MotionListener != nil {
		r.MotionListener(direction)
	}
}

func (r *Runner) canMoveInsideDungeonTowards(movement Movement) bool {
	return r.isInsideDungeon() && r.currentDungeon.CanMoveTowards(movement, &r.Rect)
}

func (r *Runner) canMoveInsidePathsTowards(movement Movement) bool {
	canMove := false

	for _, path := range r.currentPaths {
		if path.CanMoveTowards(movement, &r.Rect) {
			canMove = true
			break
		}
	}
	return canMove
}

func (r *Runner) walkLeft() {
	r.Rect.MoveLeft(movementLengthPx)
}

func (r *Runner) walkUp() {
	r.Rect.MoveTop(movementLengthPx)
}

func (r *Runner) walkRight() {
	r.Rect.MoveRight(movementLengthPx)
}

func (r *Runner) walkDown() {
	r.Rect.MoveBottom(movementLengthPx)
}

func (r *Runner) isInsideDungeon() bool {
	return r.currentDungeon != nil
}

func (r *Runner) setPosition(x int, y int) {
	r.Rect.setPosition(x, y)
}

func NewRunner() Runner {
	scale := 1.0
	runner := Runner{
		Rect: NewRect(
			0,
			0,
			int(frameWidth*scale),
			int(frameHeight*scale),
		),
		Scale:          scale,
		inputType:      InputTypeKeyboard,
		count:          0,
		currentDungeon: nil,
		currentPaths:   []*Path{},
	}
	runner.Center()

	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))

	if err != nil {
		log.Fatal(err)
	}
	runner.image = ebiten.NewImageFromImage(img)
	return runner
}
