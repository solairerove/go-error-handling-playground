package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("some_text.txt")
	if err != nil {
		log.Println("shit happens")
	}
}
