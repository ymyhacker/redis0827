package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/ymyhacker/redis0827/tree/YmY-branch/commands"

)

func main() {
	fmt.Println("Welcome to Simplified Redis!")

	// Initialize the database
	// dm := commands.NewDatabaseManager()
	dm := commands.NewDatabaseManager()
	for {
		fmt.Print("> ")
		var input string
		fmt.Scanln(&input)

		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		args := parts[1:]
		if command == "Exits"
			break
		response := dm.commands.ExecuteCommand(db, command, args)
		fmt.Println(response)
	}
}
