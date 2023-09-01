package auth

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"main/config"
	"main/functions"
	"net/http"
	"net/url"
	"os"
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
		session, _, err := config.AuthClient.StartSession(code, state)
		if err != nil {
			fmt.Println(err)
			c.String(500, "Error starting session")
			return
		}
		config.Sessions[identifier] = session
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
func Logout(c *gin.Context) {
	redirect := c.Query("redirect")
	session1 := sessions.Default(c)
	if session1.Get("token") != nil {
		delete(config.Sessions, session1.Get("token").(string))
		session1.Delete("token")
	}
	if redirect != "" {
		urlredirect, err := url.Parse(os.Getenv("BASE_URL") + redirect)
		if err != nil {
			c.String(403, "Invalid redirect URI, this only supports redirects to the same domain")
			return
		}
		c.Redirect(http.StatusSeeOther, urlredirect.Path)
		return
	}
	c.Status(200)
}
