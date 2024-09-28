package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: Routes(),
	}
	server.ListenAndServe()
}
