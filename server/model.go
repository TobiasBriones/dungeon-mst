/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

const (
	DataTypeGameInitialization = 0
	DataTypeUpdate             = 1
)

type ResponseData struct {
	Type int
	Body string
}
