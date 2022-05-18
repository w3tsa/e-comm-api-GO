package main

import (
	"fmt"
	"github.com/w3tsa/e-comm-api-GO/greetings"
)

func main() {
	message := greetings.Hello("Rahim")
	fmt.Println(message)
}
