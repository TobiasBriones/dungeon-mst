// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package graphic

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Graphic defines a unit of asset (i.e. an image) of the game.
type Graphic struct {
	*ebiten.Image
}

// Name defines the name of one asset. It should be the physical name of the
// image file.
type Name string

// NamedGraphic it defines a Graphic object that has a name.
//
// See Name.
type NamedGraphic interface {
	Name() Name
}
