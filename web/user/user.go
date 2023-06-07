package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"main/config"
	"net/http"
)

func GetUser(c *gin.Context) {
	session := sessions.Default(c)
	tok := session.Get("token")
	if tok == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization is required"})
		return
	}
	user, err := config.AuthClient.GetUser(config.Sessions[tok.(string)])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization is required"})
		return
	}
	c.JSON(200, user)
}
