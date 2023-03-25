// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package drawing

import (
	"dungeon-mst/core/graphic"
	"dungeon-mst/dungeon"
	"dungeon-mst/game/asset"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

const (
	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
)

type runnerDrawing struct {
	graphics asset.EntityGraphics[asset.RunnerGraphic]
	*dungeon.Runner
}

func (r *runnerDrawing) Draw(screen *ebiten.Image) {
	x := r.Rect().Left()
	y := r.Rect().Top()
	op := &ebiten.DrawImageOptions{}
	i := (r.Count() / 5) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY
	rect := image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)
	fullImage := r.graphics.Get(asset.Runner)

	op.GeoM.Scale(r.Scale(), r.Scale())
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(fullImage.SubImage(rect).(*ebiten.Image), op)
}

func NewRunnerDrawing(
	graphics asset.EntityGraphics[asset.RunnerGraphic],
	runner *dungeon.Runner,
) graphic.Draw {
	return &runnerDrawing{
		graphics: graphics,
		Runner:   runner,
	}
}
