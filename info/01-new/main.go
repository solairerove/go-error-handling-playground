package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	v, err := f()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(v)
}

func f() (int, error) {
	return 0, errors.New("shit happens")
}
