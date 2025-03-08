package main

import (
	"fmt"
	"github.com/AgustinCardozo/labo-so/utils/greetings"
)

func main() {
	message := greetings.Hello("Burak")
	fmt.Println(message)
}
