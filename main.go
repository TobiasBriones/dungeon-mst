package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

const windowWidth = 1280
const windowHeight = 720

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(*ebiten.Image) {
}

func (g *Game) Layout(int, int) (screenWidth int, screenHeight int) {
	return windowWidth, windowHeight
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Dungeon MST")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
