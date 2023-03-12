/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	"dungeon-mst/game/model"
	"time"
)

const (
	DataTypeGameInitialization = 0
	DataTypeUpdate             = 1
	DataTypeServerMessage      = 2
	DataTypeJoinAccepted       = 3
	DataTypePlayerJoin         = 4
	DataTypePlayerLeft         = 5
)

type ResponseData struct {
	Type int
	Body string
}

type MatchInit struct {
	MatchJSON     *model.MatchJSON
	RemainingTime time.Duration
	Players       []*PlayerJoin
}

type JoinAccepted struct {
	Id int
}

type PlayerJoin struct {
	Id        int
	Name      string
	PointJSON model.PointJSON
	Score     int
}

type Update struct {
	Id int
	//Move int // use point for now
	PointJSON    model.PointJSON
	DiamondIndex int
}
