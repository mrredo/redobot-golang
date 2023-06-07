package cons

import (
	"fmt"
	"regexp"
)

type Placeholder string

const (
	UserPlaceholder               Placeholder = "{user}"
	UserMentionPlaceholder        Placeholder = "{usermention}"
	UserIDPlaceholder             Placeholder = "{userid}"
	ServerPlaceholder             Placeholder = "{server}"
	ServerOwnerPlaceholder        Placeholder = "{serverowner}"
	ServerOwnerMentionPlaceholder Placeholder = "{serverownermention}"
	JoinDatePlaceholder           Placeholder = "{joindate}"
	ServerIconPlaceholder         Placeholder = "{servericon}"
	UserIconPlaceholder           Placeholder = "{usericon}"
	MemberCountCurrent            Placeholder = "{currentmembers}"
	MemberCountPrevious           Placeholder = "{lastmembers}"
)

var (
	PlaceHolderRegex = regexp.MustCompile(`{\w+}`)
)

func FindPlaceHoldersAndReplace(s string, placeholder map[Placeholder]any) string {
	slist := PlaceHolderRegex.FindAllString(s, -1)
	news := s
	for _, v := range slist {
		news = PlaceHolderRegex.ReplaceAllString(news, fmt.Sprintf("%v", placeholder[Placeholder(v)]))
	}
	return news
}
