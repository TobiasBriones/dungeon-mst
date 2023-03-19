// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/dungeon"
	graphicdungeon "dungeon-mst/game/graphic/dungeon"
)

type Match struct {
	Graphics graphicdungeon.Graphics
	Dungeons []*dungeon.Dungeon
	Paths    []*dungeon.Path
	Diamonds []*Diamond
}

func (m *Match) ToMatchJSON() *dungeon.MatchJSON {
	var dungeons []*dungeon.Dungeon
	var paths []*dungeon.Path
	var diamonds []*dungeon.Diamond

	for _, d := range m.Dungeons {
		dungeons = append(dungeons, d)
	}

	for _, path := range m.Paths {
		paths = append(paths, path)
	}

	for _, diamond := range m.Diamonds {
		diamonds = append(diamonds, &diamond.Diamond)
	}
	match := dungeon.Match{
		Dungeons: dungeons,
		Paths:    paths,
		Diamonds: diamonds,
	}
	return dungeon.NewMatchJSON(&match)
}

func NewMatch(m *dungeon.Match) *Match {
	graphics := graphicdungeon.LoadGraphics()
	var dungeons []*dungeon.Dungeon
	var paths []*dungeon.Path
	var diamonds []*Diamond

	for _, d := range m.Dungeons {
		dungeons = append(dungeons, d)
	}

	for _, path := range m.Paths {
		paths = append(paths, path)
	}

	for _, diamond := range m.Diamonds {
		diamonds = append(diamonds, NewDiamondFrom(*diamond, graphics))
	}
	return &Match{
		Dungeons: dungeons,
		Paths:    paths,
		Diamonds: diamonds,
	}
}
