// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/core/geo"
	"dungeon-mst/core/graphic"
	"math/rand"
	"time"
)

type BackgroundGraphic uint8

const (
	Background1 BackgroundGraphic = iota
	Background2
	Background3
)

func (g BackgroundGraphic) Name() graphic.Name {
	return map[BackgroundGraphic]graphic.Name{
		Background1: "bg_1.png",
		Background2: "bg_2.png",
		Background3: "bg_3.png",
	}[g]
}

type BackgroundGraphics map[BackgroundGraphic]*graphic.Graphic

func LoadBackgroundGraphics(load graphic.Load) BackgroundGraphics {
	return BackgroundGraphics{
		Background1: load(Background1),
		Background2: load(Background2),
		Background3: load(Background3),
	}
}

func NewBackgroundDrawing(
	graphics EntityGraphics[BackgroundGraphic],
	bg BackgroundGraphic,
) graphic.Draw {
	rect := geo.NewRect(0, 0, 1, 1) // Full size from origin
	return graphic.NewDrawing(graphics.Get(bg), &rect)
}

func GetRandomBackground() BackgroundGraphic {
	rand.Seed(time.Now().UnixNano())
	bgNumber := rand.Intn(3) + 1
	return BackgroundGraphic(bgNumber - 1)
}
