package commands

import (
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func TesGuildMem(e *events.ApplicationCommandInteractionCreate) {
	guild, _ := e.Client().Caches().Guild(*e.GuildID())
	e.CreateMessage(discord.NewMessageCreateBuilder().SetContent(fmt.Sprintf("members: %d", guild.MemberCount)).SetEphemeral(true).Build())
}
