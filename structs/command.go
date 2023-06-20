package structs

import (
	"encoding/json"
	"errors"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/config"
	mongof "main/functions/mongo"
)

/*
STORE COMMANDS AS A MAP
*/
type Command struct {
	Name        string `json:"name,omitempty"`
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Response    string `json:"response,omitempty"`
	Registered  bool   `json:"registered"`
}

type CommandObject struct {
	Commands map[string]Command `json:"commands"`
	GuildID  string             `json:"id"`
}

func NewCommand(name string, description string, response string) *Command {
	return &Command{Name: name, Description: description, Response: response}
}
func (c *Command) NameLen() int {
	return len(c.Name)
}
func (c *Command) DescriptionLen() int {
	return len(c.Description)
}
func (c *Command) ValidName() bool {
	return c.NameLen() >= 1 && c.NameLen() <= 32
}
func (c *Command) ValidDescription() bool {
	return c.NameLen() >= 1 && c.NameLen() <= 100
}
func (c *Command) ValidResponse() bool {
	return c.Response != ""
}

func ConvertCommandToProperCommand(object Command) (discord.SlashCommandCreate, error) {
	b, err := json.Marshal(object)
	if err != nil {
		return discord.SlashCommandCreate{}, err
	}
	list := discord.SlashCommandCreate{}
	if err := json.Unmarshal(b, &list); err != nil {
		return discord.SlashCommandCreate{}, err
	}
	return list, err
}
func ConvertCommandsToProperCommands(object CommandObject) ([]discord.ApplicationCommandCreate, error) {
	b, err := json.Marshal(&object.Commands)
	if err != nil {
		return nil, err
	}
	var list []discord.ApplicationCommandCreate
	if err := json.Unmarshal(b, &list); err != nil {
		return nil, err
	}
	return list, err
}
func RegisterCommand(guild snowflake.ID, command discord.SlashCommandCreate) (discord.ApplicationCommand, error) {
	cmd, err := config.BotClient.Rest().CreateGuildCommand(config.BotClient.ApplicationID(), guild, command)
	return cmd, err
}
func SetCommands(guild snowflake.ID, commands []discord.ApplicationCommandCreate) ([]discord.ApplicationCommand, error) {
	cmd, err := config.BotClient.Rest().SetGuildCommands(config.BotClient.ApplicationID(), guild, commands)
	return cmd, err
}

/*
DATABASE FUNCTIONS
NEW INSERTION +
UPDATE +
REPLACE +
DELETE

SCHEMA
CommandObject {}
*/
func (co *CommandObject) Fetch(id snowflake.ID) error {
	data, err := mongof.FindOne(bson.M{"id": id.String()}, options.FindOne(), config.MongoDatabase, "commands")
	if err != nil {
		return err
	}
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, &co); err != nil {
		return err
	}
	return nil

}
func NewCommandObjectGlobal(guild snowflake.ID) error {
	data, err := mongof.FindOne(bson.M{"id": guild.String()}, options.FindOne(), config.MongoDatabase, "commands")

	if data["id"] == nil {
		_, err = mongof.InsertOne(bson.M{"id": guild.String(), "commands": map[string]any{}}, options.InsertOne(), config.MongoDatabase, "commands")
		if err != nil {
			return errors.New("failed creating a server commands folder")
		}
	}
	return nil
}
func (c *Command) Register(guild snowflake.ID) error {
	data, err := RegisterCommand(guild, discord.SlashCommandCreate{
		Name:        c.Name,
		Description: c.Description,
	})
	if err != nil {
		return errors.New("failed registering command")
	}
	c.ID = data.ID().String()
	c.Registered = true
	//if c.Exists(guild) {
	return c.Update(guild)
	//}
	//maps, _ := c.ToMap()
	//_, err = mongof.InsertOne(maps, options.InsertOne(), config.MongoDatabase, "commands")
	//return err
}
func (c *Command) DeleteCommand(guild snowflake.ID) error {
	if err := c.Find(guild); err != nil {
		return err
	}
	id, err := snowflake.Parse(c.ID)
	if err != nil {
		return errors.New("invalid command id")
	}

	if err := config.BotClient.Rest().DeleteGuildCommand(config.BotClient.ApplicationID(), guild, id); err != nil {
		return errors.New("failed deleting the command")
	}
	//config.BotClient.Rest().DeleteGuildCommand(config.BotClient.ApplicationID(), guild)
	_, err = mongof.UpdateOne(bson.M{
		"$unset": bson.M{
			"commands." + c.Name: "",
		},
	}, bson.M{
		"id": guild.String(),
	}, options.Update(), config.MongoDatabase, "commands")
	if err != nil {
		return errors.New("failed removing command")
	}

	return nil
}
func (c *Command) Update(guild snowflake.ID) error {
	//if !c.Exists(guild) {
	//	return errors.New("command doesn't exist")
	//}
	b, err := json.Marshal(c)
	if err != nil {
		return errors.New("failed updating command")
	}
	var data = map[string]any{}
	if err = json.Unmarshal(b, &data); err != nil {
		return errors.New("failed updating command")
	}
	_, err = mongof.UpdateOne(bson.M{
		"$set": bson.M{
			"commands." + c.Name: data,
		},
	}, bson.M{
		"id": guild.String(),
	}, options.Update().SetUpsert(true), config.MongoDatabase, "commands")
	if err != nil {
		return errors.New("failed updating command")
	}
	return nil
}
func (c *Command) Find(guildid snowflake.ID) error {
	data, err := mongof.FindOne(bson.M{
		"id": guildid.String(),
	}, options.FindOne(), config.MongoDatabase, "commands")
	if _, ok := data["commands"].(primitive.M)[c.Name]; !ok || err != nil {
		return errors.New("command doesn't exist")
	}
	b, err := json.Marshal(data["commands"].(primitive.M)[c.Name])
	if err != nil {
		return errors.New("command doesn't exist")
	}
	if err = json.Unmarshal(b, &c); err != nil {
		return errors.New("command doesn't exist")
	}
	return nil
}
func (c *Command) Exists(guildid snowflake.ID) bool {
	data, err := mongof.FindOne(bson.M{
		"id": guildid.String(),
	}, options.FindOne(), config.MongoDatabase, "commands")
	if err != nil || data["id"] == "" {
		return false
	}
	if _, ok := data["commands"].(primitive.M)[c.Name]; !ok {
		return ok
	}
	return true
}
func (c *Command) ToMap() (map[string]interface{}, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
