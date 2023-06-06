package guilds

import (
	"github.com/disgoorg/snowflake/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/bot1"
	mongof "main/functions/mongo"
)

func JoinMessage(c *gin.Context) {
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
	c.String(200, string(mType)+id.String())

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

func (msg *GuildMessage) Save() error {
	data, err := msg.FindInMongo()
	if err != nil {
		return err
	}
	if data["id"] == "" {
		_, err := msg.CreateInMongo()
		return err

		//create new message
	} else {
		_, err := msg.UpdateInMongo()
		return err
	}

}
func (msg *GuildMessage) UpdateInMongo() (*mongo.UpdateResult, error) {
	return mongof.UpdateOne(msg, bson.M{
		"id":         msg.ID,
		"channel_id": msg.ChannelID,
	}, options.Update(), bot1.MongoDatabase, "messages")
}
func (msg *GuildMessage) CreateInMongo() (*mongo.InsertOneResult, error) {
	return mongof.InsertOne(msg, options.InsertOne(), bot1.MongoDatabase, "messages")
}
func (msg *GuildMessage) FindInMongo() (bson.M, error) {
	return mongof.FindOne(bson.M{
		"id":         msg.ID,
		"channel_id": msg.ChannelID,
	}, options.FindOne(), bot1.MongoDatabase, "messages")
}
