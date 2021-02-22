package main

import (
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
)

type Game struct {
	runner model.Runner
}

func (g *Game) Update() error {
	g.runner.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bgImage, nil)

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
	//dungeons = ai.GenerateDungeons(getSize())
	dungeons = genSomeDungeons()

	genSomeNeighbors(dungeons)
	//ai.GetNeighborhoods(dungeons)
	testRectIntersect()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Dungeon MST")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(dungeons))
}

func getSize() model.Dimension {
	return model.Dimension{Width: screenWidth, Height: screenHeight}
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

func genSomeDungeons() []*model.Dungeon {
	d0 := model.NewDungeon(model.Point{}, model.DimensionFactor{Width: 1, Height: 1})
	d1 := model.NewDungeon(model.Point{X: 20, Y: 540}, model.DimensionFactor{Width: 4, Height: 1})
	d2 := model.NewDungeon(model.Point{X: 200, Y: 140}, model.DimensionFactor{Width: 3, Height: 2})
	d3 := model.NewDungeon(model.Point{X: 350, Y: 90}, model.DimensionFactor{Width: 4, Height: 1})
	return []*model.Dungeon{&d0, &d1, &d2, &d3}
}

func genSomeNeighbors(dungeons []*model.Dungeon) {
	dungeons[0].AddNeighbor(dungeons[1])
	dungeons[0].AddNeighbor(dungeons[2])

	dungeons[1].AddNeighbor(dungeons[0])
	dungeons[1].AddNeighbor(dungeons[2])

	dungeons[2].AddNeighbor(dungeons[0])
	dungeons[2].AddNeighbor(dungeons[1])
	dungeons[2].AddNeighbor(dungeons[3])

	dungeons[3].AddNeighbor(dungeons[2])
}

func testRectIntersect() {
	r1 := model.Rect{
		Left:   0,
		Top:    0,
		Right:  50,
		Bottom: 40,
	}
	r2 := model.Rect{
		Left:   10,
		Top:    20,
		Right:  30,
		Bottom: 30,
	}
	r3 := model.Rect{
		Left:   24,
		Top:    30,
		Right:  30,
		Bottom: 70,
	}
	r4 := model.Rect{
		Left:   30,
		Top:    20,
		Right:  330,
		Bottom: 300,
	}
	r5 := model.Rect{
		Left:   100,
		Top:    20,
		Right:  300,
		Bottom: 30,
	}
	r6 := model.Rect{
		Left:   10,
		Top:    200,
		Right:  80,
		Bottom: 230,
	}

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
