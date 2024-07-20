package main

import (
	"log"

	"github.com/SleepWlaker/GoTerminalIndicator/server"
)

func main() {
	s := server.NewServer()
	log.Fatal(s.Run())
}
