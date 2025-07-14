package chatbot

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
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
	bs, err := os.ReadFile("./data/jokes.txt")
	if err != nil {
		fmt.Println("Error:",err)
		os.Exit(1)
	}

	jokes := strings.Split(string(bs),",")

	n := rand.Intn(len(jokes)-1)

	return jokes[n]
}