package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("some_text.txt")
	if err != nil {
		fmt.Println("shit happens")
	}
}
