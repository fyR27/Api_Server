package main

import (
	"D1-api/iternal/app/apiserver"
	"log"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
