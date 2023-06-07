package channel

import (
	"github.com/disgoorg/snowflake/v2"
	"github.com/gin-gonic/gin"
	"main/config"
)

var filters = map[string]map[int]string{
	"text": {
		0:  "GUILD_TEXT",
		1:  "DM",
		3:  "GROUP_DM",
		5:  "GUILD_ANNOUNCEMENT",
		10: "ANNOUNCEMENT_THREAD",
		11: "PUBLIC_THREAD",
		12: "PRIVATE_THREAD",
		15: "GUILD_FORUM",
	},
}

func GetChannels(c *gin.Context) {
	/*









		FILTER CHANNELS THAT USER DOESNT HAS ACCESS TO










	*/
	guildid, err := snowflake.Parse(c.Param("guild_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid guild id"})
		return
	}
	typeChan := c.Query("type")

	guild, ok := config.BotClient.Caches().Guild(guildid)
	if !ok {
		c.JSON(404, gin.H{"error": "guild not found"})
		return
	}

	channels, err := config.BotClient.Rest().GetGuildChannels(guild.ID)
	if _, ok := filters[typeChan]; !ok {
		c.JSON(200, channels)
		return
	}
	newChannels := []SimpleChannel{}
	for _, v := range channels {
		if _, ok := filters[typeChan][int(v.Type())]; ok {
			newChannels = append(newChannels, SimpleChannel{
				Name: v.Name(),
				ID:   v.ID().String(),
			})
		}
	}
	c.JSON(200, newChannels)
}

type SimpleChannel struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}
