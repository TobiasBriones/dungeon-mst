// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package graphic

import (
	"dungeon-mst/core/geo"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

// Graphic defines a unit of graphic (i.e. an image) of the game.
type Graphic struct {
	*ebiten.Image
}

// LoadGraphicFromAssets loads the given asset from the "assets" directory where
// the program is being executed.
func LoadGraphicFromAssets(name string) *Graphic {
	image, _, err := ebitenutil.NewImageFromFile("assets/" + name)

	if err != nil {
		log.Fatal(err)
	}
	return &Graphic{image}
}

// Name defines the name of one graphic. It should be the physical name of the
// image file.
type Name string

// NamedGraphic it defines a Graphic object that has a name.
//
// See Name.
type NamedGraphic interface {
	Name() Name
}

// Load defines a func type that loads a NamedGraphic from the FS to an
// in-memory high-level Graphic.
type Load func(g NamedGraphic) *Graphic

// Draw defines an object that draws itself on a given canvas like the game
// screen.
type Draw interface {
	Draw(screen *ebiten.Image)
}

// Drawing Defines a simple drawable object that has one Graphic and a geo.Rect
// as position model.
//
// See Draw.
type Drawing struct {
	*Graphic
	*geo.Rect
}

func NewDrawing(graphic *Graphic, rect *geo.Rect) Drawing {
	return Drawing{
		Graphic: graphic,
		Rect:    rect,
	}
}

func (d Drawing) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(d.Left()), float64(d.Top()))
	screen.DrawImage(d.Image, op)
}
