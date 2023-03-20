// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package main

import (
	"dungeon-mst/dungeon"
	"dungeon-mst/game/client"
	game "dungeon-mst/game/dungeon"
	"dungeon-mst/geo"
	"encoding/json"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"io/ioutil"
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
	user    User
)

type User struct {
	Id   int
	Name string
}

type Game struct {
	match         *game.Match
	arena         *Arena
	count         int
	legendImage   *ebiten.Image
	matchCh       chan *client.MatchInit
	updateCh      chan *client.Update
	sendUpdateCh  chan *client.Update
	joinCh        chan *client.PlayerJoin
	leaveCh       chan int
	quit          chan bool
	remainingTime time.Duration
}

func (g *Game) IsPaused() bool {
	return len(g.match.Diamonds) == 0
}

func (g *Game) SetMatch(value *game.Match) {
	g.match = value

	g.arena.init(user, g.match)
	g.arena.player.SetScore(0)

	for _, player := range g.arena.remotePlayers {
		player.SetScore(0)
	}
}

func (g *Game) Update() error {
	if g.match == nil {
		return nil
	}
	if g.IsPaused() {
		return nil
	}
	g.count++
	diamondIndex := -1

	for i, diamond := range g.match.Diamonds {
		if g.arena.checkDiamondCollision(&diamond.Diamond) {
			diamondIndex = i
			break
		}
	}

	if diamondIndex != -1 {
		g.match.Diamonds = remove(g.match.Diamonds, diamondIndex)
	}

	g.arena.Update(g.setCurrentDungeonAndPaths)

	position := g.arena.player.GetPosition()
	update := &client.Update{
		Id: user.Id,
		//Move: move,
		PointJSON:    *dungeon.NewPointJSON(&position),
		DiamondIndex: diamondIndex,
	}
	g.sendUpdateCh <- update

	/*
		// Generate random dungeons
		if g.count%5 == 0 {
			if ebiten.IsKeyPressed(ebiten.KeyR) {
				g.reset()
			}
		}*/
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.match == nil {
		g.drawStartScreen(screen)
		return
	}
	screen.DrawImage(bgImage, nil)

	for _, d := range g.match.Dungeons {
		d.DrawBarrier(screen)
	}
	for _, path := range g.match.Paths {
		path.Draw.Draw(screen)
	}
	for _, d := range g.match.Dungeons {
		d.Draw(screen)
	}

	// Draw legend image
	screen.DrawImage(g.legendImage, nil)

	// Draw diamonds
	for _, diamond := range g.match.Diamonds {
		diamond.Draw.Draw(screen)
	}

	// Draw remote players
	g.arena.Draw(screen)

	timeLeft := strconv.FormatInt(int64(g.remainingTime/time.Second), 10)
	text.Draw(screen, timeLeft, mplusNormalFont, screenWidth-200, 96, color.White)

	if g.IsPaused() {
		g.drawPauseScreen(screen)
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) onCharacterMotion(move int) {

}

func (g *Game) setCurrentDungeonAndPaths(runner *dungeon.Runner) {
	var currentDungeon *dungeon.Dungeon = nil
	var currentPaths []*dungeon.Path

	for _, dungeon := range g.match.Dungeons {
		if dungeon.InBounds(runner.Rect()) {
			currentDungeon = dungeon
			break
		}
	}
	for _, path := range g.match.Paths {
		if path.InBounds(runner.Rect()) {
			currentPaths = append(currentPaths, &path.Path)
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

func (g *Game) drawPauseScreen(screen *ebiten.Image) {
	player := g.arena.player
	str := player.GetName() + "(" + strconv.Itoa(player.GetScore()) + ")"
	text.Draw(screen, str, mplusNormalFont, screenWidth/2-200, 96, color.Black)

	for i, player := range g.arena.remotePlayers {
		str := player.GetName() + "(" + strconv.Itoa(player.GetScore()) + ")"
		text.Draw(screen, str, mplusNormalFont, screenWidth/2-200, 64+(i+1)*96, color.Black)
	}
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
	g := newGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Dungeon MST")

	go func() {
		for {
			select {
			case g.quit <- true:
			default:
			}
			m := <-g.matchCh

			gameMatch := game.NewMatch(m.Match)

			g.SetMatch(gameMatch)
			g.remainingTime = m.RemainingTime

			for _, player := range m.Players {
				g.arena.PushRemotePlayer(player.Id, player.Name, player.Score)
			}
			go g.watchRemainingTime()
		}
	}()
	go func() {
		for {
			u := <-g.updateCh

			if u.Id == user.Id {
				continue
			}
			//log.Println("Receiving update for player:", u.Id)
			g.arena.SetRemotePlayerPosition(u.Id, u.PointJSON.ToPoint())

			if u.DiamondIndex != -1 {
				g.match.Diamonds = remove(g.match.Diamonds, u.DiamondIndex)

				g.arena.SetRemotePlayerScore(u.Id)
			}

			//game.arena.PushRemotePlayerInput(u.Id, u.Move)
		}
	}()
	go func() {
		for {
			j := <-g.joinCh

			if j.Id == user.Id {
				continue
			}
			log.Println("Joining player:", j.Id)
			g.arena.PushRemotePlayer(j.Id, j.Name, 0)
		}
	}()

	go func() {
		for {
			lid := <-g.leaveCh

			g.arena.RemoveRemotePlayer(lid)
		}
	}()

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

func newGame() Game {
	legendImage := loadLegendImage()
	arena := NewArena()
	game := Game{
		arena:       &arena,
		legendImage: legendImage,
	}

	matchCh := make(chan *client.MatchInit)
	updateCh := make(chan *client.Update)
	sendUpdateCh := make(chan *client.Update)
	joinCh := make(chan *client.PlayerJoin)
	leaveCh := make(chan int)
	quit := make(chan bool)

	game.matchCh = matchCh
	game.updateCh = updateCh
	game.sendUpdateCh = sendUpdateCh
	game.joinCh = joinCh
	game.leaveCh = leaveCh
	game.quit = quit

	acceptedCh := make(chan *client.JoinAccepted)

	go client.Run(user.Name, acceptedCh, matchCh, updateCh, sendUpdateCh, joinCh, leaveCh)

	accepted := <-acceptedCh

	arena.SetOnCharacterMotion(game.onCharacterMotion)

	user.Id = accepted.Id

	log.Println("Accepted", accepted.Id)
	return game
}

func getSize() geo.Dimension {
	return geo.NewDimension(screenWidth, screenHeight)
}

func init() {
	loadBg()
	loadUser()
	dungeon.InitAssets()
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

func loadUser() {
	content, err := ioutil.ReadFile("../user.json")

	if err != nil {
		log.Fatal("Failed to read user")
	}

	if err := json.Unmarshal(content, &user); err != nil {
		log.Fatal("Failed to parse user")
	}
}

func loadLegendImage() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("./assets/keyboard_legend.png")

	if err != nil {
		log.Fatal(err)
	}
	return img
}

func remove(slice []*game.Diamond, s int) []*game.Diamond {
	return append(slice[:s], slice[s+1:]...)
}

// Place this here for now
var (
	mplusNormalFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    72,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}
