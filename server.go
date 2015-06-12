package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/ws", websocket.Handler(handleSession))

	http.ListenAndServe(":9876", nil)
}

func handleSession(ws *websocket.Conn) {
	session := &Session{
		socketConn:  ws,
		jsonDecoder: json.NewDecoder(ws),
		jsonEncoder: json.NewEncoder(ws),
	}

	wrapper := &NetworkWrapper{}
	for {
		err := session.jsonDecoder.Decode(wrapper)
		if err != nil {
			log.Printf("Failed to read from socket, closing down: %s", err)
			break
		}

		// TODO: Main Logic Here.
		if wrapper.Command == "lock" {

		} else if wrapper.Command == "unlock" {

		}
	}
	session.socketConn.WriteClose(1)
}

type NetworkWrapper struct {
	Command string `json:"commmand"`
}

// Session represents a connection between server and client.
type Session struct {
	socketConn  *websocket.Conn
	jsonDecoder *json.Decoder
	jsonEncoder *json.Encoder
}

// Generically write any object to JSON over the socket.
func (se *Session) WriteObject(i interface{}) error {
	err := se.jsonEncoder.Encode(i)
	if err != nil {
		log.Printf("Failed to send message on socket: %s", err)
		return err
	}
	return err
}
