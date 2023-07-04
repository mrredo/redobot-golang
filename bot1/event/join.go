package event

import (
	"encoding/json"
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/snowflake/v2"
	"main/cons"
	"main/functions"
	"main/structs"
)

func UserJoin(e *events.GuildMemberJoin) {
	user := e.Member.User
	mem := e.Member

	guild, err := e.Client().Rest().GetGuild(e.GuildID, true)
	if err != nil {
		return
	}
	owner, err := e.Client().Rest().GetMember(e.GuildID, guild.OwnerID)
	if err != nil {
		return
	}
	msg := &structs.GuildMessage{
		ID:   guild.ID.String(),
		Type: structs.JOIN,
	}
	if err := msg.FetchData(); err != nil {
		return
	}
	decoded, err := functions.Base64Decode(msg.JsonData)
	if err != nil {
		return
	}

	var placeholder = map[cons.Placeholder]any{
		cons.UserPlaceholder:               e.Member.User.Username,
		cons.UserMentionPlaceholder:        fmt.Sprintf("<@%s>", e.Member.User.ID.String()),
		cons.UserIDPlaceholder:             user.ID.String(),
		cons.ServerPlaceholder:             guild.Name,
		cons.ServerOwnerPlaceholder:        owner.User.Username,
		cons.ServerOwnerMentionPlaceholder: fmt.Sprintf("<@%s>", owner.User.ID.String()),
		cons.JoinDatePlaceholder:           mem.JoinedAt.Second(),
		//cons.ServerIconPlaceholder:         *guild.Icon, //could throw an error if guild doest have a icon
		//cons.UserIconPlaceholder:           *user.Avatar,
		cons.MemberCountCurrent:  guild.MemberCount,
		cons.MemberCountPrevious: guild.MemberCount - 1,
	}
	if guild.IconURL() == nil {
		placeholder[cons.ServerIconPlaceholder] = "https://assets-global.website-files.com/6257adef93867e50d84d30e2/636e0a6ca814282eca7172c6_icon_clyde_white_RGB.svg"
	} else {
		placeholder[cons.ServerIconPlaceholder] = *guild.IconURL()
	}
	if user.AvatarURL() == nil {
		placeholder[cons.UserIconPlaceholder] = "https://assets-global.website-files.com/6257adef93867e50d84d30e2/636e0a6ca814282eca7172c6_icon_clyde_white_RGB.svg"
	} else {
		placeholder[cons.UserIconPlaceholder] = *user.AvatarURL()
	}
	newJsonFilledData := cons.FindPlaceHoldersAndReplace(decoded, placeholder)
	creator := discord.MessageCreate{}
	//b, err := json.Marshal(newJsonFilledData)
	//if err != nil {
	//	return
	//}
	//fmt.Println(5)
	chanid, err := snowflake.Parse(msg.ChannelID)
	if err != nil {
		return
	}
	if err := json.Unmarshal([]byte(newJsonFilledData), &creator); err != nil {

		_, err = e.Client().Rest().CreateMessage(chanid, discord.NewMessageCreateBuilder().SetContent("Invalid json data set for join messages, visit dashboard to fix.").SetEphemeral(true).Build())
		return
	}

	if !msg.Enabled {
		return
	}
	_, err = e.Client().Rest().CreateMessage(chanid, creator)
	if err != nil {
		fmt.Println(err)
		e.Client().Rest().CreateMessage(chanid, discord.NewMessageCreateBuilder().SetContent("Invalid json data set for join messages, visit dashboard to fix.").SetEphemeral(true).Build())
		return
	}
	//cons.FindPlaceHoldersAndReplace(s, placeholder)

	/*
		get json data string from guild id and join type then decode base64 then replace all placeholders like:
		- {user.name}
		- {user.id}
		etc.
	*/
}
