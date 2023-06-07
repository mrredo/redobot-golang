package channel

import "github.com/gin-gonic/gin"

func GetChannels(c *gin.Context) {
	c.Param("guild_id")
}
