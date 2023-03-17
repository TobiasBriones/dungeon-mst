// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package graphic

import "dungeon-mst/dungeon"

type Graphics struct {
	graphics map[dungeon.Entity]*Graphic
}

func (g *Graphics) Get(entity dungeon.Entity) *Graphic {
	return g.graphics[entity]
}

// LoadGraphics loads the graphic assets of the game into memory.
// These graphics only need to be loaded once, and be reused when drawing.
// That is, there's a 1:n relation between a graphic and the instances of the
// same entities that appear on the game.
func LoadGraphics() *Graphics {
	return &Graphics{map[dungeon.Entity]*Graphic{
		dungeon.DiamondEntity: newGraphicFromDungeon(dungeon.DiamondEntity),
	}}
}

const (
	DiamondWidthPx  = 32
	DiamondHeightPx = 26
)

var (
	imageNameMatch = map[dungeon.Entity]string{
		dungeon.DiamondEntity: "diamond.png",
	}
)

func newGraphicFromDungeon(d dungeon.Entity) *Graphic {
	name := imageNameMatch[d]
	return NewGraphicFromAssets(name)
}
