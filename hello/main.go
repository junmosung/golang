package main

import (
	"awesomeProject/greetings"
	"fmt"
	"log"
)

func main() {
	//message := greetings.Hello("junmo")
	//fmt.Println(message)

	//result, err := greetings.HelloAddError("")
	//fmt.Println(result)

	//////////////////////////////////////////////////

	//// Set properties of the predefined Logger, including
	//// the log entry prefix and a flag to disable printing
	//// the time, source file, and line number.
	//log.SetPrefix("greetings: ")
	//log.SetFlags(0)
	//
	//// Request a greeting message.
	//message, err := greetings.HelloAddError("")
	//// If an error was returned, print it to the console and
	//// exit the program.
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// If no error was returned, print the returned message
	//// to the console.
	//fmt.Println(message)

	///////////////////////////////////////////////////////

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("junmo")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
