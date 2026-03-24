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

	// err := os.MkdirAll("out", 0755)
	// if err != nil {
	// 	panic(err)
	// }

	// file, err := os.Create("out/test.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// file.Write([]byte("my test file"))

	// defer file.Close()
}
