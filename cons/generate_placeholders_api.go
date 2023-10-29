package cons

import (
	"encoding/json"
	"github.com/disgoorg/disgo/discord"
)

//	func GenerateAPIPlacholder() map[string]any {
//		//fakeG, fakeU, fakeM, fakeC := discord.Guild{}, discord.User{}, discord.Member{}, discord.SlashCommandInteractionData{}
//		return map[string]any{}
//	}
func GenerateAPIGuild() (data map[string]any) {
	b, _ := json.Marshal(discord.Guild{})
	if err := json.Unmarshal(b, &data); err != nil {
		return map[string]any{}
	}
	return
}
func GenerateAPICommand() (data map[string]any) {
	b, _ := json.Marshal(discord.SlashCommandInteractionData{})
	if err := json.Unmarshal(b, &data); err != nil {
		return map[string]any{}
	}
	data["options"] = map[string]any{}
	return
}
func GenerateAPIUser() (data map[string]any) {
	b, _ := json.Marshal(discord.User{})
	if err := json.Unmarshal(b, &data); err != nil {
		return map[string]any{}
	}
	return
}
func GenerateAPIMember() (member map[string]any) {
	b, _ := json.Marshal(discord.Member{})
	if err := json.Unmarshal(b, &member); err != nil {
		return map[string]any{}
	}

	return
}
