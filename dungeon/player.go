// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/geo"
)

type Player struct {
	Id             int
	name           string
	character      *Runner
	score          int
	motionListener MotionListener
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetScore() int {
	return p.score
}

func (p *Player) SetScore(value int) {
	p.score = value
}

func (p *Player) GetCharacter() *Runner {
	return p.character
}

func (p *Player) GetPosition() geo.Point {
	return geo.NewPoint(p.character.Rect().Left(), p.character.Rect().Top())
}

func (p *Player) SetPosition(x int, y int) {
	p.character.setPosition(x, y)
}

func (p *Player) PushInput(value int) {
	p.character.PushInput(value)
}

func (p *Player) SetMotionListener(value MotionListener) {
	p.motionListener = value
}

func (p *Player) Update() {
	runner := p.character

	if p.motionListener != nil && len(runner.inputs) > 0 {
		p.motionListener(runner.inputs)
	}
	runner.Update()
}

func NewPlayer(name string) Player {
	character := NewRunner()

	return Player{
		name:           name,
		character:      &character,
		score:          0,
		motionListener: nil,
	}
}

type MotionListener func([]int)
