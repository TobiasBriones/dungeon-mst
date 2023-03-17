// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/dungeon"
	"dungeon-mst/game/graphic"
	"dungeon-mst/geo"
)

type Diamond struct {
	dungeon.Diamond
	graphic.Drawing
}

// NewDiamond Creates a new game Diamond with the given initial position.
func NewDiamond(pos geo.Point, gs *graphic.Graphics) *Diamond {
	d := dungeon.NewDiamond(pos, graphic.DiamondWidthPx, graphic.DiamondHeightPx)
	g := gs.Get(dungeon.DiamondEntity)
	return &Diamond{
		d,
		graphic.NewDrawing(g, d.Rect()),
	}
}

// NewDiamondFrom Creates a new game Diamond from the given state.
func NewDiamondFrom(diamond dungeon.Diamond, gs *graphic.Graphics) *Diamond {
	g := gs.Get(dungeon.DiamondEntity)
	return &Diamond{
		diamond,
		graphic.NewDrawing(g, diamond.Rect()),
	}
}
