package chatbot

import (
	"fmt"
	"time"
)

func (b Bot) getGreetingResponse() string {
	return fmt.Sprintf("Hello %s, I am %s!", b.UserName, b.Name)
}

func (Bot) getTTimeResponse() string {
	time := time.Now()
	return fmt.Sprintf("The time right now is %s",time)
}

func (Bot) getJokeResponse() string {
	
}