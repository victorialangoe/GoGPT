package main

import (
	"GoGPT/chatbot"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	b := chatbot.Bot{Name: "GoGPT"}
    reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to GoGPT! Please enter your name: ")

    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)
	b.UserName = name

    fmt.Printf("Hello, %s!\n", name)

	for {
		fmt.Println("What would you like me to do?")
		fmt.Println("I can 1) introduce myself, 2) tell you the time, or 3) tell a joke or 0 to exit")
		
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			break
		}
		
		input = strings.TrimSpace(input)
		if input == "0" {
			fmt.Println("Exiting GoGPT!")
			break
		}
		b.ProcessMsg(input)
	}
}