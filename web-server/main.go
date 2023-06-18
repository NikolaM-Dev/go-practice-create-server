package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewServer(":3000")

	server.Handle(http.MethodGet, "/", server.AddMiddlewares(HandleRoot, Logging()))
	server.Handle(http.MethodPost, "/api", server.AddMiddlewares(HandleHome, CheckAuth(), Logging()))
	server.Handle(http.MethodPost, "/create", server.AddMiddlewares(HandlePost, Logging()))
	server.Handle(http.MethodPost, "/user", server.AddMiddlewares(HandlePostUser, Logging()))

	err := server.Listen()
	if err != nil {
		log.Fatal(err)
	}
}
