package main

import (
	"log"

	"github.com/assaidy/url-shortener/server"
)

func main() {
	server := server.NewFiberServer()
	server.RegisterRoutes()
	log.Fatal(server.Listen(":8080"))
}
