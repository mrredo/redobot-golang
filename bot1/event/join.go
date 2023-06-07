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
	guild, ok := e.Client().Caches().Guild(e.GuildID)
	if !ok {
		return
	}
	owner, err := e.Client().Rest().GetMember(e.GuildID, guild.OwnerID)
	if err != nil {
		return
	}
	msg := &structs.GuildMessage{}
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
		cons.ServerIconPlaceholder:         *guild.IconURL(), //could throw an error if guild doest have a icon
		cons.UserIconPlaceholder:           *user.AvatarURL(),
	}
	newJsonFilledData := cons.FindPlaceHoldersAndReplace(decoded, placeholder)
	creator := discord.NewMessageCreateBuilder()
	b, err := json.Marshal(newJsonFilledData)
	if err != nil {
		return
	}
	if err := json.Unmarshal(b, &creator); err != nil {
		return
	}
	chanid, err := snowflake.Parse(msg.ChannelID)
	if err != nil {
		return
	}
	_, err = e.Client().Rest().CreateMessage(chanid, creator.Build())
	if err != nil {
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
