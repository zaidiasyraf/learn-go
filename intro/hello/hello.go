package main

import (
	"fmt"
	"log"

	"rsc.io/quote/v4"
	"zaidi.com/greeting"
)

func main() {
	// Hello world
	fmt.Println("Hello World")

	// quote from external resource
	fmt.Println(quote.Go())

	// greeting from other module
	message, err := greeting.Hello("Ken")
	if err != nil {
		// do nothing
	}
	fmt.Println(message)

	log.SetPrefix("Greeting: ")
	log.SetFlags(0)

	// Simulate error
	//message, err = greeting.Hello("")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(message)

	// multiple name
	names := []string{
		"Kim",
		"Jack",
		"Sam",
	}

	messages, err := greeting.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
