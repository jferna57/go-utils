package greetings

import "fmt"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	message := fmt.Sprintf("Hello, %s!", name)
	return message
}
