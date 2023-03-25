// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package asset

import (
	"dungeon-mst/core/graphic"
)

// GameGraphic defines general game graphics
type GameGraphic uint8

const (
	Legend GameGraphic = iota
)

func (g GameGraphic) Name() graphic.Name {
	return map[GameGraphic]graphic.Name{
		Legend: "keyboard_legend.png",
	}[g]
}

type GameGraphics map[GameGraphic]*graphic.Graphic

func LoadGameGraphics(load graphic.Load) GameGraphics {
	return GameGraphics{Legend: load(Legend)}
}
