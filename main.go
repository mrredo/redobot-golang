package main

import (
	"github.com/joho/godotenv"
	"log"
	"main/bot"
	"main/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := bot.Start()
	web.Start(client)

}
