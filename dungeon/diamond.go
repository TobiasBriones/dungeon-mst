// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/core/geo"
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

type DiamondDimension struct {
	geo.Dimension
}

func NewDiamondDimension(width, height int) DiamondDimension {
	return DiamondDimension{geo.NewDimension(width, height)}
}

func NewDiamond(point geo.Point, dim DiamondDimension) Diamond {
	rect := geo.NewRect(
		point.X(),
		point.Y(),
		point.X()+dim.Width(),
		point.Y()+dim.Height(),
	)
	return Diamond{rect}
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
