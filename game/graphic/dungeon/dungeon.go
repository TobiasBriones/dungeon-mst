// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/game/graphic"
	"dungeon-mst/geo"
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
