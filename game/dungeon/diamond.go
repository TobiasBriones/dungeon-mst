// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/core/graphic"
	"dungeon-mst/dungeon"
	"dungeon-mst/game/asset"
)

type Diamond struct {
	dungeon.Diamond
	graphic.Draw
}

// NewDiamondFrom Creates a new game Diamond from the given state.
func NewDiamondFrom(diamond dungeon.Diamond, gs *asset.Graphics) *Diamond {
	return &Diamond{
		diamond,
		asset.NewDiamondDrawing(*gs.DiamondGraphics, diamond.Rect()),
	}
}
