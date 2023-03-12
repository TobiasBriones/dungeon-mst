/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	"dungeon-mst/server/model"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Client struct {
	PointJSON model.PointJSON
	Score     int
	id        int
	name      string
	conn      *websocket.Conn
	ch        chan *ResponseData
	quit      chan struct{}
}

func (c *Client) InitGame(match *model.Match, time time.Duration, players []*PlayerJoin) {
	matchJSON := model.NewMatchJSON(match)
	matchInit := &MatchInit{
		MatchJSON:     matchJSON,
		RemainingTime: time,
		Players:       players,
	}
	enc, err := json.Marshal(matchInit)

	if err != nil {
		log.Println(err)
		return
	}

	data := &ResponseData{
		Type: DataTypeGameInitialization,
		Body: string(enc),
	}
	if err := c.conn.WriteJSON(data); err != nil {
		log.Println("WS write error:", err)
		return
	}
}

func (c *Client) SendId() {
	accepted := &JoinAccepted{Id: c.id}
	enc, err := json.Marshal(accepted)

	if err != nil {
		log.Println(err)
		return
	}
	data := &ResponseData{
		Type: DataTypeJoinAccepted,
		Body: string(enc),
	}
	if err := c.conn.WriteJSON(data); err != nil {
		log.Println("WS write error:", err)
		return
	}
}

func (c *Client) Handle() {
	for {
		select {
		case <-c.quit:
			if err := c.conn.Close(); err != nil {
				log.Printf("Failed to close %s client connection: %v\n", c.id, err)
			}
			return
		case data := <-c.ch:
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

func NewClient(conn *websocket.Conn, id int, name string) *Client {
	return &Client{
		id:   id,
		name: name,
		conn: conn,
		ch:   make(chan *ResponseData),
		quit: make(chan struct{}),
	}
}
