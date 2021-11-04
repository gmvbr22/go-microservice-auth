//go:build !test

package main

import (
	"log"

	"github.com/gmvbr/go-microservice-auth/api/app"
)

func main() {
	server := app.Setup()
	log.Fatal(server.Listen(":3000"))
}
