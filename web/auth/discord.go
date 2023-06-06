package auth

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"main/bot1"
	"main/functions"
	"net/http"
	//"github.com/imroc/req/v3"
)

func AuthDiscord(c *gin.Context) {
	session1 := sessions.Default(c)
	var (
		code  = c.Query("code")
		state = c.Query("state")
	)
	if code != "" && state != "" {
		identifier := functions.RandStr(32)
		session, _, err := bot1.AuthClient.StartSession(code, state)
		if err != nil {
			fmt.Println(err)
			c.String(500, "Error starting session")
			return
		}
		bot1.Sessions[identifier] = session
		session1.Set("token", identifier)
		err1 := session1.Save()
		if err1 != nil {
			fmt.Println(err)
			c.String(404, "Failed saving session")
			return
		}

	}
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
