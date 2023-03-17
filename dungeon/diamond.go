// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/geo"
)

type Diamond struct {
	rect geo.Rect
}

func (d *Diamond) Rect() *geo.Rect {
	return &d.rect
}

func (d *Diamond) Collides(rect *geo.Rect) bool {
	return d.rect.Intersects(rect)
}

func (d *Diamond) Update() {}

func NewDiamond(point geo.Point, width, height int) Diamond {
	rect := geo.NewRect(
		point.X(),
		point.Y(),
		point.X()+width,
		point.Y()+height,
	)
	return Diamond{
		rect: rect,
	}
}

type DiamondJSON struct {
	*RectJSON
}

func (d *DiamondJSON) ToDiamond() *Diamond {
	diamond := Diamond{*d.RectJSON.ToRect()}
	return &diamond
}

func NewDiamondJSON(d *Diamond) *DiamondJSON {
	return &DiamondJSON{NewRectJSON(&d.rect)}
}
