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
	remotePlayers []*model.Runner
}

func (a *Arena) Update(update UpdateRemotePlayer) {
	// Temporarily update remote players this way
	for _, runner := range a.remotePlayers {
		// Make the runner receive the socket input rather than keyboard
		runner.CustomInput = randInput()

		update(runner)
	}
}

func (a *Arena) Draw(screen *ebiten.Image) {
	for _, runner := range a.remotePlayers {
		runner.Draw(screen)
	}
}

func NewArena() Arena {
	remotePlayers := getTempPlayers()
	return Arena{remotePlayers: remotePlayers}
}

type UpdateRemotePlayer func(runner *model.Runner)

func getTempPlayers() []*model.Runner {
	var remotePlayers []*model.Runner
	remotePlayer := model.NewRunner()

	remotePlayer.SetInputType(model.InputTypeCustom)

	remotePlayers = append(remotePlayers, &remotePlayer)
	return remotePlayers
}

func randInput() string {
	r := rand.Intn(4)

	switch r {
	case 0:
		return "W"
	case 1:
		return "A"
	case 2:
		return "D"
	case 3:
		return "S"
	}
	return ""
}
