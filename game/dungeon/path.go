// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/core/graphic"
	"dungeon-mst/dungeon"
	graphicdungeon "dungeon-mst/game/graphic/dungeon"
)

type Path struct {
	dungeon.Path
	graphic.Draw
}

// NewPathFrom Creates a new game Diamond from the given state.
func NewPathFrom(path dungeon.Path, gs *graphicdungeon.Graphics) *Path {
	return &Path{
		path,
		graphicdungeon.NewPathDrawing(
			*gs.PathGraphics,
			path.HRect(),
			path.VRect(),
		),
	}
}
