/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	"github.com/gorilla/websocket"
	"log"
)

var match = NewRandomMatch()

type Hub struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan *ResponseData
	quit       chan struct{}
}

func (h *Hub) Start() {
	var broadcast = func(message *ResponseData) {
		for _, client := range h.clients {
			client.ch <- message
		}
	}

	var register = func(client *Client) {
		h.push(client)
		client.InitGame(match)
		go h.listen(client)
	}

	var unregister = func(client *Client) {
		h.delete(client)
	}

	for {
		select {
		case client := <-h.register:
			register(client)
		case client := <-h.unregister:
			unregister(client)
		case message := <-h.broadcast:
			broadcast(message)
		case <-h.quit:
			log.Println("Hub QUIT")
			return
		}
	}
}

func (h *Hub) Register(c *Client) {
	log.Printf("Client %s connected.\n", c.id)
	h.register <- c
}

func (h *Hub) Unregister(c *Client) {
	log.Printf("Client %s disconnected.\n", c.id)
	h.unregister <- c
}

func (h *Hub) push(client *Client) {
	h.clients[client.id] = client
}

func (h *Hub) delete(client *Client) {
	delete(h.clients, client.id)
}

func (h *Hub) listen(client *Client) {
	conn := client.conn

	for {
		_, p, err := conn.ReadMessage()

		if err != nil {
			if websocket.IsCloseError(err) {
				log.Println("Client disconnected", client.id)
			}
			client.Close()
			h.Unregister(client)
			return
		}
		message := string(p)
		id := client.id

		log.Println("Message " + message + " sent from client " + id)

		h.broadcast <- &ResponseData{
			Type: 0,
			Body: message,
		}
	}
}

func NewHub(ch chan *ResponseData, quit chan struct{}) *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  ch,
		quit:       quit,
	}
}
