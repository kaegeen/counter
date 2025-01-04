package main

import (
	"fmt"
	"time"
)

func countdown(seconds int) {
	for seconds > 0 {
		fmt.Printf("Time left: %d seconds\n", seconds)
		time.Sleep(1 * time.Second)
		seconds--
	}
	fmt.Println("Time's up!")
}

func main() {
	fmt.Println("Countdown Timer")
	fmt.Println("=================")
	fmt.Println("Type 'start' to begin a countdown.")
	fmt.Println("Type 'exit' to quit.\n")

	for {
		fmt.Print("Command: ")
		var command string
		fmt.Scanln(&command)

		if command == "exit" {
			fmt.Println("Goodbye!")
			break
		} else if command == "start" {
			var seconds int
			fmt.Print("Enter countdown time in seconds: ")
			_, err := fmt.Scanln(&seconds)
			if err != nil || seconds <= 0 {
				fmt.Println("Error: Please enter a valid positive number.\n")
				continue
			}
			fmt.Println("Starting countdown...")
			countdown(seconds)
			fmt.Println()
		} else {
			fmt.Println("Invalid command. Use 'start' or 'exit'.\n")
		}
	}
}
