/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

const (
	PathWidthPx = 36
)

type Path struct {
	hLine Line
	hRect Rect
	vLine Line
	vRect Rect
}

func (p *Path) InBounds(rect *Rect) bool {
	return p.hRect.InBounds(rect) || p.vRect.InBounds(rect)
}

func (p *Path) CanMoveTowards(movement Movement, rect *Rect) bool {
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

func NewPath(hl Line, vl Line) Path {
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
	sw := PathWidthPx / 2
	hRect := NewRect(
		hl.p1.X()-sw,
		hl.p1.Y()-sw,
		hl.p2.X()+sw,
		hl.p1.Y()+sw,
	)
	vRect := NewRect(
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

func (p *PathJSON) ToPath() *Path {
	path := NewPath(*p.HLineJSON.ToLine(), *p.VLineJSON.ToLine())
	return &path
}

type Line struct {
	p1 Point
	p2 Point
}

func (l *Line) IsDegenerate() bool {
	return Distance(l.p1, l.p2) == 0
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
