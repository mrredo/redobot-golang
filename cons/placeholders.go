package cons

import (
	"fmt"
	"regexp"
)

type Placeholder string

const (
	User                Placeholder = "{user}"
	UserMention         Placeholder = "{usermention}"
	UserID              Placeholder = "{userid}"
	Server              Placeholder = "{server}"
	ServerOwner         Placeholder = "{serverowner}"
	ServerOwnerMention  Placeholder = "{serverownermention}"
	JoinDate            Placeholder = "{joindate}"
	ServerIcon          Placeholder = "{servericon}"
	UserIcon            Placeholder = "{usericon}"
	MemberCountCurrent  Placeholder = "{currentmembers}"
	MemberCountPrevious Placeholder = "{lastmembers}"
	CommandName         Placeholder = "{commandname}"
)

var PlaceholderLists = map[string][]Placeholder{
	"command": {
		CommandName, User, UserMention, UserID, Server, ServerOwner, ServerOwnerMention, JoinDate, ServerIcon, UserIcon,
	},
	"message": {
		User, UserMention, UserID, Server, ServerOwner, ServerOwnerMention, JoinDate, ServerIcon, UserIcon, MemberCountPrevious, MemberCountCurrent,
	},
}
var (
	PlaceHolderRegex = regexp.MustCompile(`{\w+}`)
)

func FindPlaceHoldersAndReplace(s string, placeholder map[Placeholder]any) string {
	slist := PlaceHolderRegex.FindAllString(s, -1)
	news := s
	for _, v := range slist {
		//fmt.Println(slist)
		reg := regexp.MustCompile(v)
		news = reg.ReplaceAllString(news, fmt.Sprintf("%v", placeholder[Placeholder(v)]))
		//news = PlaceHolderRegex.ReplaceAllString(news, fmt.Sprintf("%v", placeholder[Placeholder(v)]))
	}
	return news
}
