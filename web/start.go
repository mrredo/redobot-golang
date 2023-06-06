package web

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"main/bot1"
	"main/web/auth"
	"main/web/guilds"
	"main/web/user"
	"net/http"
)

var (
	r = gin.Default()
)

func Start(client bot.Client) *gin.Engine {
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("token", store))
	api := r.Group("/api/")
	api.GET("/guilds", func(c *gin.Context) {
		guilds.GetGuilds(c)
	})
	api.PUT("/guilds/:id/:channel_id/:type", guilds.JoinMessage)
	r.GET("/addbot", auth.AddBot)
	api.GET("/user", func(c *gin.Context) {
		user.GetUser(c)
	})
	auth1 := r.Group("/auth/")
	auth1.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, bot1.AuthClient.GenerateAuthorizationURL("http://localhost:4000/auth/trylogin", discord.PermissionsNone, 0, false, discord.OAuth2ScopeIdentify, discord.OAuth2ScopeGuilds))
	})
	auth1.GET("/trylogin", func(c *gin.Context) {
		auth.AuthDiscord(c)
	})
	return r
}
