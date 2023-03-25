// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package drawing

import (
	"dungeon-mst/core/geo"
	"dungeon-mst/core/graphic"
	"dungeon-mst/dungeon"
	"dungeon-mst/game/asset"
)

func NewDiamondDrawing(
	graphics asset.EntityGraphics[asset.DiamondGraphic],
	rect *geo.Rect,
) graphic.Draw {
	return graphic.NewBasicDrawing(graphics.Get(asset.Diamond), rect)
}

func DiamondSize() dungeon.DiamondDimension {
	return dungeon.NewDiamondDimension(32, 26)
}
