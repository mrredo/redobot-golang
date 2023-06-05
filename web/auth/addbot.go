package auth

import (
	"github.com/gin-gonic/gin"
	"os"
)

func AddBot(c *gin.Context) {
	c.Redirect(307, os.Getenv("BOT_URL")+"&"+"guild_id="+c.Query("guild_id"))
}
