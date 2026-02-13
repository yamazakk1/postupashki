package main

import (
	"log"
	"net"
)


func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("Ошибка при запуске сервера: ", err)
		return
	}
	defer listener.Close()

	log.Println("Сервер запущен на порту 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Ошибка принятия соедиинения: ", err)
			continue
		}

		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	_, err := conn.Write([]byte("OK\n"))
	if err != nil {
		log.Println("Ошибка отправки ответа: ", err)
		return
	}
	log.Println("Ответ отправлен, соединение закрывается йоу")
}
