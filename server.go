package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"time"

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
	log.Printf("Starting listener for: %s", ws.RemoteAddr())
	for {
		err := session.jsonDecoder.Decode(wrapper)
		if err != nil {
			log.Printf("Failed to read from socket, closing down: %s", err)
			break
		}
		log.Printf("Message: '%s'", wrapper.Command)
		if wrapper.Command == "lock" {
			lockCmd := exec.Command("gnome-screensaver-command", "-l")
			speakCmd := exec.Command("spd-say", "-r", "-50", "ALERT ALERT ALERT. Colby step away from the computer. ALERT ALERT ALERT ALERT.")
			go func() {
				speakCmd.Run()
			}()
			go func() {
				time.Sleep(time.Second * 3)
				lockCmd.Run() // this locks the computer!
			}()
		} else if wrapper.Command == "unlock" {

		}
	}
	session.socketConn.WriteClose(1)
}

type NetworkWrapper struct {
	Command string `json:"command"`
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
