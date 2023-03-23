// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/core/geo"
	"dungeon-mst/dungeon"
	"dungeon-mst/game/graphic"
)

type DiamondGraphic uint8

const (
	Diamond DiamondGraphic = iota
)

func (g DiamondGraphic) Name() graphic.Name {
	return "diamond.png"
}

type DiamondGraphics map[DiamondGraphic]*graphic.Graphic

func LoadDiamondGraphics(load graphic.Load) DiamondGraphics {
	return DiamondGraphics{Diamond: load(Diamond)}
}

func NewDiamondDrawing(
	graphics EntityGraphics[DiamondGraphic],
	rect *geo.Rect,
) graphic.Draw {
	return graphic.NewDrawing(graphics.Get(Diamond), rect)
}

func DiamondSize() dungeon.DiamondDimension {
	return dungeon.NewDiamondDimension(32, 26)
}
