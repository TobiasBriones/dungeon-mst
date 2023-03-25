// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package asset

import (
	"dungeon-mst/core/graphic"
)

type DungeonBackgroundGraphic uint8

const (
	DungeonBackground DungeonBackgroundGraphic = iota
)

func (g DungeonBackgroundGraphic) Name() graphic.Name {
	return "dungeon_bg.png"
}

type DungeonBackgroundGraphics map[DungeonBackgroundGraphic]*graphic.Graphic

func LoadDungeonBackgroundGraphics(load graphic.Load) DungeonBackgroundGraphics {
	return DungeonBackgroundGraphics{
		DungeonBackground: load(DungeonBackground),
	}
}

type BrickGraphic uint8

const (
	Brick  BrickGraphic = iota
	BrickY BrickGraphic = iota
)

func (g BrickGraphic) Name() graphic.Name {
	return map[BrickGraphic]graphic.Name{
		Brick:  "brick.png",
		BrickY: "brick_y.png",
	}[g]
}

type BrickGraphics map[BrickGraphic]*graphic.Graphic

func LoadBrickGraphics(load graphic.Load) BrickGraphics {
	return BrickGraphics{
		Brick:  load(Brick),
		BrickY: load(BrickY),
	}
}
