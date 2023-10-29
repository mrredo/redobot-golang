package cons

import (
	"github.com/disgoorg/disgo/events"
)

func GeneratePlaceholderCommand(e *events.ApplicationCommandInteractionCreate) (end map[string]any) {
	us := e.User()
	end = map[string]any{
		NewUser:        GenerateUserMap(&us),
		NewMember:      GenerateMemberMap(e.Member()),
		NewCommandData: GenerateCommandData(e.SlashCommandInteractionData()),
	}
	if g, ok := e.Guild(); ok {
		end[NewServer] = GenerateServerMap(&g)
	}
	return end
}
