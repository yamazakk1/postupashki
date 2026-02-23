package main

import (
	"log"
	"net"
	"time"
)

const (
	protocol       string = "tcp"
	address        string = "localhost:8080"
	expectedAnswer string = "OK\n"
)

func main() {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		log.Println("Connection error: ", err)
		return
	}
	defer conn.Close()

	var ans []byte
	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	_, err = conn.Read(ans)

	if err != nil {
		log.Println("Unable to read server-answer: ", err)
		return
	}
	if string(ans) != expectedAnswer {
		log.Println("Wrong server-answer received: ", string(ans))
	}
}
