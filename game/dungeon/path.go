// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/core/graphic"
	"dungeon-mst/dungeon"
	"dungeon-mst/game/asset"
	"dungeon-mst/game/drawing"
)

type Path struct {
	dungeon.Path
	graphic.Draw
}

// NewPathFrom Creates a new game Diamond from the given state.
func NewPathFrom(path dungeon.Path, gs *asset.Graphics) *Path {
	return &Path{
		path,
		drawing.NewPathDrawing(
			*gs.PathGraphics,
			path.HRect(),
			path.VRect(),
		),
	}
}
