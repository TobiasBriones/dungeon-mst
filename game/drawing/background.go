// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package drawing

import (
	"dungeon-mst/core/geo"
	"dungeon-mst/core/graphic"
	"dungeon-mst/game/asset"
)

func NewBackgroundDrawing(
	graphics asset.EntityGraphics[asset.BackgroundGraphic],
	bg asset.BackgroundGraphic,
) graphic.Draw {
	rect := geo.NewRect(0, 0, 1, 1) // Full size from origin
	return graphic.NewBasicDrawing(graphics.Get(bg), &rect)
}
