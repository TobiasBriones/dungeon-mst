// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package model

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

func NewImageFromAssets(name string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFile("./assets/" + name)

	if err != nil {
		log.Fatal(err)
	}
	return image
}
