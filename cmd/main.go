package main

import (
	"fmt"
	configs "go-rest-api/configs"
	auth "go-rest-api/internal/auth"
	hello "go-rest-api/internal/hello"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	config := configs.GenerateConfig()
	auth.NewAuthHandler(mux, auth.AuthHandlerDeps{
		Config: config,
	})
	hello.NewHelloHandler(mux)

	server := http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
