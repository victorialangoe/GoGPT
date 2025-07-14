package chatbot

import (
	"GoGPT/handlers"
	"fmt"
)

type Bot struct{
	Name string
	UserName string
	InputHandler *handlers.InputHandler
}

func (b Bot) ProcessMsg (input string) {
	// Check for number-based commands first
	switch input {
		case "1":
			fmt.Println(b.getGreetingResponse())
			return
		case "2":
			fmt.Println(b.getTTimeResponse())
			return
		case "3":
			fmt.Println(b.getJokeResponse())
			return
	}
	
	// Check for natural language commands
	greetingKeywords := []string{"hello", "hi", "introduce", "greeting", "who", "yourself"}
	timeKeywords := []string{"time", "clock", "hour", "minute", "now"}
	jokeKeywords := []string{"joke", "funny", "laugh", "humor", "tell"}
	
	if b.InputHandler.MatchKeywords(input, greetingKeywords) {
		fmt.Println(b.getGreetingResponse())
	} else if b.InputHandler.MatchKeywords(input, timeKeywords) {
		fmt.Println(b.getTTimeResponse())
	} else if b.InputHandler.MatchKeywords(input, jokeKeywords) {
		fmt.Println(b.getJokeResponse())
	} else {
		fmt.Println("I don't understand that. You can:")
		fmt.Println("- Type 1 or say 'hello' or 'introduce yourself'")
		fmt.Println("- Type 2 or ask for 'time'")
		fmt.Println("- Type 3 or ask for a 'joke'")
		fmt.Println("- Type 0 to exit")
	}
}