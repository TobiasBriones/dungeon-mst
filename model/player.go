/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

import (
	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	name           string
	character      *Runner
	input          int
	motionListener MotionListener
}

func (p *Player) GetCharacter() *Runner {
	return p.character
}

func (p *Player) SetInput(value int) {
	p.input = value
}

func (p *Player) SetMotionListener(value MotionListener) {
	p.motionListener = value
}

func (p *Player) Update() {
	runner := p.character

	runner.SetInput(p.input)
	runner.Update()

	if p.motionListener != nil {
		p.motionListener(p.input)
	}
	p.input = MoveNone
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.character.Draw(screen)
}

func NewPlayer(name string) Player {
	character := NewRunner()

	return Player{
		name:           name,
		character:      &character,
		input:          MoveNone,
		motionListener: nil,
	}
}

type MotionListener func(int)
