package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello function returns a greeting for the named person
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty name")
	}
	//Returns a greeting with the name ebeed in a message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v, welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
