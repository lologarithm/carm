package main

import (
	"encoding/json"
	"fmt"
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

var pingTime = time.Second
var numPerps = 0

func handleSession(ws *websocket.Conn) {
	session := &Session{
		socketConn:  ws,
		jsonDecoder: json.NewDecoder(ws),
		jsonEncoder: json.NewEncoder(ws),
	}

	wrapper := &NetworkWrapper{}
	log.Printf("Starting listener for: %s", ws.RemoteAddr())
	pinger := time.AfterFunc(pingTime, func() {
		log.Printf("No ping after %d seconds!", pingTime/time.Second)
		RunLock()
	})
	for {
		err := session.jsonDecoder.Decode(wrapper)
		if err != nil {
			log.Printf("Failed to read from socket, closing down: %s", err)
			break
		}
		if wrapper.Command == "lock" {
			log.Printf("LOCKING IT DOWN")
			RunLock()
		} else if wrapper.Command == "disarm" {
			log.Printf("Disarming deadman switch!")
			pinger.Stop()
		} else if wrapper.Command == "ping" {
			pinger.Reset(pingTime)
		}
	}
	session.socketConn.WriteClose(1)
}

func RunLock() {
	lockCmd := exec.Command("gnome-screensaver-command", "-l")
	speakCmd := exec.Command("spd-say", "-r", "-50", "ALERT ALERT ALERT. Colby step away from the computer. ALERT ALERT ALERT ALERT.")
	pictureCmd := exec.Command("fswebcam", "-r", "1280x720", "--jpeg", "90", fmt.Sprintf("perps/perp%d.jpg", numPerps))
	numPerps++
	go func() {
		speakCmd.Run()
	}()
	go func() {
		pictureCmd.Run()
	}()
	go func() {
		time.Sleep(time.Second)
		lockCmd.Run() // this locks the computer!
	}()
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
