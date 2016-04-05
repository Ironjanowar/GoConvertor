package main

import (
	"fmt"
	"github.com/rockneurotiko/go-tgbot"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func help(bot tgbot.TgBot, msg tgbot.Message, text string) *string {
	commands := map[string]string{
		"start": "Starts the bot!",
		"help":  "Show commands abailables",
		"todec": "Takes an hexadecimal or binary number and returns its decimal equivalent",
		"tohex": "Takes a decimal number and returns its hexadecimal equivalent",
		"tobin": "Takes a decimal number and returns its binary equivalent",
	}

	helpMessage := "Available commands:\n\n"

	for key, value := range commands {
		helpMessage = helpMessage + "- /" + key + " ::\n    " + value + "\n"
	}

	bot.Answer(msg).Text(helpMessage).End()
	return nil
}

func tobin(bot tgbot.TgBot, msg tgbot.Message, vals []string, kvals map[string]string) *string {
	if intNum, err := strconv.Atoi(vals[1]); err == nil {
		binNum := strconv.FormatInt(int64(intNum), 2)
		bot.Answer(msg).Text("0b" + binNum).ReplyToMessage(msg.ID).End()
		return nil
	}
	error := "That's not a decimal number!"

	bot.Answer(msg).Text(error).ReplyToMessage(msg.ID).End()
	return nil
}

func tohex(bot tgbot.TgBot, msg tgbot.Message, vals []string, kvals map[string]string) *string {
	if len(vals) == 1 {
		error := "Try:\n /tohex [number]\n\nExample:\n /tohex 10"
		bot.Answer(msg).Text(error).ReplyToMessage(msg.ID).End()
		return nil
	}

	if intNum, err := strconv.Atoi(vals[1]); err == nil {
		hexNum := strconv.FormatInt(int64(intNum), 16)
		bot.Answer(msg).Text("0x" + hexNum).ReplyToMessage(msg.ID).End()
		return nil
	}
	error := "That's not a decimal number"

	bot.Answer(msg).Text(error).ReplyToMessage(msg.ID).End()
	return nil
}

func todec(bot tgbot.TgBot, msg tgbot.Message, vals []string, kvals map[string]string) *string {
	if len(vals) == 1 {
		error := "Try:\n /todec [number]\n\nExample:\n /todec 0xa"
		bot.Answer(msg).Text(error).ReplyToMessage(msg.ID).End()
		return nil
	}

	if strings.HasPrefix(vals[1], "0x") {
		if num, err := strconv.ParseInt(vals[1], 0, 64); err == nil {
			bot.Answer(msg).Text(strconv.Itoa(int(num))).ReplyToMessage(msg.ID).End()
			return nil
		}
		error := "That's not an hexadecimal number.\nTry:\n  /todec 0xa"
		bot.Answer(msg).Text(error).ReplyToMessage(msg.ID).End()
		return nil
	}

	if strings.HasPrefix(vals[1], "0b") {
		binNum := strings.Split(vals[1], "b")[1]
		if num, err := strconv.ParseInt(binNum, 2, 64); err == nil {
			bot.Answer(msg).Text(strconv.Itoa(int(num))).ReplyToMessage(msg.ID).End()
			return nil
		}
		error := "That's not a binary number.\nTry:\n  /todec 0b1010"
		bot.Answer(msg).Text(error).ReplyToMessage(msg.ID).End()
		return nil
	}

	bot.Answer(msg).Text("That's neither a binary or hexadecimal number.\nTry:\n  /todec 0xa\n  /todec 0b1010").ReplyToMessage(msg.ID).End()
	return nil
}

func Stobin(bot tgbot.TgBot, msg tgbot.Message, text string) *string {
	error := "Try:\n /tobin [number]\n\nExample:\n /tobin 10"
	bot.Answer(msg).Text(error).ReplyToMessage(msg.ID).End()
	return nil
}

func Stohex(bot tgbot.TgBot, msg tgbot.Message, text string) *string {
	error := "Try:\n /tohex [number]\n\nExample:\n /tohex 10"
	bot.Answer(msg).Text(error).ReplyToMessage(msg.ID).End()
	return nil
}

func Stodec(bot tgbot.TgBot, msg tgbot.Message, text string) *string {
	error := "Try:\n /todec [number]\n\nExample:\n /todec 0xa"
	bot.Answer(msg).Text(error).ReplyToMessage(msg.ID).End()
	return nil
}

func main() {
	filename := "convertor.token"
	TOKEN, err := ioutil.ReadFile(filename)
	check(err)

	bot := tgbot.NewTgBot(string(TOKEN)).
		SimpleCommandFn("help", help).
		CommandFn("tobin (.+)", tobin).
		CommandFn("tohex (.+)", tohex).
		CommandFn("todec (.+)", todec).
		SimpleCommandFn("tobin", Stobin).
		SimpleCommandFn("tohex", Stohex).
		SimpleCommandFn("todec", Stodec)

	fmt.Println("Running..")
	bot.SimpleStart()
}
