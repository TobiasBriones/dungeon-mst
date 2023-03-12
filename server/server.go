// Copyright (c) 2021 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/dungeon-mst

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	addr = "localhost:8080"
)

var globalId = -1

func main() {
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()

	dataCh := make(chan *ResponseData)
	quitCh := make(chan struct{})
	hub := NewHub(dataCh, quitCh)

	defer close(quitCh)
	go hub.Start()

	r.GET("/", wsHandler(getUpgrader(), quitCh, hub))
	err := r.Run(addr)

	if err != nil {
		log.Fatal("Unable to run server: " + err.Error())
	}
}

func wsHandler(updgrader *websocket.Upgrader, quit chan struct{}, hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := updgrader.Upgrade(c.Writer, c.Request, nil)

		if err != nil {
			log.Println(err)
		}
		id, name := waitForConfirm(conn)

		if len(name) == 0 {
			return
		}

		client := NewClient(conn, id, name)

		client.SendId()
		go client.Handle()

		hub.Register(client)
	}
}

func getUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
}

func waitForConfirm(conn *websocket.Conn) (int, string) {
	_, p, err := conn.ReadMessage()
	globalId++

	if err != nil {
		log.Println(err)
		return globalId, ""
	}
	return globalId, string(p)
}
