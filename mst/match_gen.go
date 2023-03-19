// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package mst

import (
	"dungeon-mst/dungeon"
	"dungeon-mst/geo"
)

func NewRandomMatch(
	dimension geo.Dimension,
	diamondDimension dungeon.DiamondDimension,
) *dungeon.Match {
	dungeons := GenerateDungeons(dimension)
	paths := GetPaths(dungeons)
	diamonds := generateDiamonds(dungeons, diamondDimension)
	return &dungeon.Match{
		Dungeons: dungeons,
		Paths:    paths,
		Diamonds: diamonds,
	}
}

func generateDiamonds(dungeons []*dungeon.Dungeon, dim dungeon.DiamondDimension) []*dungeon.Diamond {
	var diamonds []*dungeon.Diamond

	for _, d := range dungeons {
		point := d.RandomPoint(dim.Width())
		diamond := dungeon.NewDiamond(point, dim)
		diamonds = append(diamonds, &diamond)
	}
	return diamonds
}
