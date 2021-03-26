/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

import (
	"math"
)

type Point struct {
	x int
	y int
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) Y() int {
	return p.y
}

func (p *Point) Equals(point *Point) bool {
	return p.X() == point.X() && p.Y() == point.Y()
}

func NewPoint(x int, y int) Point {
	if x < 0 || y < 0 {
		panic("Point coordinates must be non-negative integer numbers")
	}
	return Point{x, y}
}

type PointJSON struct {
	X int
	Y int
}

func (p *PointJSON) ToPoint() *Point {
	point := NewPoint(p.X, p.Y)
	return &point

}

func NewPointJSON(p *Point) *PointJSON {
	return &PointJSON{
		p.x, p.y,
	}
}

type PointPair struct {
	PointA Point
	PointB Point
}

type Dimension struct {
	width  int
	height int
}

func (d *Dimension) Width() int {
	return d.width
}

func (d *Dimension) Height() int {
	return d.height
}

func (d *Dimension) SemiWidth() int {
	return d.Width() / 2
}

func (d *Dimension) SemiHeight() int {
	return d.Height() / 2
}

func NewDimension(width int, height int) Dimension {
	if width <= 0 || height <= 0 {
		panic("Width and Height must be positive integer numbers")
	}
	return Dimension{width, height}
}

type Rect struct {
	left   int
	top    int
	right  int
	bottom int
}

func (r *Rect) Left() int {
	return r.left
}

func (r *Rect) Top() int {
	return r.top
}

func (r *Rect) Right() int {
	return r.right
}

func (r *Rect) Bottom() int {
	return r.bottom
}

func (r *Rect) Width() int {
	return r.Right() - r.Left()
}

func (r *Rect) SemiWidth() int {
	return r.Width() / 2
}

func (r *Rect) Height() int {
	return r.Bottom() - r.Top()
}

func (r *Rect) SemiHeight() int {
	return r.Height() / 2
}

func (r *Rect) Cx() int {
	return r.Left() + r.SemiWidth()
}

func (r *Rect) Cy() int {
	return r.Top() + r.SemiHeight()
}

func (r *Rect) Center() Point {
	return Point{r.Cx(), r.Cy()}
}

func (r *Rect) CenterLeft() Point {
	return Point{r.Cx() - r.SemiWidth(), r.Cy()}
}

func (r *Rect) CenterTop() Point {
	return Point{r.Cx(), r.Cy() - r.SemiHeight()}
}

func (r *Rect) CenterRight() Point {
	return Point{r.Cx() + r.SemiWidth(), r.Cy()}
}

func (r *Rect) CenterBottom() Point {
	return Point{r.Cx(), r.Cy() + r.SemiHeight()}
}

func (r *Rect) Intersects(rect *Rect) bool {
	xi := r.Left() <= rect.Right() &&
		r.Right() >= rect.Left()

	yi := r.Top() <= rect.Bottom() &&
		r.Bottom() >= rect.Top()

	return xi && yi
}

func (r *Rect) InBounds(rect *Rect) bool {
	return r.left <= rect.Left() &&
		r.top <= rect.Top() &&
		r.right >= rect.Right() &&
		r.bottom >= rect.Bottom()
}

func (r *Rect) setPosition(x int, y int) {
	if x < 0 || y < 0 {
		panic("X and Y must be non-negative integer numbers")
	}
	width := r.Width()
	height := r.Height()
	r.left = x
	r.top = y
	r.right = x + width
	r.bottom = y + height
}

func (r *Rect) MoveLeft(length int) {
	if length < 0 {
		panic("Length must be a non-negative integer number")
	}
	left := r.left - length
	right := r.right - length

	if left < 0 {
		panic("Invalid movement")
	}
	r.left = left
	r.right = right
}

func (r *Rect) MoveTop(length int) {
	if length < 0 {
		panic("Length must be a non-negative integer number")
	}
	top := r.top - length
	bottom := r.bottom - length

	if top < 0 {
		panic("Invalid movement")
	}
	r.top = top
	r.bottom = bottom
}

func (r *Rect) MoveRight(length int) {
	if length < 0 {
		panic("Length must be a non-negative integer number")
	}
	left := r.left + length
	right := r.right + length
	r.left = left
	r.right = right
}

func (r *Rect) MoveBottom(length int) {
	if length < 0 {
		panic("Length must be a non-negative integer number")
	}
	top := r.top + length
	bottom := r.bottom + length
	r.top = top
	r.bottom = bottom
}

func (r *Rect) Clone() Rect {
	return NewRect(
		r.Left(),
		r.Top(),
		r.Right(),
		r.Bottom(),
	)
}

func NewRect(left int, top int, right int, bottom int) Rect {
	if left < 0 || top < 0 || right < 0 || bottom < 0 {
		panic("Left, top, right and bottom must be non-negative integer numbers")
	}
	if !(left < right && top < bottom) {
		panic("Left must be less than right and top must be less than bottom")
	}
	return Rect{left, top, right, bottom}
}

type RectJSON struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

func (r *RectJSON) ToRect() *Rect {
	rect := NewRect(
		r.Left,
		r.Top,
		r.Right,
		r.Bottom,
	)
	return &rect
}

func NewRectJSON(r *Rect) *RectJSON {
	return &RectJSON{
		r.left,
		r.top,
		r.right,
		r.bottom,
	}
}

func Distance(p1 Point, p2 Point) int {
	dx := p1.X() - p2.X()
	dy := p1.Y() - p2.Y()
	return int(math.Sqrt(float64(dx*dx + dy*dy)))
}
