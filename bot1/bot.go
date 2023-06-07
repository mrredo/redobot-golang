package bot1

import (
	"context"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/oauth2"
	"github.com/disgoorg/disgo/rest"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
	"main/bot1/event"
	"main/config"
	"os"
)

func StrToSnowflake(id string) snowflake.ID {
	id1, _ := snowflake.Parse(id)
	return id1
}
func Start() bot.Client {
	log.Info("starting example...")
	log.Infof("disgo version: %s", disgo.Version)

	config.AuthClient = oauth2.New(StrToSnowflake(os.Getenv("ID")), os.Getenv("SECRET"), oauth2.WithRestClientConfigOpts(rest.WithHTTPClient(config.HttpClient)))
	client, err := disgo.New(os.Getenv("TOKEN"),
		bot.WithGatewayConfigOpts(gateway.WithIntents(
			gateway.IntentGuilds,
			gateway.IntentGuildMessages,
			gateway.IntentDirectMessages,
			gateway.IntentMessageContent,
			gateway.IntentGuildMembers,
			gateway.IntentsAll,
		)),
		bot.WithCacheConfigOpts(
			cache.WithCaches(cache.FlagGuilds, cache.FlagChannels, cache.FlagMembers),
		),
		bot.WithEventListenerFunc(func(e *events.Ready) {

		}),
		bot.WithEventListenerFunc(event.UserJoin),
	)
	config.BotClient = client
	if err != nil {
		log.Fatal("error while building bot1: ", err)
	}

	defer config.BotClient.Close(context.TODO())

	return config.BotClient
}
