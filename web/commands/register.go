package commands

import (
	"github.com/disgoorg/snowflake/v2"
	"github.com/gin-gonic/gin"
	"main/structs"
)

//POST /api/guilds/:id/commands
/*
body:
*/
func RegisterCommand(c *gin.Context) {
	/*



		CHECK IF USE AUTHENTICATED AND HAS PERMISSIONS TO GUILD OTHERWISE CANCEL




	*/
	guild, err := snowflake.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid server id"})
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
		c.JSON(400, gin.H{"error": "empty response"})
		return
	}
	err = structs.NewCommandObjectGlobal(guild)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = cmd.Register(guild)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, cmd)
	//cmd := structs.CommandObject{Commands: []structs.Command{
	//	{
	//		Name:        "hello",
	//		Description: "hell dwadawdwao",
	//		ID:          "1118163260638773319",
	//		Response:    `{"content": "hello"}`,
	//	},
	//}}
	//arr := []discord.ApplicationCommandCreate{}
	//for _, v := range cmd.Commands {
	//	arr = append(arr, discord.SlashCommandCreate{
	//		Name:        v.Name,
	//		Description: v.Description,
	//	})
	//}
	//commands, err := structs.ConvertCommandsToProperCommands(cmd) // issue
	//fmt.Println(cmd, commands)
	//if err != nil {
	//	c.JSON(400, err)
	//	return
	//}
	//_, err = structs.SetCommands(guild, []discord.ApplicationCommandCreate{
	//	discord.SlashCommandCreate{
	//		Name:        "",
	//		Description: "",
	//	},
	//})
	//fmt.Println(err)
	//c.JSON(200, err)
}
