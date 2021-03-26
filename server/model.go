/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	"server/ai"
	"server/model"
)

const (
	DataTypeGameInitialization = 0
	DataTypeUpdate             = 1
	DataTypeServerMessage      = 2
)

type ResponseData struct {
	Type int
	Body string
}

type MatchJSON struct {
	Dungeons []*model.DungeonJSON
	Paths    []*model.PathJSON
	Diamonds []*model.DiamondJSON
}

func NewRandomMatch() *MatchJSON {
	dungeons := ai.GenerateDungeons()
	paths := ai.GetPaths(dungeons)
	diamonds := generateDiamonds(dungeons)
	var dungeonsJSON []*model.DungeonJSON
	var pathsJSON []*model.PathJSON
	var diamondsJSON []*model.DiamondJSON

	for _, dungeon := range dungeons {
		dungeonsJSON = append(dungeonsJSON, model.NewDungeonJSON(dungeon))
	}

	for _, path := range paths {
		pathsJSON = append(pathsJSON, model.NewPathJSON(path))
	}

	for _, diamond := range diamonds {
		diamondsJSON = append(diamondsJSON, model.NewDiamondJSON(diamond))
	}
	return &MatchJSON{
		Dungeons: dungeonsJSON,
		Paths:    pathsJSON,
		Diamonds: diamondsJSON,
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
