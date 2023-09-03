package guilds

import (
	"fmt"
	"github.com/disgoorg/snowflake/v2"
	"github.com/gin-gonic/gin"
	"main/config"
	"main/functions"
)

func IsBotInGuild(c *gin.Context) {
	fmt.Println(111)
	id, err := snowflake.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, functions.Error("invalid guild id"))
		return
	}
	_, ok := config.BotClient.Caches().Guild(id)
	if !ok {
		c.Status(404)
		return
	}
	c.Status(200)
}
