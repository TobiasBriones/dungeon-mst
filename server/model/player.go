/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

type Player struct {
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
