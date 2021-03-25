/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

import (
	_ "image/png"
	"math/rand"
)

const (
	horizontalUnitWidthPx  = 64
	horizontalUnitHeightPx = 12
	wallWidth              = horizontalUnitHeightPx
)

type Dungeon struct {
	rect    Rect
	barrier Barrier
}

func (d *Dungeon) Width() int {
	return d.rect.Width()
}

func (d *Dungeon) Height() int {
	return d.rect.Height()
}

func (d *Dungeon) Cx() int {
	return d.rect.Cx()
}

func (d *Dungeon) Cy() int {
	return d.rect.Cy()
}

func (d *Dungeon) Center() Point {
	return NewPoint(d.Cx(), d.Cy())
}

func (d *Dungeon) GetPathFor(dungeon *Dungeon) *Path {
	center := d.Center()
	hp1x := min(center.X(), dungeon.Cx())
	hp2x := max(center.X(), dungeon.Cx())
	hpy := dungeon.Cy()
	hl := Line{
		p1: Point{hp1x, hpy},
		p2: Point{hp2x, hpy},
	}

	vp1x := center.X()
	vp1y := min(center.Y(), dungeon.Cy())
	vp2y := max(center.Y(), dungeon.Cy())
	vl := Line{
		p1: Point{vp1x, vp1y},
		p2: Point{vp1x, vp2y},
	}

	path := NewPath(hl, vl)
	return &path
}

func (d *Dungeon) Intersects(rect *Rect) bool {
	return d.rect.Intersects(rect)
}

func (d *Dungeon) InBounds(rect *Rect) bool {
	return d.rect.InBounds(rect)
}

func (d *Dungeon) CanMoveTowards(movement Movement, rect *Rect) bool {
	if !d.InBounds(rect) {
		return true
	}
	return !d.barrier.WillCollide(movement, rect)
}

func (d *Dungeon) RandomPoint(p int) Point {
	x := rand.Intn(d.Width()-wallWidth*2-p) + d.rect.Left() + wallWidth
	y := rand.Intn(d.Height()-wallWidth*2-p) + d.rect.Top() + wallWidth
	return Point{x, y}
}

func NewDungeon(p0 Point, factor DimensionFactor) Dungeon {
	x0 := p0.X()
	y0 := p0.Y()
	w := factor.Width * horizontalUnitWidthPx
	h := factor.Height * horizontalUnitWidthPx
	rect := NewRect(x0, y0, x0+w, y0+h)
	barrier := NewBarrier(rect, factor)
	return Dungeon{
		rect,
		barrier,
	}
}

type DimensionFactor struct {
	Width  int
	Height int
}

type Wall struct {
	rect Rect
}

type Barrier struct {
	factor     DimensionFactor
	leftWall   Wall
	topWall    Wall
	rightWall  Wall
	bottomWall Wall
}

func (b *Barrier) WillCollide(movement Movement, objRect *Rect) bool {
	return WillCollide(movement, &b.leftWall.rect, objRect) ||
		WillCollide(movement, &b.topWall.rect, objRect) ||
		WillCollide(movement, &b.rightWall.rect, objRect) ||
		WillCollide(movement, &b.bottomWall.rect, objRect)
}

func NewBarrier(rect Rect, factor DimensionFactor) Barrier {
	return Barrier{
		factor: factor,
		leftWall: Wall{
			rect: NewRect(
				rect.Left(),
				rect.Top(),
				rect.Left()+wallWidth,
				rect.Bottom(),
			),
		},
		topWall: Wall{
			rect: NewRect(
				rect.Left(),
				rect.Top(),
				rect.Right(),
				rect.Top()+wallWidth,
			),
		},
		rightWall: Wall{
			rect: NewRect(
				rect.Right()-wallWidth,
				rect.Top(),
				rect.Right(),
				rect.Bottom(),
			),
		},
		bottomWall: Wall{
			rect: NewRect(
				rect.Left(),
				rect.Bottom()-wallWidth,
				rect.Right(),
				rect.Bottom(),
			),
		},
	}
}

func GetDungeonHorizontalUnitSize() Dimension {
	return NewDimension(
		horizontalUnitWidthPx,
		horizontalUnitHeightPx,
	)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
