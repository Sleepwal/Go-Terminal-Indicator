package main

import (
	"log"

	"github.com/SleepWlaker/GoTerminalIndicator/internal"
)

func main() {
	s := internal.NewServer()
	log.Fatal(s.Run())
}
