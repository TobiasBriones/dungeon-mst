/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	id   string
	conn *websocket.Conn
	ch   chan ResponseData
	quit chan struct{}
}

func (c *Client) Handle() {
	for {
		select {
		case <-c.quit:
			if err := c.conn.Close(); err != nil {
				log.Printf("Failed to close %s client connection: %v\n", c.id, err)
			}
			return
		case v := <-c.ch:
			data := map[string]interface{}{
				"dataType": v.Type,
				"body":     v.Body,
			}
			if err := c.conn.WriteJSON(data); err != nil {
				log.Println("WS write error:", err)
				return
			}
		}
	}
}

func (c *Client) Close() {
	close(c.quit)
}

func NewClient(conn *websocket.Conn, id string) *Client {
	return &Client{
		id:   id,
		conn: conn,
		ch:   make(chan ResponseData),
		quit: make(chan struct{}),
	}
}
