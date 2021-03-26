/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

const (
	DiamondWidthPx  = 32
	DiamondHeightPx = 26
)

type Diamond struct {
	rect Rect
}

func (d *Diamond) Collides(rect *Rect) bool {
	return d.rect.Intersects(rect)
}

func NewDiamond(point Point) Diamond {
	rect := Rect{
		left:   point.X(),
		top:    point.Y(),
		right:  point.X() + DiamondWidthPx,
		bottom: point.Y() + DiamondHeightPx,
	}
	return Diamond{
		rect: rect,
	}
}

type DiamondJSON struct {
	*PointJSON
}

func (d *DiamondJSON) ToDiamond() *Diamond {
	diamond := NewDiamond(*d.PointJSON.ToPoint())
	return &diamond
}

func NewDiamondJSON(d *Diamond) *DiamondJSON {
	point := &Point{d.rect.left, d.rect.top}
	return &DiamondJSON{NewPointJSON(point)}
}
