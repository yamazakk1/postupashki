package main

import (
  "log"
  "net"
)


const(
  port string = ":8080"
  protocol string = "tcp"
  payload string = "OK\n"
)

func main() {
  listener, err := net.Listen(protocol, port)
  if err != nil {
    log.Println("Unable to start server: ", err)
    return
  }
  defer listener.Close()

  log.Println("Server started at localhost:8080")

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Println("Connection receiving error: ", err)
      continue
    }

    go handleConnection(conn)

  }
}

func handleConnection(conn net.Conn) {
  defer conn.Close()
  _, err := conn.Write([]byte(payload))
  if err != nil {
    log.Println("Unable to write answer: ", err)
    return
  }
  log.Println("Answer sent!")
}