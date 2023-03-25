// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package asset

import (
	"bytes"
	"dungeon-mst/core/graphic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"image"
	"log"
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
