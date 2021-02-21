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
	dungeons []model.Dungeon
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
	drawSomeDungeons(screen)
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

func drawSomeDungeons(screen *ebiten.Image) {
	for _, dungeon := range dungeons {
		dungeon.Draw(screen)
	}
}

func genSomeDungeons() []model.Dungeon {
	return []model.Dungeon{
		model.NewDungeon(model.Point{}, 1, 1),
		model.NewDungeon(model.Point{X: 20, Y: 540}, 4, 1),
		model.NewDungeon(model.Point{X: 200, Y: 140}, 3, 2),
		model.NewDungeon(model.Point{X: 350, Y: 90}, 4, 1),
	}
}
