/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package client

import (
	"bufio"
	"encoding/json"
	"flag"
	"game/model"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

type ResponseData struct {
	Type int
	Body string
}

type JoinAccepted struct {
	Id int
}

type PlayerJoin struct {
	Id        int
	Name      string
	PointJSON model.PointJSON
}

type MatchInit struct {
	MatchJSON     *model.MatchJSON
	Match         *model.Match
	RemainingTime time.Duration
	Players       []*PlayerJoin
}

type Update struct {
	Id int
	//Move int // use point for now
	PointJSON model.PointJSON
}

func Run(
	name string,
	accepted chan *JoinAccepted,
	matchCh chan *MatchInit,
	ch chan *Update,
	sendUpdate chan *Update,
	joinCh chan *PlayerJoin,
) {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: ""}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	done := make(chan struct{})

	waitAccepted(name, accepted, conn)
	readMessages(done, conn, matchCh, ch, joinCh)
	writeMessages(done, conn, sendUpdate)

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func waitAccepted(name string, acceptedCh chan *JoinAccepted, conn *websocket.Conn) {
	if err := conn.WriteMessage(websocket.TextMessage, []byte(name)); err != nil {
		log.Println("Name write error:", err)
		return
	}

	_, p, err := conn.ReadMessage()

	if err != nil {
		log.Println("Read error:", err)
		return
	}
	data := &ResponseData{}

	if err := json.Unmarshal(p, data); err != nil {
		log.Println("Read ResponseData error:", err)
		return
	}

	if data.Type != 3 {
		log.Println("Failed to connect, invalid server accepted response")
		return
	}
	accepted := &JoinAccepted{}

	if err := json.Unmarshal([]byte(data.Body), accepted); err != nil {
		log.Println("JoinAccepted read error:", err)
		return
	}
	acceptedCh <- accepted
}

func readMessages(
	done chan struct{},
	conn *websocket.Conn,
	h chan *MatchInit,
	ch chan *Update,
	joinCh chan *PlayerJoin,
) {
	init := func(body string) {
		matchInit := &MatchInit{}

		if err := json.Unmarshal([]byte(body), matchInit); err != nil {
			log.Println("Match read error:", err)
			return
		}
		//fmt.Printf("%+v\n", match)
		matchInit.Match = matchInit.MatchJSON.ToMatch()
		h <- matchInit
	}

	update := func(body string) {
		update := &Update{}

		if err := json.Unmarshal([]byte(body), update); err != nil {
			log.Println("Update read error:", err)
			return
		}
		ch <- update
	}

	join := func(body string) {
		join := &PlayerJoin{}

		if err := json.Unmarshal([]byte(body), join); err != nil {
			log.Println("Join read error:", err)
			return
		}
		joinCh <- join
	}

	readResponse := func(data *ResponseData) {
		switch data.Type {
		case 0:
			init(data.Body)
		case 1:
			update(data.Body)
		case 4:
			join(data.Body)
		}
	}

	go func() {
		defer close(done)
		for {
			_, p, err := conn.ReadMessage()

			if err != nil {
				log.Println("Read error:", err)
				return
			}
			//log.Printf("recv: %s", p)
			data := &ResponseData{}

			if err := json.Unmarshal(p, data); err != nil {
				log.Println("Read ResponseData error:", err)
				return
			}
			//log.Printf("Response: %+v\n", data)
			readResponse(data)
		}
	}()
}

func writeMessages(done chan struct{}, conn *websocket.Conn, ch chan *Update) {
	go func() {
		for {
			select {
			case <-done:
				return
			case u := <-ch:
				sendUpdate(conn, u)
			}
		}
	}()
}

func sendUpdate(conn *websocket.Conn, update *Update) {
	enc, err := json.Marshal(update)

	if err != nil {
		log.Println("Update encoding error:", err)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, enc); err != nil {
		log.Println("Update write error:", err)
		return
	}
}
