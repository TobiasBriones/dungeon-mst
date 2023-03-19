// Copyright (c) 2023 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/dungeon"
	"dungeon-mst/game/graphic"
)

type PathGraphic uint8

const (
	Path PathGraphic = iota
	PathY
)

func (g PathGraphic) Name() graphic.Name {
	return map[PathGraphic]graphic.Name{
		Path:  "path.png",
		PathY: "path_y.png",
	}[g]
}

type PathGraphics map[PathGraphic]*graphic.Graphic

func LoadPathGraphics(load graphic.Load) PathGraphics {
	return PathGraphics{
		Path:  load(Path),
		PathY: load(PathY),
	}
}

func PathSize() dungeon.PathDimension {
	return 36
}
