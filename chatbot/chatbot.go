package chatbot

type Bot struct{
	Name string
	UserName string
}

func (b Bot) ProcessMsg (input string) {
	switch input {
		case "1":
			b.getGreetingResponse()
		case "2":
			b.getTTimeResponse()
		case "3":
			
	}
}