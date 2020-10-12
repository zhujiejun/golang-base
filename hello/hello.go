package main

import (
	"example.com/greetings"
	"fmt"
	"os"
)

func main() {
	message := greetings.Hello("Golang")
	if len(os.Args) > 1 {
		message = greetings.Hello(os.Args[1])
	}
	fmt.Println(message)
}
