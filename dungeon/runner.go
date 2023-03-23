// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/core/geo"
)

// TODO these constants are duplicated with the graphic impl and should be
// avoided due to coupling reasons

const (
	screenWidth  = 1280
	screenHeight = 720

	frameWidth       = 32
	frameHeight      = 32
	movementLengthPx = 1
)

type Runner struct {
	rect           *geo.Rect
	scale          float64
	inputs         []int
	count          int
	currentDungeon *Dungeon
	currentPaths   []*Path
}

func (r *Runner) Rect() *geo.Rect {
	return r.rect
}

func (r *Runner) Scale() float64 {
	return r.scale
}

func (r *Runner) Count() int {
	return r.count
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
	x := int(-(frameWidth*r.scale)/2) + screenWidth/2
	y := int(-(frameHeight*r.scale)/2) + screenHeight/2
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
	return r.isInsideDungeon() && r.currentDungeon.CanMoveTowards(movement, r.rect)
}

func (r *Runner) canMoveInsidePathsTowards(movement Movement) bool {
	canMove := false

	for _, path := range r.currentPaths {
		if path.CanMoveTowards(movement, r.rect) {
			canMove = true
			break
		}
	}
	return canMove
}

func (r *Runner) walkLeft() {
	r.rect.MoveLeft(movementLengthPx)
}

func (r *Runner) walkUp() {
	r.rect.MoveTop(movementLengthPx)
}

func (r *Runner) walkRight() {
	r.rect.MoveRight(movementLengthPx)
}

func (r *Runner) walkDown() {
	r.rect.MoveBottom(movementLengthPx)
}

func (r *Runner) isInsideDungeon() bool {
	return r.currentDungeon != nil
}

func (r *Runner) setPosition(x int, y int) {
	r.rect.SetPosition(x, y)
}

func (r *Runner) CheckDiamondCollision(diamond *Diamond) bool {
	return diamond.Collides(r.rect)
}

func NewRunner() Runner {
	scale := 1.0
	rect := geo.NewRect(
		0,
		0,
		int(frameWidth*scale),
		int(frameHeight*scale),
	)
	runner := Runner{
		rect:           &rect,
		scale:          scale,
		inputs:         []int{},
		count:          0,
		currentDungeon: nil,
		currentPaths:   []*Path{},
	}
	runner.Center()
	return runner
}
