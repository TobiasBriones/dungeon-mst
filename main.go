package main

import (
	"dungeon-mst/model"
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
	bgImage *ebiten.Image
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

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Dungeon MST")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
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
	dungeons := genSomeDungeons()

	for _, dungeon := range dungeons {
		dungeon.Draw(screen)
	}
}

func genSomeDungeons() []model.Dungeon {
	return []model.Dungeon{
		model.NewDungeon(0, 0, 1, 1),
		model.NewDungeon(20, 540, 4, 1),
		model.NewDungeon(200, 140, 3, 2),
		model.NewDungeon(350, 90, 4, 1),
	}
}
