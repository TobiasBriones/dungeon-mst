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
	motionListener MotionListener
}

func (p *Player) GetCharacter() *Runner {
	return p.character
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

func (p *Player) Draw(screen *ebiten.Image) {
	p.character.Draw(screen)
}

func NewPlayer(name string) Player {
	character := NewRunner()

	return Player{
		name:           name,
		character:      &character,
		motionListener: nil,
	}
}

type MotionListener func([]int)
