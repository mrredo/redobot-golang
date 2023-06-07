package main

import (
	"context"
	"encoding/gob"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"main/bot1"
	"main/config"
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
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO")).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client1, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	config.MongoClient = client1
	config.MongoDatabase = client1.Database("redobot")
	defer func() {
		if err = client1.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	gob.Register(map[string]interface{}{})
	gob.Register(map[interface{}]interface{}{})
	client := bot1.Start()
	r := web.Start(client)
	s := make(chan os.Signal, 1)
	go func() {
		if err = client.OpenGateway(context.TODO()); err != nil {
			log.Fatal("error while connecting to gateway: ", err)
		}

		log.Println("example is now running. Press CTRL-C to exit.")

		// Block the goroutine to keep the bot1 running
		//select {}
	}()

	// Start a goroutine to run the Gin server
	go func() {
		err := r.Run("localhost:4000")
		if err != nil {
			panic(err)
		}
	}()

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
	// Block the main goroutine to keep the application running
	//select {}

}
