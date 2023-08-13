package config

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/oauth2"
	"github.com/disgoorg/snowflake/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
)

var (
	clientID      = snowflake.GetEnv("client_id")
	clientSecret  = os.Getenv("client_secret")
	baseURL       = os.Getenv("base_url")
	logger        = log.Default()
	HttpClient    = http.DefaultClient
	AuthClient    oauth2.Client
	BotClient     bot.Client
	MongoClient   *mongo.Client
	MongoDatabase *mongo.Database
	Sessions      = map[string]oauth2.Session{}
)
