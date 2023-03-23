// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/core/graphic"
	"dungeon-mst/dungeon"
	graphicdungeon "dungeon-mst/game/graphic/dungeon"
)

type Runner struct {
	*dungeon.Runner
	graphic.Draw
}

func NewRunnerFrom(
	runner *dungeon.Runner,
	gs *graphicdungeon.Graphics,
) *Runner {
	return &Runner{
		runner,
		graphicdungeon.NewRunnerDrawing(
			*gs.RunnerGraphics,
			runner,
		),
	}
}
