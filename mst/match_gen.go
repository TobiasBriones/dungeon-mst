// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package mst

import (
	"dungeon-mst/dungeon"
	"dungeon-mst/game/graphic"
	"dungeon-mst/geo"
)

func NewRandomMatch(dimension geo.Dimension) *dungeon.Match {
	dungeons := GenerateDungeons(dimension)
	paths := GetPaths(dungeons)
	diamonds := generateDiamonds(dungeons)
	return &dungeon.Match{
		Dungeons: dungeons,
		Paths:    paths,
		Diamonds: diamonds,
	}
}

func generateDiamonds(dungeons []*dungeon.Dungeon) []*dungeon.Diamond {
	var diamonds []*dungeon.Diamond

	for _, d := range dungeons {
		point := d.RandomPoint(graphic.DiamondWidthPx)
		diamond := dungeon.NewDiamond(point, graphic.DiamondWidthPx, graphic.DiamondHeightPx)
		diamonds = append(diamonds, &diamond)
	}
	return diamonds
}
