package guilds

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"main/config"
)

func GetGuilds(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("token") != nil {
		guilds, err := config.AuthClient.GetGuilds(config.Sessions[session.Get("token").(string)])
		guilds1 := []Guild{}

		for _, v := range guilds {
			_, is := config.BotClient.Caches().Guild(v.ID)

			if !v.Permissions.Has(discord.PermissionManageGuild) {
				continue
			}
			guilds1 = append(guilds1, Guild{
				Id:          v.ID.String(),
				Name:        v.Name,
				BotInServer: is,
				Icon:        v.Icon,
			})
		}
		if err != nil {
			c.JSON(404, gin.H{
				"error": "failed fetching guilds",
			})
			return
		}
		c.JSON(200, guilds1)

	}

}

type Guild struct {
	Id          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Icon        *string `json:"icon,omitempty"`
	BotInServer bool    `json:"botInServer,omitempty"`
}
