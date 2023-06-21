package web

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"main/config"
	"main/web/auth"
	"main/web/channel"
	"main/web/commands"
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

	api.Use(func(c *gin.Context) {
		session := sessions.Default(c)
		tok := session.Get("token")
		if tok != nil {
			if _, ok := config.Sessions[tok.(string)]; !ok {
				c.JSON(401, gin.H{
					"error": "Authorization required",
				})
				c.Abort()
			}
		} else {
			c.JSON(401, gin.H{
				"error": "Authorization required",
			})
			c.Abort()
		}
		c.Next()
	})
	api.POST("/guilds/:id/commands", commands.RegisterCommand)
	api.GET("/guilds/:id/commands/:command", commands.GetSingleCommand)
	api.GET("/guilds/:id/commands", commands.GetCommands)

	api.DELETE("/guilds/:id/commands", commands.DeleteCommand)
	/*

		when user messgaes load their current data and add enabled to front end so user can disable and enable it
		/api/guilds/:id/:channel_id/:type GET
		current data
	*/

	api.GET("/guilds/messages/:id/:type", guilds.GetMessage)
	api.GET("/guilds/:id/channels", channel.GetChannels)
	api.GET("/guilds", func(c *gin.Context) {
		guilds.GetGuilds(c)
	})
	api.PUT("/guilds/:id/:channel_id/:type", guilds.JoinMessagePUT)
	r.GET("/addbot", auth.AddBot)
	api.GET("/user", func(c *gin.Context) {
		user.GetUser(c)
	})
	auth1 := r.Group("/auth/")
	auth1.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, config.AuthClient.GenerateAuthorizationURL("http://localhost:4000/auth/trylogin", discord.PermissionsNone, 0, false, discord.OAuth2ScopeIdentify, discord.OAuth2ScopeGuilds))
	})
	auth1.GET("/trylogin", func(c *gin.Context) {
		auth.AuthDiscord(c)
	})
	return r
}
