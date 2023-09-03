package functions

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/oauth2"
	"github.com/disgoorg/snowflake/v2"
	"main/config"
)

func IsUserInGuild(session oauth2.Session, guildId snowflake.ID) bool {
	guilds, _ := config.AuthClient.GetGuilds(session)
	for _, guild := range guilds {
		if (guild.Permissions.Has(discord.PermissionManageGuild) || guild.Owner) && guild.ID.String() == guildId.String() {
			return true
		}
	}
	return false
}
