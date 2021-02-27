/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	"dungeon-mst/ai"
	"dungeon-mst/model"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	_ "image/png"
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
)

type Game struct {
	runner model.Runner
}

func (g *Game) Update() error {
	var currentDungeon *model.Dungeon = nil
	var currentPaths []*model.Path

	for _, dungeon := range dungeons {
		if dungeon.InBounds(&g.runner.Rect) {
			currentDungeon = dungeon
			break
		}
	}
	for _, path := range paths {
		if path.InBounds(&g.runner.Rect) {
			currentPaths = append(currentPaths, path)
		}
	}
	g.runner.SetCurrentDungeon(currentDungeon)
	g.runner.SetCurrentPaths(currentPaths)

	if g.runner.IsOutSide() {
		g.runner.SetDungeon(dungeons[0])
	}

	g.runner.Update()

	// Generate random dungeons
	//for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
	//	if ebiten.IsKeyPressed(k) {
	//		dungeons = ai.GenerateDungeons(getSize())
	//		ai.GetPaths(dungeons)
	//	}
	//}
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
}

func NewGame() Game {
	runner := model.NewRunner()
	return Game{
		runner: runner,
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := NewGame()

	testRectIntersect()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Dungeon MST")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
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

func genSomeDungeons() []*model.Dungeon {
	d0 := model.NewDungeon(model.Point{}, model.DimensionFactor{Width: 1, Height: 1})
	d1 := model.NewDungeon(model.NewPoint(20, 540), model.DimensionFactor{Width: 4, Height: 1})
	d2 := model.NewDungeon(model.NewPoint(200, 140), model.DimensionFactor{Width: 3, Height: 2})
	d3 := model.NewDungeon(model.NewPoint(350, 90), model.DimensionFactor{Width: 4, Height: 1})
	return []*model.Dungeon{&d0, &d1, &d2, &d3}
}

// Neighbors will be replaced by doors
/*
func genSomeNeighbors(dungeons []*model.Dungeon) {
	dungeons[0].AddNeighbor(dungeons[1])
	dungeons[0].AddNeighbor(dungeons[2])

	dungeons[1].AddNeighbor(dungeons[2])

	dungeons[2].AddNeighbor(dungeons[3])
}
*/

func testRectIntersect() {
	r1 := model.NewRect(
		0,
		0,
		50,
		40,
	)
	r2 := model.NewRect(
		10,
		20,
		30,
		30,
	)
	r3 := model.NewRect(
		24,
		30,
		30,
		70,
	)
	r4 := model.NewRect(
		30,
		20,
		330,
		300,
	)
	r5 := model.NewRect(
		100,
		20,
		300,
		30,
	)
	r6 := model.NewRect(
		10,
		200,
		80,
		230,
	)

	if !r1.Intersects(&r2) {
		fmt.Println("FAILED R1-R2")
	}
	if !r2.Intersects(&r1) {
		fmt.Println("FAILED R1-R2")
	}
	if !r1.Intersects(&r3) {
		fmt.Println("FAILED R1-R3")
	}
	if !r3.Intersects(&r1) {
		fmt.Println("FAILED R1-R3")
	}
	if !r1.Intersects(&r4) {
		fmt.Println("FAILED R1-R4")
	}
	if !r4.Intersects(&r1) {
		fmt.Println("FAILED R1-R4")
	}
	if r1.Intersects(&r5) {
		fmt.Println("FAILED R1-R5")
	}
	if r5.Intersects(&r1) {
		fmt.Println("FAILED R1-R5")
	}
	if r1.Intersects(&r6) {
		fmt.Println("FAILED R1-R6")
	}
	if r6.Intersects(&r1) {
		fmt.Println("FAILED R1-R6")
	}
}
