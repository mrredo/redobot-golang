package commands

import (
	"github.com/disgoorg/snowflake/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"main/config"
	"main/functions"
	"main/structs"
)

func GetSingleCommand(c *gin.Context) {
	session := sessions.Default(c)
	guild, err := snowflake.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid server id"})
		return
	}
	cmdName := c.Param("command")

	if !functions.IsUserInGuild(config.Sessions[session.Get("token").(string)], guild) {
		c.JSON(401, gin.H{"error": "you are not in this server and/or you dont have MANAGE_SERVER permission"})
		return
	}
	cmd := structs.Command{Name: cmdName}
	if err = cmd.Find(guild); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, cmd)
}
