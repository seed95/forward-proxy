package main

import (
	"log"

	"github.com/seed95/forward-proxy/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("cannot run the app, why? %v", err)
	}
}
