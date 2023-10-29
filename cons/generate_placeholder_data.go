package cons

import (
	"encoding/json"
	"github.com/disgoorg/disgo/discord"
	"main/def"
)

func GenerateUserMap(user *discord.User) (userMap map[string]any) {
	b, err := json.Marshal(user)
	if err != nil {
		return map[string]any{}
	}
	if err := json.Unmarshal(b, &userMap); err != nil {
		return map[string]any{}
	}
	if user.AvatarURL() == nil {
		userMap["avatar_url"] = user.DefaultAvatarURL()
	} else {
		userMap["avatar_url"] = *user.AvatarURL()
	}
	if user.BannerURL() != nil {
		userMap["banner_url"] = *user.BannerURL()
	} else {
		userMap["banner_url"] = def.ServerIcon
	}

	return userMap
}
func GenerateMemberMap[T *discord.Member | *discord.ResolvedMember](member T) (userMap map[string]any) {
	b, err := json.Marshal(member)
	if err != nil {
		return map[string]any{}
	}

	if err := json.Unmarshal(b, &userMap); err != nil {
		return map[string]any{}
	}
	delete(userMap, "user")
	return userMap
}
func GenerateServerMap(server *discord.Guild) (serverMap map[string]any) {
	b, err := json.Marshal(server)
	if err != nil {
		return map[string]any{}
	}
	if server.MemberCount == 0 {
		server.MemberCount = server.ApproximateMemberCount
	}
	if err := json.Unmarshal(b, &serverMap); err != nil {
		return map[string]any{}
	}
	if server.IconURL() == nil {
		serverMap["icon_url"] = def.ServerIcon
	} else {
		serverMap["icon_url"] = *server.IconURL()
	}
	return serverMap
}
func GenerateCommandData(data discord.SlashCommandInteractionData) (dataMap map[string]any) {
	b, err := json.Marshal(data)
	if err != nil {
		return map[string]any{}
	}
	if err := json.Unmarshal(b, &dataMap); err != nil {
		return map[string]any{}
	}
	dataMap["name"] = data.CommandName()
	dataMap["id"] = data.CommandID().String()
	option := make(map[string]any, 0)
	for _, v := range data.Options {
		var Newv any
		b, _ := v.Value.MarshalJSON()
		err := json.Unmarshal(b, &Newv)
		if err != nil {
			continue
		}
		option[v.Name] = Newv
	}
	dataMap["options"] = option
	return
}
func GenerateRestServerMap(server *discord.RestGuild) (serverMap map[string]any) {
	b, err := json.Marshal(server)
	if err != nil {
		return map[string]any{}
	}
	if server.MemberCount == 0 {
		server.MemberCount = server.ApproximateMemberCount
	}
	if err := json.Unmarshal(b, &serverMap); err != nil {
		return map[string]any{}
	}
	if server.IconURL() == nil {
		serverMap["icon_url"] = def.ServerIcon
	} else {
		serverMap["icon_url"] = *server.IconURL()
	}
	return serverMap
}
