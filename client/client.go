package main

import "net"
import "log"



func main(){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil{
		log.Println("Ошибка подключения к серверу: ", err)
		return
	}
	defer conn.Close()

	var ans []byte
	_, err = conn.Read(ans)
	if err != nil{
		log.Println("Не удалось прочитать ответ от сервера: ", err)
		return
	}
	if string(ans) != "OK\n"{
		log.Println("От сервера получен неверный ответ: ", string(ans))
	}
}
