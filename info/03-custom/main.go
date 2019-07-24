package main

import (
	"fmt"
	"log"
)

type myError struct {
	msg string
	err error
}

func (e myError) Error() string {
	return fmt.Sprintf("shit happens %v %v", e.err, e.msg)
}

func main() {
	v, err := f()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(v)
}

func f() (int, error) {
	e := fmt.Errorf("shit happen again")
	return 0, myError{
		msg: "Jason Bourne",
		err: e,
	}
}
