// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/core/graphic"
)

type Type interface {
	DiamondGraphic |
		PathGraphic |
		RunnerGraphic |
		BackgroundGraphic |
		BrickGraphic |
		DungeonBackgroundGraphic |
		GameGraphic
}

// EntityGraphics defines and loads the physical graphics for the given
// dungeon.Entity.
type EntityGraphics[T Type] struct {
	graphics map[T]*graphic.Graphic
}

func (g *EntityGraphics[T]) Get(t T) *graphic.Graphic {
	return g.graphics[t]
}

func NewEntityGraphics[T Type](gs map[T]*graphic.Graphic) *EntityGraphics[T] {
	return &EntityGraphics[T]{graphics: gs}
}

type Graphics struct {
	DiamondGraphics           *EntityGraphics[DiamondGraphic]
	PathGraphics              *EntityGraphics[PathGraphic]
	RunnerGraphics            *EntityGraphics[RunnerGraphic]
	BackgroundGraphics        *EntityGraphics[BackgroundGraphic]
	BrickGraphics             *EntityGraphics[BrickGraphic]
	DungeonBackgroundGraphics *EntityGraphics[DungeonBackgroundGraphic]
	GameGraphics              *EntityGraphics[GameGraphic]
}

// LoadGraphics loads the graphic assets of the game into memory.
// These graphics only need to be loaded once, and be reused when drawing.
// That is, there's a 1:n relation between a graphic and the instances of the
// same entities that appear on the game.
func LoadGraphics() *Graphics {
	diamonds := LoadDiamondGraphics(loadNamedGraphic)
	paths := LoadPathGraphics(loadNamedGraphic)
	runners := LoadRunnerGraphics()
	backgrounds := LoadBackgroundGraphics(loadNamedGraphic)
	bricks := LoadBrickGraphics(loadNamedGraphic)
	dungeonBackgrounds := LoadDungeonBackgroundGraphics(loadNamedGraphic)
	games := LoadGameGraphics(loadNamedGraphic)

	return &Graphics{
		DiamondGraphics:           NewEntityGraphics(diamonds),
		PathGraphics:              NewEntityGraphics(paths),
		RunnerGraphics:            NewEntityGraphics(runners),
		BackgroundGraphics:        NewEntityGraphics(backgrounds),
		BrickGraphics:             NewEntityGraphics(bricks),
		DungeonBackgroundGraphics: NewEntityGraphics(dungeonBackgrounds),
		GameGraphics:              NewEntityGraphics(games),
	}
}

func loadNamedGraphic(g graphic.NamedGraphic) *graphic.Graphic {
	return graphic.LoadGraphicFromAssets(string(g.Name()))
}
