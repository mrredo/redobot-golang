package bot1

import (
	"context"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/oauth2"
	"github.com/disgoorg/disgo/rest"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
	"net/http"
	"os"
)

var (
	clientID     = snowflake.GetEnv("client_id")
	clientSecret = os.Getenv("client_secret")
	baseURL      = os.Getenv("base_url")
	logger       = log.Default()
	httpClient   = http.DefaultClient
	AuthClient   oauth2.Client
	BotClient    bot.Client
	Sessions     = map[string]oauth2.Session{}
)

func StrToSnowflake(id string) snowflake.ID {
	id1, _ := snowflake.Parse(id)
	return id1
}
func Start() bot.Client {
	log.Info("starting example...")
	log.Infof("disgo version: %s", disgo.Version)

	AuthClient = oauth2.New(StrToSnowflake(os.Getenv("ID")), os.Getenv("SECRET"), oauth2.WithLogger(logger), oauth2.WithRestClientConfigOpts(rest.WithHTTPClient(httpClient)))
	client, err := disgo.New(os.Getenv("TOKEN"),
		bot.WithGatewayConfigOpts(gateway.WithIntents(
			gateway.IntentGuilds,
			gateway.IntentGuildMessages,
			gateway.IntentDirectMessages,
			gateway.IntentMessageContent)),
	)
	BotClient = client
	if err != nil {
		log.Fatal("error while building bot1: ", err)
	}

	defer BotClient.Close(context.TODO())

	return BotClient
}
