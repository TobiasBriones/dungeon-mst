// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package drawing

import (
	"dungeon-mst/core/geo"
	"dungeon-mst/core/graphic"
	"dungeon-mst/game/asset"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type pathDrawing struct {
	graphics asset.EntityGraphics[asset.PathGraphic]
	hRect    *geo.Rect
	vRect    *geo.Rect
}

func (p pathDrawing) Draw(screen *ebiten.Image) {
	p.drawHorizontalLine(screen)
	p.drawVerticalLine(screen)
}

func (p pathDrawing) drawHorizontalLine(screen *ebiten.Image) {
	rect := p.hRect
	img := p.graphics.Get(asset.Path)
	x := rect.Left()
	y := rect.Top()
	op := &ebiten.DrawImageOptions{}
	subRect := image.Rect(0, 0, rect.Width(), rect.Height())

	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img.SubImage(subRect).(*ebiten.Image), op)
}

func (p pathDrawing) drawVerticalLine(screen *ebiten.Image) {
	rect := p.vRect
	img := p.graphics.Get(asset.PathY)
	x := rect.Left()
	y := rect.Top()
	op := &ebiten.DrawImageOptions{}
	subRect := image.Rect(0, 0, rect.Width(), rect.Height())

	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img.SubImage(subRect).(*ebiten.Image), op)
}

func NewPathDrawing(
	graphics asset.EntityGraphics[asset.PathGraphic],
	hRect *geo.Rect,
	vRect *geo.Rect,
) graphic.Draw {
	return pathDrawing{
		graphics: graphics,
		hRect:    hRect,
		vRect:    vRect,
	}
}
