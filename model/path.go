/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"log"
)

const (
	PathWidthPx = 36
)

var (
	pathImage  = getPathImage()
	pathYImage = getPathYImage()
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

func (p *Path) Draw(screen *ebiten.Image) {
	p.drawHorizontalLine(screen)
	p.drawVerticalLine(screen)
}

func (p *Path) drawHorizontalLine(screen *ebiten.Image) {
	rect := p.hRect
	x := rect.Left()
	y := rect.Top()
	op := &ebiten.DrawImageOptions{}
	subRect := image.Rect(0, 0, rect.Width(), rect.Height())

	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(pathImage.SubImage(subRect).(*ebiten.Image), op)
}

func (p *Path) drawVerticalLine(screen *ebiten.Image) {
	rect := p.vRect
	x := rect.Left()
	y := rect.Top()
	op := &ebiten.DrawImageOptions{}
	subRect := image.Rect(0, 0, rect.Width(), rect.Height())

	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(pathYImage.SubImage(subRect).(*ebiten.Image), op)
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

func getPathImage() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("./assets/path.png")

	if err != nil {
		log.Fatal(err)
	}
	return img
}

func getPathYImage() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("./assets/path_y.png")

	if err != nil {
		log.Fatal(err)
	}
	return img
}
