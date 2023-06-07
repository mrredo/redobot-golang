package structs

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/config"
	mongof "main/functions/mongo"
)

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
	Enabled   bool         `json:"enabled" bson:"enabled,omitempty"`
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
	}, options.Update(), config.MongoDatabase, "messages")
}
func (msg *GuildMessage) CreateInMongo() (*mongo.InsertOneResult, error) {
	return mongof.InsertOne(msg.ToMap(), options.InsertOne(), config.MongoDatabase, "messages")
}
func (msg *GuildMessage) FindInMongo() (bson.M, error) {
	return mongof.FindOne(bson.M{
		"id":   msg.ID,
		"type": msg.Type,
	}, options.FindOne(), config.MongoDatabase, "messages")
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
func (msg *GuildMessage) FetchData() error {
	data, err := msg.FindInMongo()
	if err != nil {
		return err
	}
	b, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &msg); err != nil {
		return err
	}
	return nil
}
