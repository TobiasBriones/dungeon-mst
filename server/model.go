/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	"server/model"
	"time"
)

const (
	DataTypeGameInitialization = 0
	DataTypeUpdate             = 1
	DataTypeServerMessage      = 2
	DataTypeJoinAccepted       = 3
)

type ResponseData struct {
	Type int
	Body string
}

type MatchInit struct {
	MatchJSON            *model.MatchJSON
	RemainingTimeSeconds time.Duration
}

type JoinAccepted struct {
	Id int
}

type Update struct {
	Id   int
	Move int
}
