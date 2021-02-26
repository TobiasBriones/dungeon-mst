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
