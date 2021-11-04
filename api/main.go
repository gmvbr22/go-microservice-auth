//go:build !test

package main

import (
	"github.com/gmvbr/go-rest/api/app"
	"log"
)

func main() {
	server := app.Setup()
	log.Fatal(server.Listen(":3000"))
}
