package main

import (
	"GoGPT/chatbot"
	"bufio"
	"fmt"
	"os"
)

func main() {
	b := chatbot.Bot{Name: "GoGPT"}
    sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to GoGPT! Please enter your name: ")

    sc.Scan()

    name := sc.Text()
	b.UserName = name

    fmt.Printf("Hello, %s!\n", name)

	for {
		fmt.Println("What would you like me to do?")
		fmt.Println("I can 1) introduce myself, 2) tell you the time, or 3) tell a joke or 0 to exit")
		if !sc.Scan() {
			break
		}
		input := sc.Text()
		if input == "0" {
			fmt.Println("Exiting GoGPT!")
			os.Exit(1)
		}
		b.ProcessMsg(input)
	}
}