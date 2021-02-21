package app

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/diego3/hotel-booking-api/core"
)

// StartCommandLineApplication start a command line application
func StartCommandLineApplication() {
	command := Command("Enter the command:")

	if command == "ping" {
		core.ConnectionPing()
	}

	if command == "create" {
		roomName := Command("Enter the room name:")
		if len(roomName) == 0 {
			log.Fatalln("room name cant be empty")
		}
		core.RegisterRoom(roomName)
	}
}

// Command represents a simple command line with instructions before execution
func Command(instructions string) string {
	fmt.Print(instructions)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return ""
}
