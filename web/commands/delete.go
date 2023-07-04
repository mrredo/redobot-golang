package commands

import (
	"github.com/disgoorg/snowflake/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"main/config"
	"main/functions"
	"main/structs"
)

func DeleteCommand(c *gin.Context) {
	guild, err := snowflake.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid server id"})
		return
	}
	name := c.Param("command")
	if 1 <= len(name) && len(name) >= 32 {
		c.JSON(400, gin.H{"error": "invalid command name"})
		return
	}
	if !functions.IsUserInGuild(config.Sessions[sessions.Default(c).Get("token").(string)], guild) {
		c.JSON(401, gin.H{"error": "you are not in this server and/or you dont have MANAGE_SERVER permission"})
		return
	}
	cmd := structs.Command{Name: name}
	err = cmd.DeleteCommand(guild)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
