// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package drawing

import (
	"dungeon-mst/core/geo"
	"dungeon-mst/core/graphic"
	"dungeon-mst/dungeon"
	"dungeon-mst/game/asset"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// TODO check for duplication on constants

const (
	horizontalUnitWidthPx  = 64
	horizontalUnitHeightPx = 12
	wallWidth              = horizontalUnitHeightPx
)

type dungeonBackgroundDrawing struct {
	graphics asset.EntityGraphics[asset.DungeonBackgroundGraphic]
	rect     *geo.Rect
}

func (d dungeonBackgroundDrawing) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	img := d.graphics.Get(asset.DungeonBackground)
	imgRect := image.Rect(0, 0, d.rect.Width()-2*wallWidth, d.rect.Height()-2*wallWidth)

	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Left()+wallWidth), float64(d.rect.Top()+wallWidth))
	screen.DrawImage(img.SubImage(imgRect).(*ebiten.Image), op)
}

func NewDungeonBackgroundDrawing(
	graphics asset.EntityGraphics[asset.DungeonBackgroundGraphic],
	rect *geo.Rect,
) graphic.Draw {
	return dungeonBackgroundDrawing{
		graphics: graphics,
		rect:     rect,
	}
}

type brickDrawing struct {
	graphics asset.EntityGraphics[asset.BrickGraphic]
	barrier  *dungeon.Barrier
}

func (b brickDrawing) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	factor := b.barrier.Factor()
	wFactor := factor.Width
	hFactor := factor.Height
	blockWidth := horizontalUnitHeightPx
	brickImage := b.graphics.Get(asset.Brick).Image
	brickYImage := b.graphics.Get(asset.BrickY).Image

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
	graphics asset.EntityGraphics[asset.BrickGraphic],
	barrier *dungeon.Barrier,
) graphic.Draw {
	return brickDrawing{
		graphics: graphics,
		barrier:  barrier,
	}
}
