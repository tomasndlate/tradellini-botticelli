package main

import (
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World, \nWelcome to Alpaca Golang Demo!")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading Alpaca API keys from .env file")
	}

	alpacaClient := alpaca.NewClient(alpaca.ClientOpts{
		ApiKey:    os.Getenv("ALPACA_ACCESS_ID"),
		ApiSecret: os.Getenv("ALPACA_SECRET_KEY"),
		BaseURL:   "https://paper-api.alpaca.markets",
	})

	fmt.Println(alpacaClient.GetAccount())
}
