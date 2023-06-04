package web

import (
	"fmt"
	"github.com/disgoorg/disgo/bot"
	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

func Start(client bot.Client) *gin.Engine {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello")
	})
	fmt.Println(r.Run("localhost:4000"))
	return r
}
