package chatbot

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func (b Bot) getGreetingResponse() string {
	return fmt.Sprintf("Hello %s, I am %s! \n", b.UserName, b.Name)
}

func (Bot) getTTimeResponse() string {
	t := time.Now()
	layout := "15:04:05"
	return fmt.Sprintf("The time right now is %s and the day is %sth of %s", t.Format(layout), fmt.Sprintf("%d", t.Day()), t.Month())
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