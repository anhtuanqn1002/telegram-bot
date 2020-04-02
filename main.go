package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Coin struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func getBTCprice(symbol string) string {
	url := "https://api.binance.com/api/v3/ticker/price?symbol=" + strings.ToUpper(symbol) + "BTC"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var coin Coin
	json.Unmarshal(body, &coin)
	return coin.Price
}

func main() {
	fmt.Printf("Goooo...")
	bot, err := tgbotapi.NewBotAPI("1193668041:AAHWLyqlEAql50aJZPVKq8TCpqB3tLqUL0k")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		price := getBTCprice("ltc")

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, price)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
