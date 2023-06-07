package guilds

import (
	"encoding/json"
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
	"github.com/gin-gonic/gin"
	"main/config"
	"main/functions"
	"main/structs"
)

func JoinMessage(c *gin.Context) {
	/*

		CHECK IF USER HAS AUTHENTICATED


	*/

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
	}
	_, ok := config.BotClient.Caches().Guild(id)
	if !ok {
		c.JSON(400, gin.H{
			"error": "bot is not in this guild",
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
