/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	addr = "localhost:8080"
)

var id = 0

func main() {
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()

	dataCh := make(chan *ResponseData)
	quitCh := make(chan struct{})
	hub := NewHub(dataCh, quitCh)

	defer close(quitCh)
	go hub.Start()

	sendFakeUpdate(hub)

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
		name := "remote " + strconv.Itoa(id)
		client := NewClient(conn, name)
		id++

		go client.Handle()

		hub.Register(client)
		hub.broadcast <- &ResponseData{
			Type: DataTypeServerMessage,
			Body: "New client connected " + client.id + "...",
		}
	}
}

func getUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
}

type Update struct {
	M int
}

func sendFakeUpdate(hub *Hub) {
	ticker := time.NewTicker(100 * time.Millisecond)

	go func() {
		for range ticker.C {
			u := &Update{
				M: rand.Intn(4),
			}
			enc, _ := json.Marshal(u)

			hub.broadcast <- &ResponseData{
				Type: DataTypeUpdate,
				Body: string(enc),
			}
		}
	}()
}
