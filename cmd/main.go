package main

import (
	"kagami/server/v1/udp"
	"log"
	"os"
)

func main() {
	args := os.Args

	switch {
	case len(args) == 4 && args[1] == "server" && args[2] == "udp":
		udp.Loop(args[3])
	default:
		log.Fatal("Usage:", args[0], "server", "udp", "0.0.0.0:9999")
	}
}
