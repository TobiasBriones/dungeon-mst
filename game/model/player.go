/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

import (
	"dungeon-mst/math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"strconv"
)

var (
	mplusNormalFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

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

func (p *Player) GetPosition() math.Point {
	return math.NewPoint(p.character.Rect.Left(), p.character.Rect.Top())
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

func (p *Player) Draw(screen *ebiten.Image) {
	p.character.Draw(screen)
	p.drawName(screen)
}

func (p *Player) drawName(screen *ebiten.Image) {
	name := p.name
	str := name + "(" + strconv.Itoa(p.score) + ")"
	character := p.character
	x := character.Rect.Left()
	y := character.Rect.Top()
	text.Draw(screen, str, mplusNormalFont, x, y, color.Black)
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
