package commands

import (
	"github.com/disgoorg/snowflake/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"main/config"
	"main/functions"
	"main/structs"
)

/*
/guilds/:id/commands
*/
func GetCommands(c *gin.Context) {
	guild, err := snowflake.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid server id"})
		return
	}
	if !functions.IsUserInGuild(config.Sessions[sessions.Default(c).Get("token").(string)], guild) {
		c.JSON(401, gin.H{"error": "you are not in this server and/or you dont have MANAGE_SERVER permission"})
		return
	}
	commands, err := config.BotClient.Rest().GetGuildCommands(config.BotClient.ApplicationID(), guild, false)
	if err != nil {
		c.JSON(404, gin.H{"error": "failed getting commands"})
		return
	}
	commandOb := structs.CommandObject{}
	if err := commandOb.Fetch(guild); err != nil {

		c.JSON(200, gin.H{})
		return
	}

	for _, c := range commandOb.Commands {
		c.Registered = false
		commandOb.Commands[c.Name] = c
		for range commands {
			c.Registered = true
			commandOb.Commands[c.Name] = c
			break

		}
	}

	c.JSON(200, commandOb.Commands)

}
