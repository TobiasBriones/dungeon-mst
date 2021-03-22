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
	player        *model.Player
	remotePlayers []*model.Player
}

func (a *Arena) Update(update SetCurrentDungeonAndPaths) {

	// local player input
	input := model.MoveNone
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			switch k {
			case ebiten.KeyUp, ebiten.KeyW:
				input = model.MoveDirTop
				a.player.PushInput(input)
			case ebiten.KeyDown, ebiten.KeyS:
				input = model.MoveDirBottom
				a.player.PushInput(input)
			case ebiten.KeyLeft, ebiten.KeyA:
				input = model.MoveDirLeft
				a.player.PushInput(input)
			case ebiten.KeyRight, ebiten.KeyD:
				input = model.MoveDirRight
				a.player.PushInput(input)
			}
		}
	}
	//
	setCurrentDungeonAndPaths(a.player.GetCharacter())

	a.player.Update()

	// Temporarily update remote players this way
	for _, player := range a.remotePlayers {

		// Make the player receive the socket input rather than keyboard
		player.PushInput(randInput())

		// Temp implementation
		update(player.GetCharacter())

		player.Update()
	}
}

func (a *Arena) Draw(screen *ebiten.Image) {
	a.player.Draw(screen)

	for _, player := range a.remotePlayers {
		player.Draw(screen)
	}
}

func NewArena() Arena {
	player := model.NewPlayer("local")
	remotePlayers := getTempPlayers()
	return Arena{player: &player, remotePlayers: remotePlayers}
}

type SetCurrentDungeonAndPaths func(runner *model.Runner)

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
