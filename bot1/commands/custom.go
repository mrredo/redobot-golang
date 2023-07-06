package commands

import (
	"encoding/json"
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"main/cons"
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

	guild, _ := e.Guild()
	owner, err := e.Client().Rest().GetMember(*e.GuildID(), guild.OwnerID)

	creator := discord.MessageCreate{}
	var Placeholders = map[cons.Placeholder]any{
		cons.CommandName:        data.CommandName(),
		cons.User:               e.User().Username,
		cons.UserMention:        fmt.Sprintf("<@%s>", e.User().ID.String()),
		cons.UserID:             e.User().ID.String(),
		cons.Server:             guild.Name,
		cons.ServerOwner:        owner.User.Username,
		cons.ServerOwnerMention: fmt.Sprintf("<@%s>", owner.User.ID.String()),
		cons.JoinDate:           e.Member().JoinedAt.Second(),
	}
	if guild.IconURL() == nil {
		Placeholders[cons.ServerIcon] = "https://assets-global.website-files.com/6257adef93867e50d84d30e2/636e0a6ca814282eca7172c6_icon_clyde_white_RGB.svg"
	} else {
		Placeholders[cons.ServerIcon] = *guild.IconURL()
	}
	if e.User().AvatarURL() == nil {
		Placeholders[cons.UserIcon] = "https://assets-global.website-files.com/6257adef93867e50d84d30e2/636e0a6ca814282eca7172c6_icon_clyde_white_RGB.svg"
	} else {
		Placeholders[cons.UserIcon] = *e.User().AvatarURL()
	}
	command.Response = cons.FindPlaceHoldersAndReplace(command.Response, Placeholders)
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
