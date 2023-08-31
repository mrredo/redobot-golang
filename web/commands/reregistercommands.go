package commands

import (
	"github.com/disgoorg/snowflake/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"main/structs"
)

func ReRegisterCommands(c *gin.Context) {
	session := sessions.Default(c)
	guild, err := snowflake.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid server id"})
		return
	}

	cmds := []structs.Command{}
	if err := c.BindJSON(&cmds); err != nil {
		c.JSON(400, gin.H{"error": "invalid command json"})
		return
	}
	cmd := structs.CommandObject{}
}
