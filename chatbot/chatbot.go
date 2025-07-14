package chatbot

import "fmt"

type Bot struct{
	Name string
	UserName string
}

func (b Bot) ProcessMsg (input string) {
	switch input {
		case "1":
			fmt.Println(b.getGreetingResponse())
		case "2":
			fmt.Println(b.getTTimeResponse())
		case "3":
			fmt.Println(b.getJokeResponse())
		default:
			fmt.Println("I don't understand that option. Please choose 1, 2, 3, or 0 to exit.")
	}
}