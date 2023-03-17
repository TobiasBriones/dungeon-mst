// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package graphic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Graphic struct {
	*ebiten.Image
}

func NewGraphicFromAssets(name string) *Graphic {
	image, _, err := ebitenutil.NewImageFromFile("assets/" + name)

	if err != nil {
		log.Fatal(err)
	}
	return &Graphic{image}
}
