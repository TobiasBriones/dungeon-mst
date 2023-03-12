/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package ai

import "dungeon-mst/server/model"

func NewRandomMatch() *model.Match {
	dungeons := GenerateDungeons()
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
