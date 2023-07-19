package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	var name string
	fmt.Scanf("%s", &name)
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
