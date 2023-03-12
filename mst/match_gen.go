// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package mst

import (
	"dungeon-mst/game/model"
	"dungeon-mst/geo"
)

func NewRandomMatch(dimension geo.Dimension) *model.Match {
	dungeons := GenerateDungeons(dimension)
	paths := GetPaths(dungeons)
	diamonds := generateDiamonds(dungeons)
	return &model.Match{
		Dungeons: dungeons,
		Paths:    paths,
		Diamonds: diamonds,
	}
}

func generateDiamonds(dungeons []*model.Dungeon) []*model.Diamond {
	var diamonds []*model.Diamond

	for _, dungeon := range dungeons {
		point := dungeon.RandomPoint(model.DiamondWidthPx)
		diamond := model.NewDiamond(point)
		diamonds = append(diamonds, &diamond)
	}
	return diamonds
}
