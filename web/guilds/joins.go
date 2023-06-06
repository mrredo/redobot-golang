package guilds

import (
	"encoding/json"
	"github.com/disgoorg/snowflake/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/bot1"
	mongof "main/functions/mongo"
)

func JoinMessage(c *gin.Context) {
	/*

		CHECK IF USER HAS AUTHENTICATED


	*/

	mType, b := MsgType(c.Param("type"))

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
	_, ok := bot1.BotClient.Caches().Guild(id)
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
	if channel, ok := bot1.BotClient.Caches().Channel(Chanid); !ok {
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
	messageData := make(map[string]any)
	if err := c.BindJSON(&messageData); err != nil {
		c.JSON(400, gin.H{
			"error": "failed parsing json data",
		})
		return
	}
	//convert message data to discord message and back then strip the data
	StripDiscordMessageOfUnwantedInformation(messageData)
	guildmsg := &GuildMessage{
		ID:        id.String(),
		JsonData:  "",
		ChannelID: Chanid.String(),
		Type:      mType,
	}

	data, _ := guildmsg.FindInMongo()
	if data == nil {
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
func MsgType(s string) (MessageTypes, bool) {
	return MSGsTOmsg[s], MSGsTOmsg[s] != ""
}

var MSGsTOmsg = map[string]MessageTypes{
	"join":  JOIN,
	"leave": LEAVE,
}

type MessageTypes string

const (
	JOIN  MessageTypes = "join"
	LEAVE MessageTypes = "leave"
)

type GuildMessage struct {
	ID        string       `json:"id,omitempty" bson:"_id,omitempty"`
	JsonData  string       `json:"json_data,omitempty" bson:"json_data,omitempty"`
	ChannelID string       `json:"channel_id,omitempty" bson:"channel_id,omitempty"`
	Type      MessageTypes `json:"type,omitempty" bson:"type,omitempty"`
}

// func (msg *GuildMessage) Save() error {
//
// }
func (msg *GuildMessage) UpdateInMongo() (*mongo.UpdateResult, error) {
	return mongof.UpdateOne(bson.M{
		"$set": msg.ToMap(),
	}, bson.M{

		"id":   msg.ID,
		"type": msg.Type,
	}, options.Update(), bot1.MongoDatabase, "messages")
}
func (msg *GuildMessage) CreateInMongo() (*mongo.InsertOneResult, error) {
	return mongof.InsertOne(msg.ToMap(), options.InsertOne(), bot1.MongoDatabase, "messages")
}
func (msg *GuildMessage) FindInMongo() (bson.M, error) {
	return mongof.FindOne(bson.M{
		"id":   msg.ID,
		"type": msg.Type,
	}, options.FindOne(), bot1.MongoDatabase, "messages")
}
func (msg *GuildMessage) ToMap() bson.M {
	result := make(map[string]interface{})

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil
	}
	return result
}
