package main

import (
	"log"
)

func main() {
	err := RunTerm()
	if err != nil {
		log.Fatal(err)
	}
}
