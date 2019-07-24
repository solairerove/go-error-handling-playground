package main

import (
	"fmt"

	"github.com/solairerove/go-error-handling-playground/scan"
)

func main() {
	fmt.Println("Error handling playground")

	f, e := fmt.Println("Error?")
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(f)

	scan.ExampleScan()
}
