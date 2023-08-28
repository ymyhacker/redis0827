package main

import (
	"fmt"
	"bufio"
	"strings"
	"github.com/ymyhacker/redis0827/tree/YmY-branch/commands"
	"os"
)

func main() {
	fmt.Println("Welcome to Simplified Redis!")

	// Initialize the database

	dm := commands.NewDatabase()
	for {
		fmt.Print("> ")
		fmt.Scan()
        // 从stdin中取内容直到遇到换行符，停止
        input, err := bufio.NewReader(os.Stdin).ReadString('\n') 
        if err != nil {
            panic(err)
        }
        parts := strings.Split(strings.TrimSpace(input)," ")

		if len(parts) == 0 {
			continue
		}
		command := parts[0]
		if  strings.ToLower(command) == "exit" {
			break
		}
		args := parts[1:]
		response := commands.ExecuteCommand(dm, command, args)
		fmt.Println(response)
	}
}
