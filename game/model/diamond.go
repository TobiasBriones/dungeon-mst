// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package model

import (
	"dungeon-mst/geo"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	DiamondWidthPx  = 32
	DiamondHeightPx = 26
)

var (
	diamondImage *ebiten.Image
)

type Diamond struct {
	rect geo.Rect
}

func (d *Diamond) Collides(rect *geo.Rect) bool {
	return d.rect.Intersects(rect)
}

func (d *Diamond) Update() {

}

func (d *Diamond) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(d.rect.Left()), float64(d.rect.Top()))
	screen.DrawImage(diamondImage, op)
}

func NewDiamond(point geo.Point) Diamond {
	rect := geo.NewRect(
		point.X(),
		point.Y(),
		point.X()+DiamondWidthPx,
		point.Y()+DiamondHeightPx,
	)
	return Diamond{
		rect: rect,
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
