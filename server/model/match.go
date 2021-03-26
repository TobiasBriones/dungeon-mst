/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

type Match struct {
	Dungeons []*Dungeon
	Paths    []*Path
	Diamonds []*Diamond
}

type MatchJSON struct {
	DungeonsJSON []*DungeonJSON
	PathsJSON    []*PathJSON
	DiamondsJSON []*DiamondJSON
}

func (m *MatchJSON) ToMatch() *Match {
	var dungeons []*Dungeon
	var paths []*Path
	var diamonds []*Diamond

	for _, dungeonJSON := range m.DungeonsJSON {
		dungeons = append(dungeons, dungeonJSON.ToDungeon())
	}

	for _, pathJSON := range m.PathsJSON {
		paths = append(paths, pathJSON.ToPath())
	}

	for _, diamondJSON := range m.DiamondsJSON {
		diamonds = append(diamonds, diamondJSON.ToDiamond())
	}
	return &Match{
		Dungeons: dungeons,
		Paths:    paths,
		Diamonds: diamonds,
	}
}

func NewMatchJSON(m *Match) *MatchJSON {
	var dungeonsJSON []*DungeonJSON
	var pathsJSON []*PathJSON
	var diamondsJSON []*DiamondJSON

	for _, dungeon := range m.Dungeons {
		dungeonsJSON = append(dungeonsJSON, NewDungeonJSON(dungeon))
	}

	for _, path := range m.Paths {
		pathsJSON = append(pathsJSON, NewPathJSON(path))
	}

	for _, diamond := range m.Diamonds {
		diamondsJSON = append(diamondsJSON, NewDiamondJSON(diamond))
	}
	return &MatchJSON{
		DungeonsJSON: dungeonsJSON,
		PathsJSON:    pathsJSON,
		DiamondsJSON: diamondsJSON,
	}
}
