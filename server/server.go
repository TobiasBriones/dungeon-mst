/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

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
		name := waitForId(conn)

		if len(name) == 0 {
			return
		}

		client := NewClient(conn, name)

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

func waitForId(conn *websocket.Conn) string {
	_, p, err := conn.ReadMessage()

	if err != nil {
		return ""
	}
	return string(p)
}
