package greetings

import "fmt"

func Welcome() {
	fmt.Println("Welcome!")
}

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
