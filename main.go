package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/seed95/forward-proxy/cmd"
	"log"
	//"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(fmt.Printf("cannot run the app, why? %v\n", aurora.Red(err)))
	}
}
