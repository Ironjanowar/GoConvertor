package main

import (
	"fmt"
	"github.com/rockneurotiko/go-tgbot"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func help(bot tgbot.TgBot, msg tgbot.Message, text string) *string {
	bot.Answer(msg).Text("hola").End()
	return nil
}

func main() {
	bot := tgbot.NewTgBot("").
		SimpleCommandFn("help", help)

	fmt.Println("Running..")
	bot.SimpleStart()
}
