/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	"game/client"
	"game/model"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
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
)

type Game struct {
	match         *model.Match
	arena         *Arena
	count         int
	legendImage   *ebiten.Image
	matchCh       chan *client.MatchInit
	updateCh      chan *client.Update
	quit          chan bool
	remainingTime time.Duration
}

func (g *Game) SetMatch(value *model.Match) {
	g.match = value
}

func (g *Game) Update() error {
	if g.match == nil {
		return nil
	}
	g.count++

	for i, diamond := range g.match.Diamonds {
		if g.arena.checkDiamondCollision(diamond) {
			remove(g.match.Diamonds, i)
		}
	}

	g.arena.Update(g.setCurrentDungeonAndPaths)

	// Generate random dungeons
	if g.count%5 == 0 {
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			g.reset()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.match == nil {
		g.drawStartScreen(screen)
		return
	}
	screen.DrawImage(bgImage, nil)

	for _, dungeon := range g.match.Dungeons {
		dungeon.DrawBarrier(screen)
	}
	for _, path := range g.match.Paths {
		path.Draw(screen)
	}
	for _, dungeon := range g.match.Dungeons {
		dungeon.Draw(screen)
	}

	// Draw legend image
	screen.DrawImage(g.legendImage, nil)

	// Draw diamonds
	for _, diamond := range g.match.Diamonds {
		diamond.Draw(screen)
	}

	// Draw remote players
	g.arena.Draw(screen)

	ebitenutil.DebugPrint(screen, strconv.FormatInt(int64(g.remainingTime/time.Second), 10))
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) onCharacterMotion(move int) {
	//name := g.arena.GetPlayerName()
	//println(name + " " + strconv.Itoa(move))
}

func (g *Game) setCurrentDungeonAndPaths(runner *model.Runner) {
	var currentDungeon *model.Dungeon = nil
	var currentPaths []*model.Path

	for _, dungeon := range g.match.Dungeons {
		if dungeon.InBounds(&runner.Rect) {
			currentDungeon = dungeon
			break
		}
	}
	for _, path := range g.match.Paths {
		if path.InBounds(&runner.Rect) {
			currentPaths = append(currentPaths, path)
		}
	}
	runner.SetCurrentDungeon(currentDungeon)
	runner.SetCurrentPaths(currentPaths)

	if runner.IsOutSide() {
		runner.SetDungeon(g.match.Dungeons[0])
	}
}

func (g *Game) reset() {
	//g.match = ai.NewRandomMatch(getSize())
}

func (g *Game) drawStartScreen(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 5, G: 2, B: 2, A: 255})
}

func (g *Game) watchRemainingTime() {
	for {
		select {
		case <-time.After(1 * time.Second):
			g.remainingTime -= time.Second
		case <-g.quit:
			return
		}
	}
}

func Run() {
	game := newGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Dungeon MST")

	go func() {
		for {
			select {
			case game.quit <- true:
			default:
			}
			m := <-game.matchCh

			game.SetMatch(m.Match)
			game.remainingTime = m.RemainingTime
			go game.watchRemainingTime()
		}
	}()
	go func() {
		for {
			u := <-game.updateCh
			game.arena.PushRemotePlayerInput("remote", u.M)
		}
	}()
	//sendFakeInputs(game.arena)

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

	matchCh := make(chan *client.MatchInit)
	updateCh := make(chan *client.Update)
	quit := make(chan bool)

	game.matchCh = matchCh
	game.updateCh = updateCh
	game.quit = quit

	go client.Run(matchCh, updateCh)
	return game
}

func getSize() model.Dimension {
	return model.NewDimension(screenWidth, screenHeight)
}

func init() {
	loadBg()
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
