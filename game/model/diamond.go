/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const (
	DiamondWidthPx  = 32
	DiamondHeightPx = 26
)

type Diamond struct {
	rect  Rect
	image *ebiten.Image
}

func (d *Diamond) Collides(rect *Rect) bool {
	return d.rect.Intersects(rect)
}

func (d *Diamond) Update() {

}

func (d *Diamond) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(d.rect.Left()), float64(d.rect.Top()))
	screen.DrawImage(d.image, op)
}

func NewDiamond(point Point) Diamond {
	rect := Rect{
		left:   point.X(),
		top:    point.Y(),
		right:  point.X() + DiamondWidthPx,
		bottom: point.Y() + DiamondHeightPx,
	}
	image, _, err := ebitenutil.NewImageFromFile("./assets/diamond.png")

	if err != nil {
		log.Fatal(err)
	}
	return Diamond{
		rect:  rect,
		image: image,
	}
}
