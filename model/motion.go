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
