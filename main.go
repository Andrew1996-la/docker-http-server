package main

import (
	"docker-http-server/http_server"
	"fmt"
)

func main() {
	fmt.Println("Server start")

	err := http_server.Start()
	if err != nil {
		fmt.Println("Во время работы с сервером произошла ошибка:", err)
		return
	} else {
		fmt.Println("Сервер завершился успешно")
	}
}
