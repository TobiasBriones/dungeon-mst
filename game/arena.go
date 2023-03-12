// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package main

import (
	"dungeon-mst/dungeon"
	"dungeon-mst/geo"
	"github.com/hajimehoshi/ebiten/v2"
)

type Arena struct {
	player            *dungeon.Player
	remotePlayers     []*dungeon.Player
	onCharacterMotion OnCharacterMotion
}

func (a *Arena) GetPlayerName() string {
	return a.player.GetName()
}

func (a *Arena) SetOnCharacterMotion(value OnCharacterMotion) {
	a.onCharacterMotion = value
}

func (a *Arena) Update(update SetCurrentDungeonAndPaths) {
	a.updatePlayer(update)
	a.updateRemotePlayers(update)
}

func (a *Arena) Draw(screen *ebiten.Image) {
	a.player.Draw(screen)

	for _, player := range a.remotePlayers {
		player.Draw(screen)
	}
}

func (a *Arena) PushRemotePlayerInput(id int, input int) {
	// This structure might be a hash map later
	for _, player := range a.remotePlayers {
		if player.Id == id {
			player.PushInput(input)
		}
	}
}

func (a *Arena) updatePlayer(update SetCurrentDungeonAndPaths) {
	a.updatePlayerInputs()
	update(a.player.GetCharacter())

	a.player.Update()
}

func (a *Arena) updatePlayerInputs() {
	pushInput := func(input int) {
		a.player.PushInput(input)

		if a.onCharacterMotion != nil {
			a.onCharacterMotion(input)
		}
	}

	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			switch k {
			case ebiten.KeyUp, ebiten.KeyW:
				pushInput(dungeon.MoveDirTop)
			case ebiten.KeyDown, ebiten.KeyS:
				pushInput(dungeon.MoveDirBottom)
			case ebiten.KeyLeft, ebiten.KeyA:
				pushInput(dungeon.MoveDirLeft)
			case ebiten.KeyRight, ebiten.KeyD:
				pushInput(dungeon.MoveDirRight)
			}
		}
	}
}

func (a *Arena) updateRemotePlayers(update SetCurrentDungeonAndPaths) {
	for _, player := range a.remotePlayers {
		// Temp implementation
		update(player.GetCharacter())

		player.Update()
	}
}

func (a *Arena) checkDiamondCollision(diamond *dungeon.Diamond) bool {
	collides := a.player.GetCharacter().CheckDiamondCollision(diamond)

	if collides {
		a.player.SetScore(a.player.GetScore() + 30)
		return true
	}
	return false
}

func (a *Arena) SetRemotePlayerPosition(id int, point *geo.Point) {
	for _, player := range a.remotePlayers {
		if player.Id == id {
			player.SetPosition(point.X(), point.Y())
			break
		}
	}
}

func (a *Arena) PushRemotePlayer(id int, name string, score int) {
	player := buildPlayer(id, name)
	player.SetScore(score)
	a.remotePlayers = append(a.remotePlayers, player)
}

func (a *Arena) RemoveRemotePlayer(lid int) {
	index := -1

	for i, player := range a.remotePlayers {
		if player.Id == lid {
			index = i
			break
		}
	}

	if index != -1 {
		a.remotePlayers[index] = a.remotePlayers[len(a.remotePlayers)-1]
		a.remotePlayers[len(a.remotePlayers)-1] = nil
		a.remotePlayers = a.remotePlayers[:len(a.remotePlayers)-1]
	}
}

func (a *Arena) SetRemotePlayerScore(id int) {
	for _, player := range a.remotePlayers {
		if player.Id == id {
			player.SetScore(player.GetScore() + 30)
			break
		}
	}
}

func NewArena(playerName string) Arena {
	player := dungeon.NewPlayer(playerName)
	return Arena{player: &player, remotePlayers: []*dungeon.Player{}}
}

type OnCharacterMotion func(int)

type SetCurrentDungeonAndPaths func(runner *dungeon.Runner)

func buildPlayer(id int, name string) *dungeon.Player {
	var newPlayer = func() *dungeon.Player {
		player := dungeon.NewPlayer(name)
		player.Id = id
		return &player
	}
	return newPlayer()
}
