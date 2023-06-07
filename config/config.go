package config

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/oauth2"
	"github.com/disgoorg/snowflake/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
	"time"
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
	Sessions      = map[string]oauth2.Session{
		"LZMKhogeaugjENjYqssYmAdUdBSkgrme": oauth2.Session{
			AccessToken:  "0od8WvRzyeVBZtUUUvZiiNNlGwOZxu ",
			RefreshToken: "ycFRqjlTXxIvYGel1YdT9FxaGXfw2o",
			Scopes:       []discord.OAuth2Scope{discord.OAuth2ScopeGuilds, discord.OAuth2ScopeIdentify},
			TokenType:    "Bearer",
			Expiration:   time.Time{},
		},
	}
)
