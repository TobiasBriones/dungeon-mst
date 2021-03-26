/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	"game/ai"
	"game/model"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var (
	bgImage *ebiten.Image
	match   *model.Match
)

type Game struct {
	arena       *Arena
	count       int
	legendImage *ebiten.Image
}

func (g *Game) Update() error {
	g.count++

	for i, diamond := range match.Diamonds {
		if g.arena.checkDiamondCollision(diamond) {
			remove(match.Diamonds, i)
		}
	}

	g.arena.Update(setCurrentDungeonAndPaths)

	// Generate random dungeons
	if g.count%5 == 0 {
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			reset()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bgImage, nil)

	for _, dungeon := range match.Dungeons {
		dungeon.DrawBarrier(screen)
	}
	for _, path := range match.Paths {
		path.Draw(screen)
	}
	for _, dungeon := range match.Dungeons {
		dungeon.Draw(screen)
	}

	// Draw legend image
	screen.DrawImage(g.legendImage, nil)

	// Draw diamonds
	for _, diamond := range match.Diamonds {
		diamond.Draw(screen)
	}

	// Draw remote players
	g.arena.Draw(screen)
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) onCharacterMotion(move int) {
	name := g.arena.GetPlayerName()
	println(name + " " + strconv.Itoa(move))
}

func Run() {
	game := newGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Dungeon MST")

	sendFakeInputs(game.arena)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

func newGame() Game {
	arena := NewArena()
	legendImage := loadLegendImage()
	game := Game{
		arena:       &arena,
		legendImage: legendImage,
	}

	game.arena.SetOnCharacterMotion(game.onCharacterMotion)
	return game
}

func getSize() model.Dimension {
	return model.NewDimension(screenWidth, screenHeight)
}

func init() {
	loadBg()
	match = ai.NewRandomMatch(getSize())
}

func loadBg() {
	rand.Seed(time.Now().UnixNano())
	bgNumber := rand.Intn(3) + 1
	bgName := "bg_" + strconv.Itoa(bgNumber) + ".png"
	bgImg, _, err := ebitenutil.NewImageFromFile("./assets/" + bgName)

	if err != nil {
		log.Fatal(err)
	}
	bgImage = bgImg
}

func loadLegendImage() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("./assets/keyboard_legend.png")

	if err != nil {
		log.Fatal(err)
	}
	return img
}

func setCurrentDungeonAndPaths(runner *model.Runner) {
	var currentDungeon *model.Dungeon = nil
	var currentPaths []*model.Path

	for _, dungeon := range match.Dungeons {
		if dungeon.InBounds(&runner.Rect) {
			currentDungeon = dungeon
			break
		}
	}
	for _, path := range match.Paths {
		if path.InBounds(&runner.Rect) {
			currentPaths = append(currentPaths, path)
		}
	}
	runner.SetCurrentDungeon(currentDungeon)
	runner.SetCurrentPaths(currentPaths)

	if runner.IsOutSide() {
		runner.SetDungeon(match.Dungeons[0])
	}
}

func reset() {
	match = ai.NewRandomMatch(getSize())
}

func sendFakeInputs(a *Arena) {
	ticker := time.NewTicker(50 * time.Millisecond)

	go func() {
		for range ticker.C {
			fake := randInput()
			a.PushRemotePlayerInput("remote", fake)
		}
	}()
}

func randInput() int {
	return rand.Intn(4)
}

func remove(slice []*model.Diamond, s int) []*model.Diamond {
	return append(slice[:s], slice[s+1:]...)
}
