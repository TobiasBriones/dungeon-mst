/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

const (
	MoveDirLeft   = 0
	MoveDirTop    = 1
	MoveDirRight  = 2
	MoveDirBottom = 3
)

type Movement struct {
	direction int
	length    int
}

func CheckMovement(movement Movement, rect *Rect, host Rect) bool {
	if movement.direction == MoveDirLeft {
		return rect.Left-movement.length > host.Left
	} else if movement.direction == MoveDirTop {
		return rect.Top-movement.length > host.Top
	} else if movement.direction == MoveDirRight {
		return rect.Right+movement.length < host.Right
	} else if movement.direction == MoveDirBottom {
		return rect.Bottom+movement.length < host.Bottom
	}
	return false
}

func WillCollide(movement Movement, rect *Rect, objRect *Rect) bool {
	dst := Move(objRect, movement)
	return rect.Intersects(dst)
}

func Move(rect *Rect, movement Movement) *Rect {
	length := movement.length
	dst := &Rect{
		Left:   rect.Left,
		Top:    rect.Top,
		Right:  rect.Right,
		Bottom: rect.Bottom,
	}

	if movement.direction == MoveDirLeft {
		dst.Left -= length
		dst.Right -= length
	} else if movement.direction == MoveDirTop {
		dst.Top -= length
		dst.Bottom -= length
	} else if movement.direction == MoveDirRight {
		dst.Right += length
		dst.Left += length
	} else if movement.direction == MoveDirBottom {
		dst.Bottom += length
		dst.Top += length
	}
	return dst
}
