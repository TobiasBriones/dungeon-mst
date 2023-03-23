// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"bytes"
	"dungeon-mst/core/graphic"
	"dungeon-mst/dungeon"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"image"
	"log"
)

const (
	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
)

type RunnerGraphic uint8

const (
	Runner RunnerGraphic = iota
)

func (g RunnerGraphic) Name() graphic.Name {
	return map[RunnerGraphic]graphic.Name{
		Runner: "runner.png",
	}[g]
}

type RunnerGraphics map[RunnerGraphic]*graphic.Graphic

func LoadRunnerGraphics() RunnerGraphics {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))

	if err != nil {
		log.Fatal(err)
	}
	runnerImage := ebiten.NewImageFromImage(img)

	return RunnerGraphics{Runner: &graphic.Graphic{Image: runnerImage}}
}

type runnerDrawing struct {
	graphics EntityGraphics[RunnerGraphic]
	*dungeon.Runner
}

func (r *runnerDrawing) Draw(screen *ebiten.Image) {
	x := r.Rect().Left()
	y := r.Rect().Top()
	op := &ebiten.DrawImageOptions{}
	i := (r.Count() / 5) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY
	rect := image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)
	fullImage := r.graphics.Get(Runner)

	op.GeoM.Scale(r.Scale(), r.Scale())
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(fullImage.SubImage(rect).(*ebiten.Image), op)
}

func NewRunnerDrawing(
	graphics EntityGraphics[RunnerGraphic],
	runner *dungeon.Runner,
) graphic.Draw {
	return &runnerDrawing{
		graphics: graphics,
		Runner:   runner,
	}
}
