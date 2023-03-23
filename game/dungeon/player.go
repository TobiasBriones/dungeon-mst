// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/dungeon"
	"dungeon-mst/game/asset"
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
	*dungeon.Player
	runner *Runner
}

func (p Player) Draw(screen *ebiten.Image) {
	p.runner.Draw.Draw(screen)
	p.drawName(screen)
}

func (p Player) drawName(screen *ebiten.Image) {
	name := p.GetName()
	str := name + "(" + strconv.Itoa(p.GetScore()) + ")"
	character := p.GetCharacter() // TODO Demeter law
	x := character.Rect().Left()
	y := character.Rect().Top()
	text.Draw(screen, str, mplusNormalFont, x, y, color.Black)
}

func NewPlayerFrom(
	player *dungeon.Player,
	gs *asset.Graphics,
) *Player {
	runner := player.GetCharacter()
	return &Player{
		Player: player,
		runner: NewRunnerFrom(runner, gs),
	}
}
