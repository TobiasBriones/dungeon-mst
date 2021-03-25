/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

const (
	screenWidth  = 1280
	screenHeight = 720

	frameWidth       = 32
	frameHeight      = 32
	movementLengthPx = 1
)

type Runner struct {
	Rect           Rect
	Scale          float64
	inputs         []int
	count          int
	currentDungeon *Dungeon
	currentPaths   []*Path
}

func (r *Runner) IsOutSide() bool {
	return !r.isInsideDungeon() && len(r.currentPaths) == 0
}

func (r *Runner) PushInput(value int) {
	r.inputs = append(r.inputs, value)
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
	r.move()
	r.inputs = r.inputs[:0]
}

func (r *Runner) Center() {
	x := int(-(frameWidth*r.Scale)/2) + screenWidth/2
	y := int(-(frameHeight*r.Scale)/2) + screenHeight/2
	r.setPosition(x, y)
}

func (r *Runner) move() {
	if len(r.inputs) == 0 {
		return
	}

	for _, direction := range r.inputs {
		r.moveTowards(direction)
	}
}

func (r *Runner) moveTowards(direction int) {
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

func (r *Runner) CheckDiamondCollision(diamond *Diamond) bool {
	return diamond.Collides(&r.Rect)
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
		inputs:         []int{},
		count:          0,
		currentDungeon: nil,
		currentPaths:   []*Path{},
	}
	runner.Center()
	return runner
}
