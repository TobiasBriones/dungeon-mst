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

type Runner struct {
	*dungeon.Runner
	graphic.Draw
}

func NewRunnerFrom(
	runner *dungeon.Runner,
	gs *asset.Graphics,
) *Runner {
	return &Runner{
		runner,
		drawing.NewRunnerDrawing(
			*gs.RunnerGraphics,
			runner,
		),
	}
}
