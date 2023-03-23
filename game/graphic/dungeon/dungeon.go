// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/core/geo"
	"dungeon-mst/core/graphic"
	"dungeon-mst/dungeon"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// TODO check for duplication on constants

const (
	horizontalUnitWidthPx  = 64
	horizontalUnitHeightPx = 12
	wallWidth              = horizontalUnitHeightPx
)

type DungeonBackgroundGraphic uint8

const (
	DungeonBackground DungeonBackgroundGraphic = iota
)

func (g DungeonBackgroundGraphic) Name() graphic.Name {
	return "dungeon_bg.png"
}

type DungeonBackgroundGraphics map[DungeonBackgroundGraphic]*graphic.Graphic

func LoadDungeonBackgroundGraphics(load graphic.Load) DungeonBackgroundGraphics {
	return DungeonBackgroundGraphics{
		DungeonBackground: load(DungeonBackground),
	}
}

type dungeonBackgroundDrawing struct {
	graphics EntityGraphics[DungeonBackgroundGraphic]
	rect     *geo.Rect
}

func (d dungeonBackgroundDrawing) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	img := d.graphics.Get(DungeonBackground)
	imgRect := image.Rect(0, 0, d.rect.Width()-2*wallWidth, d.rect.Height()-2*wallWidth)

	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Left()+wallWidth), float64(d.rect.Top()+wallWidth))
	screen.DrawImage(img.SubImage(imgRect).(*ebiten.Image), op)
}

func NewDungeonBackgroundDrawing(
	graphics EntityGraphics[DungeonBackgroundGraphic],
	rect *geo.Rect,
) graphic.Draw {
	return dungeonBackgroundDrawing{
		graphics: graphics,
		rect:     rect,
	}
}

type BrickGraphic uint8

const (
	Brick  BrickGraphic = iota
	BrickY BrickGraphic = iota
)

func (g BrickGraphic) Name() graphic.Name {
	return map[BrickGraphic]graphic.Name{
		Brick:  "brick.png",
		BrickY: "brick_y.png",
	}[g]
}

type BrickGraphics map[BrickGraphic]*graphic.Graphic

func LoadBrickGraphics(load graphic.Load) BrickGraphics {
	return BrickGraphics{
		Brick:  load(Brick),
		BrickY: load(BrickY),
	}
}

type brickDrawing struct {
	graphics EntityGraphics[BrickGraphic]
	barrier  *dungeon.Barrier
}

func (b brickDrawing) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	factor := b.barrier.Factor()
	wFactor := factor.Width
	hFactor := factor.Height
	blockWidth := horizontalUnitHeightPx
	brickImage := b.graphics.Get(Brick).Image
	brickYImage := b.graphics.Get(BrickY).Image

	// Draw Top
	op.GeoM.Reset()
	op.GeoM.Translate(b.barrier.GetTopTranslation())
	for i := 0; i < wFactor; i++ {
		screen.DrawImage(brickImage, op)
		op.GeoM.Translate(horizontalUnitWidthPx, 0)
	}

	// Draw Bottom
	op.GeoM.Reset()
	op.GeoM.Translate(b.barrier.GetBottomTranslation(blockWidth))
	for i := 0; i < wFactor; i++ {
		screen.DrawImage(brickImage, op)
		op.GeoM.Translate(horizontalUnitWidthPx, 0)
	}

	// Draw Left
	op.GeoM.Reset()
	op.GeoM.Translate(b.barrier.GetLeftTranslation())
	for i := 0; i < hFactor; i++ {
		screen.DrawImage(brickYImage, op)
		op.GeoM.Translate(0, horizontalUnitWidthPx)
	}

	// Draw Right
	op.GeoM.Reset()
	op.GeoM.Translate(b.barrier.GetRightTranslation(blockWidth))
	for i := 0; i < hFactor; i++ {
		screen.DrawImage(brickYImage, op)
		op.GeoM.Translate(0, horizontalUnitWidthPx)
	}
}

func NewBrickDrawing(
	graphics EntityGraphics[BrickGraphic],
	barrier *dungeon.Barrier,
) graphic.Draw {
	return brickDrawing{
		graphics: graphics,
		barrier:  barrier,
	}
}
