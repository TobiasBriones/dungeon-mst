/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

import "math"

type Point struct {
	X int
	Y int
}

type PointPair struct {
	PointA Point
	PointB Point
}

type Dimension struct {
	Width  int
	Height int
}

func (s *Dimension) SemiWidth() int {
	return s.Width / 2
}

func (s *Dimension) SemiHeight() int {
	return s.Height / 2
}

type Rect struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

func (r *Rect) Width() int {
	return r.Right - r.Left
}

func (r *Rect) SemiWidth() int {
	return r.Width() / 2
}

func (r *Rect) Height() int {
	return r.Bottom - r.Top
}

func (r *Rect) SemiHeight() int {
	return r.Height() / 2
}

func (r *Rect) Cx() int {
	return r.Left + r.SemiWidth()
}

func (r *Rect) Cy() int {
	return r.Top + r.SemiHeight()
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
	xi := r.Left <= rect.Right &&
		r.Right >= rect.Left

	yi := r.Top <= rect.Bottom &&
		r.Bottom >= rect.Top

	return xi && yi
}

func (r *Rect) InBounds(rect *Rect) bool {
	if !r.Intersects(rect) {
		return false
	}
	return r.Right >= rect.Right && r.Bottom >= rect.Bottom
}

func Distance(p1 Point, p2 Point) int {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return int(math.Sqrt(float64(dx*dx + dy*dy)))
}
