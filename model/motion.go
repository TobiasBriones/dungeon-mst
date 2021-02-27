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
		return rect.Left()-movement.length > host.Left()
	} else if movement.direction == MoveDirTop {
		return rect.Top()-movement.length > host.Top()
	} else if movement.direction == MoveDirRight {
		return rect.Right()+movement.length < host.Right()
	} else if movement.direction == MoveDirBottom {
		return rect.Bottom()+movement.length < host.Bottom()
	}
	return false
}

func WillCollide(movement Movement, rect *Rect, objRect *Rect) bool {
	dst := Move(objRect, movement)
	return rect.Intersects(dst)
}

func Move(rect *Rect, movement Movement) *Rect {
	length := movement.length
	dst := rect.Clone()

	if movement.direction == MoveDirLeft {
		dst.MoveLeft(length)
	} else if movement.direction == MoveDirTop {
		dst.MoveTop(length)
	} else if movement.direction == MoveDirRight {
		dst.MoveRight(length)
	} else if movement.direction == MoveDirBottom {
		dst.MoveBottom(length)
	}
	return &dst
}
