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
	player            *model.Player
	remotePlayers     []*model.Player
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
				pushInput(model.MoveDirTop)
			case ebiten.KeyDown, ebiten.KeyS:
				pushInput(model.MoveDirBottom)
			case ebiten.KeyLeft, ebiten.KeyA:
				pushInput(model.MoveDirLeft)
			case ebiten.KeyRight, ebiten.KeyD:
				pushInput(model.MoveDirRight)
			}
		}
	}
}

func (a *Arena) updateRemotePlayers(update SetCurrentDungeonAndPaths) {
	// Temporarily update remote players this way
	for _, player := range a.remotePlayers {

		// Make the player receive the socket input rather than keyboard
		player.PushInput(randInput())

		// Temp implementation
		update(player.GetCharacter())

		player.Update()
	}
}

func NewArena() Arena {
	player := model.NewPlayer("local")
	remotePlayers := getTempPlayers()
	return Arena{player: &player, remotePlayers: remotePlayers}
}

type OnCharacterMotion func(int)

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
