package web

import (
	"fmt"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"main/bot1"
	"main/web/auth"
	"net/http"
)

var (
	r = gin.Default()
)

func Start(client bot.Client) *gin.Engine {
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("token", store))

	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		cookie := session.Get("token")
		if cookie != nil {
			var user *discord.OAuth2User
			us, _ := bot1.AuthClient.GetUser(bot1.Sessions[cookie.(string)])
			user = us
			fmt.Println(11)
			c.String(200, *user.AvatarURL())
		}
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
