package commands

import (
	"github.com/disgoorg/snowflake/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"main/config"
	"main/functions"
	"main/structs"
	"strings"
)

//POST /api/guilds/:id/commands
/*
body:
*/
func RegisterCommand(c *gin.Context) {

	typeQuery := strings.ToLower(c.Query("type"))
	if typeQuery != "update" && typeQuery != "register" {
		c.JSON(400, gin.H{"error": "invalid type | register/update are only supported"})

		return
	}
	guild, err := snowflake.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid server id"})
		return
	}
	if !functions.IsUserInGuild(config.Sessions[sessions.Default(c).Get("token").(string)], guild) {
		c.JSON(401, gin.H{"error": "you are not in this server and/or you dont have MANAGE_SERVER permission"})
		return
	}
	cmd := structs.Command{}
	if err = c.BindJSON(&cmd); err != nil {
		c.JSON(400, gin.H{"error": "invalid json body"})
		return
	}
	if !cmd.ValidName() {
		c.JSON(400, gin.H{"error": "invalid command name, must be from 1-32 characters long"})
		return
	}
	if !cmd.ValidDescription() {
		c.JSON(400, gin.H{"error": "invalid command description, must be from 1-100 characters long"})
		return
	}
	if !cmd.ValidResponse() {
		c.JSON(400, gin.H{"error": "invalid response"})
		return
	}
	err = structs.NewCommandObjectGlobal(guild)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err = structs.CheckCommandLimit(guild); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if typeQuery == "register" {
		if err := structs.CommandExists(cmd.Name, guild); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}
	cmd.Name = strings.ToLower(cmd.Name)
	err = cmd.Register(guild)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, cmd)

}
