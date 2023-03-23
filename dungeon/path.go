// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package dungeon

import (
	"dungeon-mst/core/geo"
)

type Path struct {
	hLine Line
	hRect geo.Rect
	vLine Line
	vRect geo.Rect
}

func (p *Path) HRect() *geo.Rect {
	return &p.hRect
}

func (p *Path) VRect() *geo.Rect {
	return &p.vRect
}

func (p *Path) InBounds(rect *geo.Rect) bool {
	return p.hRect.InBounds(rect) || p.vRect.InBounds(rect)
}

func (p *Path) CanMoveTowards(movement Movement, rect *geo.Rect) bool {
	if !p.InBounds(rect) {
		return true
	}
	mx := false
	my := false

	if p.hRect.InBounds(rect) {
		mx = CheckMovement(movement, rect, p.hRect)
	}
	if p.vRect.InBounds(rect) {
		my = CheckMovement(movement, rect, p.vRect)
	}
	return mx || my
}

type PathDimension = uint

func NewPath(hl Line, vl Line, dim PathDimension) Path {
	if hl.IsDegenerate() || vl.IsDegenerate() {
		panic("Lines cannot be degenerate")
	}
	if !(hl.IsHorizontal() && vl.IsVertical()) {
		panic("The first line must be horizontal and the other vertical")
	}
	if !hl.p1.Equals(&vl.p1) && !hl.p1.Equals(&vl.p2) &&
		!hl.p2.Equals(&vl.p1) && !hl.p2.Equals(&vl.p2) {
		panic("One of the horizontal line start/end point must be the beginning of the vertical line")
	}
	if hl.p1.X() > hl.p2.X() {
		panic("The point 1 of the horizontal line must be the lowest")
	}
	if vl.p1.Y() > vl.p2.Y() {
		panic("The point 1 of the vertical line must be the lowest")
	}
	sw := int(dim) / 2
	hRect := geo.NewRect(
		hl.p1.X()-sw,
		hl.p1.Y()-sw,
		hl.p2.X()+sw,
		hl.p1.Y()+sw,
	)
	vRect := geo.NewRect(
		vl.p1.X()-sw,
		vl.p1.Y()-sw,
		vl.p1.X()+sw,
		vl.p2.Y()+sw,
	)
	return Path{hl, hRect, vl, vRect}
}

type PathJSON struct {
	HLineJSON LineJSON
	VLineJSON LineJSON
}

func (p *PathJSON) ToPath(dim PathDimension) *Path {
	path := NewPath(*p.HLineJSON.ToLine(), *p.VLineJSON.ToLine(), dim)
	return &path
}

func NewPathJSON(p *Path) *PathJSON {
	return &PathJSON{
		HLineJSON: *NewLineJSON(&p.hLine),
		VLineJSON: *NewLineJSON(&p.vLine),
	}
}

type Line struct {
	p1 geo.Point
	p2 geo.Point
}

func (l *Line) IsDegenerate() bool {
	return geo.Distance(l.p1, l.p2) == 0
}

func (l *Line) IsHorizontal() bool {
	return l.p1.Y() == l.p2.Y()
}

func (l *Line) IsVertical() bool {
	return l.p1.X() == l.p2.X()
}

type LineJSON struct {
	P1JSON PointJSON
	P2JSON PointJSON
}

func (l *LineJSON) ToLine() *Line {
	line := &Line{
		p1: *l.P1JSON.ToPoint(),
		p2: *l.P2JSON.ToPoint(),
	}
	return line
}

func NewLineJSON(l *Line) *LineJSON {
	return &LineJSON{
		P1JSON: *NewPointJSON(&l.p1),
		P2JSON: *NewPointJSON(&l.p2),
	}
}
