// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package main

import (
	"dungeon-mst/dungeon"
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
	MatchJSON     *dungeon.MatchJSON
	RemainingTime time.Duration
	Players       []*PlayerJoin
}

type JoinAccepted struct {
	Id int
}

type PlayerJoin struct {
	Id        int
	Name      string
	PointJSON dungeon.PointJSON
	Score     int
}

type Update struct {
	Id int
	//Move int // use point for now
	PointJSON    dungeon.PointJSON
	DiamondIndex int
}
