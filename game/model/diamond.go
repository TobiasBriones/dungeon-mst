/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

import (
	"dungeon-mst/geo"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	DiamondWidthPx  = 32
	DiamondHeightPx = 26
)

type Diamond struct {
	rect  geo.Rect
	image *ebiten.Image
}

func (d *Diamond) Collides(rect *geo.Rect) bool {
	return d.rect.Intersects(rect)
}

func (d *Diamond) Update() {

}

func (d *Diamond) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(d.rect.Left()), float64(d.rect.Top()))
	screen.DrawImage(d.image, op)
}

func NewDiamond(point geo.Point) Diamond {
	rect := geo.NewRect(
		point.X(),
		point.Y(),
		point.X()+DiamondWidthPx,
		point.Y()+DiamondHeightPx,
	)
	image := NewImageFromAssets("diamond.png")
	return Diamond{
		rect:  rect,
		image: image,
	}
}

type DiamondJSON struct {
	*PointJSON
}

func (d *DiamondJSON) ToDiamond() *Diamond {
	diamond := NewDiamond(*d.PointJSON.ToPoint())
	return &diamond
}

func NewDiamondJSON(d *Diamond) *DiamondJSON {
	point := geo.NewPoint(d.rect.Left(), d.rect.Top())
	return &DiamondJSON{NewPointJSON(&point)}
}
