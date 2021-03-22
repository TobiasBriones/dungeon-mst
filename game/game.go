/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package game

import (
	"dungeon-mst/ai"
	"dungeon-mst/model"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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
	bgImage  *ebiten.Image
	dungeons []*model.Dungeon
	paths    []*model.Path
	arena    Arena
)

type Game struct {
	runner      model.Runner
	count       int
	legendImage *ebiten.Image
}

func (g *Game) Update() error {
	g.count++

	// local player input
	input := model.MoveNone
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			switch k {
			case ebiten.KeyUp, ebiten.KeyW:
				input = model.MoveDirTop
			case ebiten.KeyDown, ebiten.KeyS:
				input = model.MoveDirBottom
			case ebiten.KeyLeft, ebiten.KeyA:
				input = model.MoveDirLeft
			case ebiten.KeyRight, ebiten.KeyD:
				input = model.MoveDirRight
			}
		}
	}
	g.runner.SetInput(input)
	//

	updatePlayer(&g.runner)
	arena.Update(updatePlayer)

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

	for _, dungeon := range dungeons {
		dungeon.DrawBarrier(screen)
	}
	for _, path := range paths {
		path.Draw(screen)
	}
	for _, dungeon := range dungeons {
		dungeon.Draw(screen)
	}

	g.runner.Draw(screen)

	// Draw legend image
	screen.DrawImage(g.legendImage, nil)

	// Draw remote players
	arena.Draw(screen)
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func Run() {
	game := newGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Dungeon MST")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

func newGame() Game {
	runner := model.NewRunner()
	legendImage := loadLegendImage()
	return Game{
		runner:      runner,
		legendImage: legendImage,
	}
}

func getSize() model.Dimension {
	return model.NewDimension(screenWidth, screenHeight)
}

func init() {
	loadBg()
	dungeons = ai.GenerateDungeons(getSize())
	//dungeons = genSomeDungeons()

	//genSomeNeighbors(dungeons)
	paths = ai.GetPaths(dungeons)
	arena = NewArena()
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

func updatePlayer(runner *model.Runner) {
	var currentDungeon *model.Dungeon = nil
	var currentPaths []*model.Path

	for _, dungeon := range dungeons {
		if dungeon.InBounds(&runner.Rect) {
			currentDungeon = dungeon
			break
		}
	}
	for _, path := range paths {
		if path.InBounds(&runner.Rect) {
			currentPaths = append(currentPaths, path)
		}
	}
	runner.SetCurrentDungeon(currentDungeon)
	runner.SetCurrentPaths(currentPaths)

	if runner.IsOutSide() {
		runner.SetDungeon(dungeons[0])
	}

	runner.Update()
}

func reset() {
	dungeons = ai.GenerateDungeons(getSize())
	paths = ai.GetPaths(dungeons)
}
