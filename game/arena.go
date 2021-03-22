/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package game

import (
	"dungeon-mst/model"
	"github.com/hajimehoshi/ebiten"
	"math/rand"
)

type Arena struct {
	remotePlayers []*model.Player
}

func (a *Arena) Update(update UpdateRemotePlayer) {
	// Temporarily update remote players this way
	for _, player := range a.remotePlayers {

		// Make the player receive the socket input rather than keyboard
		player.SetInput(randInput())

		// Temp implementation
		update(player.GetCharacter())

		player.Update()
	}
}

func (a *Arena) Draw(screen *ebiten.Image) {
	for _, player := range a.remotePlayers {
		player.Draw(screen)
	}
}

func NewArena() Arena {
	remotePlayers := getTempPlayers()
	return Arena{remotePlayers: remotePlayers}
}

type UpdateRemotePlayer func(runner *model.Runner)

func getTempPlayers() []*model.Player {
	var newPlayer = func(name string) *model.Player {
		player := model.NewPlayer(name)
		return &player
	}
	return []*model.Player{
		newPlayer("remote"),
	}
}

func randInput() int {
	return rand.Intn(4)
}
