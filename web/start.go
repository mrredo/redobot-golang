package web

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"main/config"
	"main/web/auth"
	"main/web/channel"
	"main/web/commands"
	"main/web/guilds"
	"main/web/placeholders"
	"main/web/user"
	"main/web/webhooks"
	"net/http"
	"os"
)

var (
	r = gin.Default()
)

func Start(client bot.Client) *gin.Engine {
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("token", store))
	api := r.Group("/api/")
	r.Use(cors.New(cors.Config{
		//  AllowAllOrigins: true,
		AllowWildcard:          true,
		AllowWebSockets:        true,
		AllowBrowserExtensions: true,
		AllowOrigins:           []string{"http://localhost:4000", "http://localhost:3000", "http://127.0.0.1:3000", "http://127.0.0.1:4000", "https://127.0.0.1:443", "http://127.0.0.1:80", "http://192.168.8.114:4000", "https://redobot-golang-1.mrredogaming.repl.co"},
		AllowMethods:           []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "Accept", "Cookie", "Set-Cookie"},
		AllowCredentials:       true,
	}))
	r.GET("/checkout", webhooks.SessionThing)
	api.POST("/webhook", webhooks.HandleWebhook)
	//r.Use(csrf.Middleware(csrf.Options{
	//	Secret: "secretthing3535353533457754",
	//	ErrorFunc: func(c *gin.Context) {
	//		fmt.Println(sessions.Default(c).Get("csrfSalt"))
	//		c.String(400, "CSRF token mismatch")
	//		c.Abort()
	//	},
	//}))
	//
	//r.GET("/protected", func(c *gin.Context) {
	//	c.String(200, csrf.GetToken(c))
	//})
	//
	//r.POST("/protected", func(c *gin.Context) {
	//	c.String(200, "CSRF token is valid")
	//})

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
	api.GET("/placeholders/:type", placeholders.GetPlaceholders)
	api.POST("/guilds/:id/commands/reregister", commands.ReRegisterCommands)
	api.POST("/guilds/:id/commands", commands.RegisterCommand)
	api.GET("/guilds/:id/commands/:command", commands.GetSingleCommand)
	api.GET("/guilds/:id/commands", commands.GetCommands)

	api.DELETE("/guilds/:id/commands/:command", commands.DeleteCommand)
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
	api.GET("/guilds/:id/botinguild", guilds.IsBotInGuild)
	api.PUT("/guilds/:id/:channel_id/:type", guilds.JoinMessagePUT)
	r.GET("/addbot", auth.AddBot)
	api.GET("/user", func(c *gin.Context) {
		user.GetUser(c)
	})
	auth1 := r.Group("/auth/")
	auth1.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, config.AuthClient.GenerateAuthorizationURL(os.Getenv("AUTH_URL"), discord.PermissionsNone, 0, false, discord.OAuth2ScopeIdentify, discord.OAuth2ScopeGuilds))
	})
	auth1.GET("/logout", auth.Logout)
	auth1.GET("/trylogin", auth.AuthDiscord)

	return r
}
