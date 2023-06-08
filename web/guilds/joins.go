package guilds

import (
	"encoding/json"
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/oauth2"
	"github.com/disgoorg/snowflake/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"main/config"
	"main/functions"
	"main/structs"
)

func GetMessage(c *gin.Context) {
	mType, b := structs.MsgType(c.Param("type"))

	if !b {
		c.JSON(400, gin.H{
			"error": "invalid message type",
		})
		return
	}
	id, err := snowflake.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed parsing guild id",
		})
		return
	}
	_, ok := config.BotClient.Caches().Guild(id)
	if !ok {
		c.JSON(400, gin.H{
			"error": "bot is not in this guild",
		})
		return

	}
	guilds, _ := config.AuthClient.GetGuilds(config.Sessions[sessions.Default(c).Get("token").(string)])
	var good bool
	for _, guild := range guilds {
		if guild.Permissions.Has(discord.PermissionManageGuild) && guild.ID.String() == id.String() {
			good = true
		}
	}
	if !good {
		c.JSON(400, gin.H{
			"error": "either you dont have MANAGE_SERVER permission or you are not in this guild",
		})
		return
	}
	msg := &structs.GuildMessage{ID: id.String(), Type: mType}
	err = msg.FetchData()
	if err != nil {
		c.JSON(404, gin.H{
			"error": "document not found",
		})
		return
	}

	/*
		check if user has access to this information by checking if hes in the guild and has manage server perm

	*/
	msg.JsonData, _ = functions.Base64Decode(msg.JsonData)
	c.JSON(200, msg)
}
func UserIsInServerAndHasPerms(session oauth2.Session, supposedGuildId snowflake.ID) {

}
func JoinMessagePUT(c *gin.Context) {

	if sessions.Default(c).Get("token") == nil {
		c.JSON(401, gin.H{
			"error": "not logged in",
		})
		return
	}
	mType, b := structs.MsgType(c.Param("type"))

	if !b {
		c.JSON(400, gin.H{
			"error": "invalid message type",
		})
		return
	}
	id, err := snowflake.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed parsing guild id",
		})
		return
	}
	_, ok := config.BotClient.Caches().Guild(id)
	if !ok {
		c.JSON(400, gin.H{
			"error": "bot is not in this guild",
		})
		return

	}
	guilds, _ := config.AuthClient.GetGuilds(config.Sessions[sessions.Default(c).Get("token").(string)])
	var good bool
	for _, guild := range guilds {
		if guild.Permissions.Has(discord.PermissionManageGuild) && guild.ID.String() == id.String() {
			good = true
		}
	}
	if !good {
		c.JSON(400, gin.H{
			"error": "either you dont have MANAGE_SERVER permission or you are not in this guild",
		})
		return
	}

	Chanid, err := snowflake.Parse(c.Param("channel_id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed parsing guild id",
		})
	}
	if channel, ok := config.BotClient.Caches().Channel(Chanid); !ok {
		if channel.GuildID().String() != id.String() {
			c.JSON(400, gin.H{
				"error": "channel is not part of the guild",
			})
			return
		}
		c.JSON(400, gin.H{
			"error": "bot doesnt have access to this channel",
		})
		return

	}
	data := discord.MessageCreate{}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"error": "failed parsing json data or invalid json data",
		})
		fmt.Println(err)
		return
	}
	bd, err1 := json.Marshal(&data)
	if err1 != nil {
		c.JSON(400, gin.H{
			"error": "failed parsing json data 2",
		})
		return
	}

	messageData := make(map[string]any)
	if err := json.Unmarshal(bd, &messageData); err != nil {
		c.JSON(400, gin.H{
			"error": "failed parsing json data 3",
		})
		return
	}
	sd, err2 := functions.MapToJSONString(messageData)
	if err2 != nil {
		c.JSON(400, gin.H{
			"error": "failed parsing json data 4",
		})
		return
	}
	//convert message data to discord message and back then strip the data
	StripDiscordMessageOfUnwantedInformation(messageData)
	guildmsg := &structs.GuildMessage{
		ID:        id.String(),
		JsonData:  functions.Base64Encode(sd),
		ChannelID: Chanid.String(),
		Type:      mType,
		Enabled:   c.Query("enabled") == "true",
	}

	datas, _ := guildmsg.FindInMongo()
	if datas == nil {
		_, err := guildmsg.CreateInMongo()
		if err != nil {
			c.JSON(400, err)
			return
		}
	} else {

		_, err := guildmsg.UpdateInMongo()
		if err != nil {
			c.JSON(400, err)
			return
		}
	}
	//fmt.Println(data["id"] == "")
	//if data["id"] == "" {
	//	d, err := guildmsg.CreateInMongo()
	//	if err != nil {
	//		c.JSON(400, err)
	//		return
	//	}
	//	fmt.Println(d)
	//	//create new message
	//} else {
	//	_, err := guildmsg.UpdateInMongo()
	//	if err != nil {
	//		c.JSON(400, err)
	//		return
	//	}
	//
	//}
	c.String(200, "lol")

}
