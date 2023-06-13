package structs

import (
	"github.com/disgoorg/snowflake/v2"
)

type Command struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Response    string `json:"response,omitempty"`
}

type CommandObject struct {
}

func NewCommand(name string, description string, response string) *Command {
	return &Command{Name: name, Description: description, Response: response}
}
func (c Command) Register(guild snowflake.ID) error {
	return RegisterCommands(guild, c)
}
func RegisterCommands(guild snowflake.ID, command ...Command) error {
	return nil
}
