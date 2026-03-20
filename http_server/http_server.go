package http_server

import (
	"errors"
	"fmt"
	"net/http"
)

func Start() error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("processed request /ping")
		
		w.Write([]byte("hello from Docker\n"))
	})

	err := http.ListenAndServe(":5050", nil)

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}