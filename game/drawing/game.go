// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package drawing

import (
	"dungeon-mst/core/graphic"
	"dungeon-mst/game/asset"
	"github.com/hajimehoshi/ebiten/v2"
)

type gameDrawing struct {
	graphics asset.EntityGraphics[asset.GameGraphic]
}

func (d gameDrawing) Draw(screen *ebiten.Image) {
	d.drawLegend(screen)
}

func (d gameDrawing) drawLegend(screen *ebiten.Image) {
	screen.DrawImage(d.graphics.Get(asset.Legend).Image, nil)
}

func NewGameDrawing(
	graphics asset.EntityGraphics[asset.GameGraphic],
) graphic.Draw {
	return gameDrawing{graphics: graphics}
}
