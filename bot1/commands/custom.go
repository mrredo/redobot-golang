package commands

import (
	"encoding/json"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"main/structs"
)

func HandleCustomCommands(e *events.ApplicationCommandInteractionCreate) {
	cmd := &structs.CommandObject{}
	err := cmd.Fetch(*e.GuildID())
	if err != nil {
		return
	}
	data := e.SlashCommandInteractionData()
	name := data.CommandName()
	command, ok := cmd.Commands[name]
	_ = command
	if !ok {
		e.CreateMessage(discord.NewMessageCreateBuilder().SetContent("Command not found!").SetEphemeral(true).Build())
		return
	}
	creator := discord.MessageCreate{}

	if err := json.Unmarshal([]byte(command.Response), &creator); err != nil {

		e.CreateMessage(discord.NewMessageCreateBuilder().SetContent("Invalid json data set for custom command messages, visit dashboard to fix.").SetEphemeral(true).Build())
		return
	}
	err = e.CreateMessage(creator)
	if err != nil {
		e.CreateMessage(discord.NewMessageCreateBuilder().SetContent("Invalid json data set for custom command messages, visit dashboard to fix.").SetEphemeral(true).Build())
		return
	}

}
