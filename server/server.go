/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type hub struct {
	clients   map[*client]bool
	broadcast chan []byte
}

type client struct {
	conn *websocket.Conn
	send chan []byte
	hub  *hub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	setRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setRoutes() {
	http.HandleFunc("/", wsHandler)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected...")

	err = ws.WriteMessage(1, []byte("This is the server response"))
	if err != nil {
		log.Println(err)
	}
	listenConn(ws)
}

func listenConn(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		err = conn.WriteMessage(messageType, p)

		if err != nil {
			log.Println(err)
			return
		}

	}
}
