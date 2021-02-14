package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

const windowWidth = 640
const windowHeight = 480

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(*ebiten.Image) {
}

func (g *Game) Layout(int, int) (screenWidth int, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Your game's title")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
