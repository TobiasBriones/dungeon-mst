// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/dungeon"
	"dungeon-mst/game/graphic"
	"dungeon-mst/geo"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type PathGraphic uint8

const (
	Path PathGraphic = iota
	PathY
)

func (g PathGraphic) Name() graphic.Name {
	return map[PathGraphic]graphic.Name{
		Path:  "path.png",
		PathY: "path_y.png",
	}[g]
}

type PathGraphics map[PathGraphic]*graphic.Graphic

func LoadPathGraphics(load graphic.Load) PathGraphics {
	return PathGraphics{
		Path:  load(Path),
		PathY: load(PathY),
	}
}

type pathDrawing struct {
	graphics EntityGraphics[PathGraphic]
	hRect    *geo.Rect
	vRect    *geo.Rect
}

func (p pathDrawing) Draw(screen *ebiten.Image) {
	p.drawHorizontalLine(screen)
	p.drawVerticalLine(screen)
}

func (p pathDrawing) drawHorizontalLine(screen *ebiten.Image) {
	rect := p.hRect
	img := p.graphics.Get(Path)
	x := rect.Left()
	y := rect.Top()
	op := &ebiten.DrawImageOptions{}
	subRect := image.Rect(0, 0, rect.Width(), rect.Height())

	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img.SubImage(subRect).(*ebiten.Image), op)
}

func (p pathDrawing) drawVerticalLine(screen *ebiten.Image) {
	rect := p.vRect
	img := p.graphics.Get(PathY)
	x := rect.Left()
	y := rect.Top()
	op := &ebiten.DrawImageOptions{}
	subRect := image.Rect(0, 0, rect.Width(), rect.Height())

	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img.SubImage(subRect).(*ebiten.Image), op)
}

func NewPathDrawing(
	graphics EntityGraphics[PathGraphic],
	hRect *geo.Rect,
	vRect *geo.Rect,
) graphic.Draw {
	return pathDrawing{
		graphics: graphics,
		hRect:    hRect,
		vRect:    vRect,
	}
}

func PathSize() dungeon.PathDimension {
	return 36
}
