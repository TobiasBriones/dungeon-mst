// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/core/graphic"
	"dungeon-mst/dungeon"
	"dungeon-mst/game/asset"
)

type NewPlayer interface {
	NewPlayer(string) *Player
}

type Match struct {
	Graphics asset.Graphics
	Dungeons []*Dungeon
	Paths    []*Path
	Diamonds []*Diamond
	Bg       graphic.Draw
	Game     graphic.Draw
}

func (m *Match) ToMatchJSON() *dungeon.MatchJSON {
	var dungeons []*dungeon.Dungeon
	var paths []*dungeon.Path
	var diamonds []*dungeon.Diamond

	for _, d := range m.Dungeons {
		dungeons = append(dungeons, d.Dungeon)
	}

	for _, path := range m.Paths {
		paths = append(paths, &path.Path)
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

func (m *Match) NewPlayer(name string) *Player {
	player := dungeon.NewPlayer(name)
	return NewPlayerFrom(&player, &m.Graphics)
}

func NewMatch(m *dungeon.Match) *Match {
	graphics := asset.LoadGraphics()
	var dungeons []*Dungeon
	var paths []*Path
	var diamonds []*Diamond

	for _, d := range m.Dungeons {
		dungeons = append(dungeons, NewDungeonFrom(d, *graphics))
	}

	for _, path := range m.Paths {
		paths = append(paths, NewPathFrom(*path, graphics))
	}

	for _, diamond := range m.Diamonds {
		diamonds = append(diamonds, NewDiamondFrom(*diamond, graphics))
	}

	bg := asset.NewBackgroundDrawing(
		*graphics.BackgroundGraphics,
		asset.GetRandomBackground(),
	)
	game := asset.NewGameDrawing(*graphics.GameGraphics)

	return &Match{
		Graphics: *graphics,
		Dungeons: dungeons,
		Paths:    paths,
		Diamonds: diamonds,
		Bg:       bg,
		Game:     game,
	}
}
