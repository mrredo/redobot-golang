package main

import (
	"context"
	"encoding/gob"
	"github.com/joho/godotenv"
	"log"
	"main/bot1"
	"main/web"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gob.Register(map[string]interface{}{})
	gob.Register(map[interface{}]interface{}{})
	client := bot1.Start()
	r := web.Start(client)
	go func() {
		if err = client.OpenGateway(context.TODO()); err != nil {
			log.Fatal("error while connecting to gateway: ", err)
		}

		log.Println("example is now running. Press CTRL-C to exit.")
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-s

		// Block the goroutine to keep the bot1 running
		select {}
	}()

	// Start a goroutine to run the Gin server
	go func() {
		err := r.Run("localhost:4000")
		if err != nil {
			panic(err)
		}
	}()

	// Block the main goroutine to keep the application running
	select {}
}
