package main

import (
	"GoGPT/chatbot"
	"GoGPT/handlers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputHandler := handlers.NewInputHandler()
	b := chatbot.Bot{
		Name: "GoGPT",
		InputHandler: inputHandler,
	}
    reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to GoGPT! Please enter your name: ")

    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)
	b.UserName = name

    fmt.Printf("Hello, %s!\n", name)

	for {
		fmt.Println("What would you like me to do?")
		fmt.Println("You can type numbers (1, 2, 3) or use natural language:")
		fmt.Println("- 1 or 'hello' / 'introduce yourself'")
		fmt.Println("- 2 or 'what time is it?'")
		fmt.Println("- 3 or 'tell me a joke'")
		fmt.Println("- 0 to exit")
		
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