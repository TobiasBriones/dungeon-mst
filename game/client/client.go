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

type MatchInitJSON struct {
	MatchJSON            *model.MatchJSON
	RemainingTimeSeconds time.Duration
}

type MatchInit struct {
	Match         *model.Match
	RemainingTime time.Duration
}

type Update struct {
	Id   int
	Move int
}

func Run(name string, accepted chan *JoinAccepted, matchCh chan *MatchInit, ch chan *Update, sendUpdate chan *Update) {
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
	readMessages(done, conn, matchCh, ch)
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
	log.Printf("Response: %+v\n", data)

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

func readMessages(done chan struct{}, conn *websocket.Conn, h chan *MatchInit, ch chan *Update) {
	init := func(body string) {
		matchInitJSON := &MatchInitJSON{}

		if err := json.Unmarshal([]byte(body), matchInitJSON); err != nil {
			log.Println("Match read error:", err)
			return
		}
		//fmt.Printf("%+v\n", matchJSON)

		matchJSON := matchInitJSON.MatchJSON
		match := matchJSON.ToMatch()
		matchInit := &MatchInit{
			Match:         match,
			RemainingTime: matchInitJSON.RemainingTimeSeconds,
		}
		//fmt.Printf("%+v\n", match)
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

	readResponse := func(data *ResponseData) {
		switch data.Type {
		case 0:
			init(data.Body)
		case 1:
			update(data.Body)
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
