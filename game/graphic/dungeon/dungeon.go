// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/game/graphic"
	"dungeon-mst/geo"
	"github.com/hajimehoshi/ebiten/v2"
)

type BackgroundGraphic uint8

const (
	Background BackgroundGraphic = iota
)

func (g BackgroundGraphic) Name() graphic.Name {
	return "dungeon_bg.png"
}

type BackgroundGraphics map[BackgroundGraphic]*graphic.Graphic

func LoadBackgroundGraphics(load graphic.Load) BackgroundGraphics {
	return BackgroundGraphics{Background: load(Diamond)}
}

func NewBackgroundDrawing(
	graphics EntityGraphics[BackgroundGraphic],
	rect *geo.Rect,
) graphic.Draw {
	return graphic.NewDrawing(graphics.Get(Background), rect)
}

type BrickGraphic uint8

const (
	Brick  BrickGraphic = iota
	BrickY BrickGraphic = iota
)

func (g BrickGraphic) Name() graphic.Name {
	return map[BrickGraphic]graphic.Name{
		Brick:  "brick.png",
		BrickY: "brick_y.png",
	}[g]
}

type BrickGraphics map[BrickGraphic]*graphic.Graphic

func LoadBrickGraphics(load graphic.Load) BrickGraphics {
	return BrickGraphics{
		Brick:  load(Brick),
		BrickY: load(BrickY),
	}
}

type brickDrawing struct {
	graphics EntityGraphics[BrickGraphic]
	rect     *geo.Rect
}

func (p brickDrawing) Draw(screen *ebiten.Image) {
	// TODO
}

func NewBrickDrawing(
	graphics EntityGraphics[BrickGraphic],
	rect *geo.Rect,
) graphic.Draw {
	return brickDrawing{
		graphics: graphics,
		rect:     rect,
	}
}
