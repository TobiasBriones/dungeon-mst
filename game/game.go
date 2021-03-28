/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	"encoding/json"
	"game/client"
	"game/model"
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
	match         *model.Match
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

func (g *Game) SetMatch(value *model.Match) {
	g.match = value

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
		if g.arena.checkDiamondCollision(diamond) {
			log.Println("Diamond collision", i, "len:", len(g.match.Diamonds))
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
		PointJSON:    *model.NewPointJSON(&position),
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

			for _, player := range m.Players {
				game.arena.PushRemotePlayer(player.Id, player.Name, player.Score)
			}
			go game.watchRemainingTime()
		}
	}()
	go func() {
		for {
			u := <-game.updateCh

			if u.Id == user.Id {
				continue
			}
			//log.Println("Receiving update for player:", u.Id)
			game.arena.SetRemotePlayerPosition(u.Id, u.PointJSON.ToPoint())

			if u.DiamondIndex != -1 {
				game.match.Diamonds = remove(game.match.Diamonds, u.DiamondIndex)

				game.arena.SetRemotePlayerScore(u.Id)
			}

			//game.arena.PushRemotePlayerInput(u.Id, u.Move)
		}
	}()
	go func() {
		for {
			j := <-game.joinCh

			if j.Id == user.Id {
				continue
			}
			log.Println("Joining player:", j.Id)
			game.arena.PushRemotePlayer(j.Id, j.Name, 0)
		}
	}()

	go func() {
		for {
			lid := <-game.leaveCh

			game.arena.RemoveRemotePlayer(lid)
		}
	}()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

func newGame() Game {
	arena := NewArena(user.Name)
	legendImage := loadLegendImage()
	game := Game{
		arena:       &arena,
		legendImage: legendImage,
	}

	game.arena.SetOnCharacterMotion(game.onCharacterMotion)

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
	arena.player.Id = accepted.Id
	user.Id = accepted.Id

	log.Println("Accepted", accepted.Id)
	return game
}

func getSize() model.Dimension {
	return model.NewDimension(screenWidth, screenHeight)
}

func init() {
	loadBg()
	loadUser()
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

func remove(slice []*model.Diamond, s int) []*model.Diamond {
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
