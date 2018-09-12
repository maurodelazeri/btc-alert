package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gopkg.in/telegram-bot-api.v4"

	"github.com/isvaldo/go-coinmarketcap-client"
)

var client coinmarket.Interface

func init() {
	var err error
	client = coinmarket.New("https://api.coinmarketcap.com")
	if err != nil {
		panic(err)
	}
}

func main() {

	go func() {
		for {
			coinTicker, _ := client.GetTicker("bitcoin")
			s1 := fmt.Sprintf("Hourly Update (BTC)  \n Price: %s (USD) %s (BTC) \n Percent Change 1H %s \n Percent Change 24H %s \n Percent Change 7D %s \n", coinTicker.PriceUsd, coinTicker.PriceBtc, coinTicker.PercentChange1H, coinTicker.PercentChange24H, coinTicker.PercentChange7D)

			coinTicker, _ = client.GetTicker("nano")
			s2 := fmt.Sprintf("Hourly Update (NANO)  \n Price: %s (USD) %s (BTC) \n Percent Change 1H %s \n Percent Change 24H %s \n Percent Change 7D %s \n", coinTicker.PriceUsd, coinTicker.PriceBtc, coinTicker.PercentChange1H, coinTicker.PercentChange24H, coinTicker.PercentChange7D)

			coinTicker, _ = client.GetTicker("ethereum")
			s3 := fmt.Sprintf("Hourly Update (ETH)  \n Price: %s (USD) %s (BTC) \n Percent Change 1H %s \n Percent Change 24H %s \n Percent Change 7D %s \n", coinTicker.PriceUsd, coinTicker.PriceBtc, coinTicker.PercentChange1H, coinTicker.PercentChange24H, coinTicker.PercentChange7D)
			sendMessage(s1 + "\n" + s2 + "\n" + s3)

			time.Sleep(time.Hour)
		}
	}()

	pairs := []string{"nano", "bitcoin", "ethereum"}
	for {
		for _, pair := range pairs {
			coinTicker, _ := client.GetTicker(pair)
			priceChange, _ := strconv.ParseFloat(coinTicker.PercentChange1H, 64)
			if priceChange > 1 {
				s := fmt.Sprintf("Alert ("+pair+")  \n Price: %s (USD) %s (BTC) \n Percent Change 1H %s", coinTicker.PriceUsd, coinTicker.PriceBtc, coinTicker.PercentChange1H)
				sendMessage(s)
				time.Sleep(time.Minute * 30)
			}
			if priceChange < -5 {
				s := fmt.Sprintf("Alert ("+pair+")  \n Price: %s (USD) %s (BTC) \n Percent Change 1H %s", coinTicker.PriceUsd, coinTicker.PriceBtc, coinTicker.PercentChange1H)
				sendMessage(s)
				time.Sleep(time.Minute * 30)
			}
		}
		time.Sleep(time.Minute * 5)
	}

}

func sendMessage(message string) {
	bot, err := tgbotapi.NewBotAPI("552268599:AAHBd9PEyiQ5XB_o5DNBU8gp_kRArXDt8ms")
	if err != nil {
		log.Panic(err)
	}
	//msg := tgbotapi.NewMessage(304403970, message)
	msg := tgbotapi.NewMessage(-295823428, message)
	bot.Send(msg)
}
