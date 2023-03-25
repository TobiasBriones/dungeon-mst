// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package graphic

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const (
	AssetsDirPath = "asset/assets"
)

func GetAssetPath(name string) string {
	return AssetsDirPath + name
}

// LoadGraphicFromAssets loads the given asset from the "asset/assets" directory
// where the program is being executed.
// Directory "asset" should be the model representation of the physical assets
// in "asset/assets", i.e. a Go package that maps those graphics to an
// application Graphic.
func LoadGraphicFromAssets(name string) *Graphic {
	path := GetAssetPath(name)
	image, _, err := ebitenutil.NewImageFromFile(path)

	if err != nil {
		log.Fatal(err)
	}
	return &Graphic{image}
}

// Load defines a func type that loads a NamedGraphic from the FS to an
// in-memory high-level Graphic.
type Load func(g NamedGraphic) *Graphic
