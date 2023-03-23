// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package asset

import (
	"dungeon-mst/core/graphic"
	"github.com/hajimehoshi/ebiten/v2"
)

// GameGraphic defines general game graphics
type GameGraphic uint8

const (
	Legend GameGraphic = iota
)

func (g GameGraphic) Name() graphic.Name {
	return map[GameGraphic]graphic.Name{
		Legend: "keyboard_legend.png",
	}[g]
}

type GameGraphics map[GameGraphic]*graphic.Graphic

func LoadGameGraphics(load graphic.Load) GameGraphics {
	return GameGraphics{Legend: load(Legend)}
}

type gameDrawing struct {
	graphics EntityGraphics[GameGraphic]
}

func (d gameDrawing) Draw(screen *ebiten.Image) {
	d.drawLegend(screen)
}

func (d gameDrawing) drawLegend(screen *ebiten.Image) {
	screen.DrawImage(d.graphics.Get(Legend).Image, nil)
}

func NewGameDrawing(
	graphics EntityGraphics[GameGraphic],
) graphic.Draw {
	return gameDrawing{graphics: graphics}
}
