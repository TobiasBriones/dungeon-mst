/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	"dungeon-mst/game/model"
	"dungeon-mst/geo"
	"dungeon-mst/mst"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

func NewRandomMatch() *model.Match {
	dimension := geo.NewDimension(screenWidth, screenHeight)
	return mst.NewRandomMatch(dimension)
}
