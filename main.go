package main

import (
	"log"
)

func main() {
	// Run the terminal menu
	err := RunTerm()
	if err != nil {
		log.Fatal(err)
	}
}
