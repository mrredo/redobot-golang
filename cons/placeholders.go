package cons

import (
	"fmt"
	"github.com/mrredo/govaluate"
	"regexp"
	"strings"
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

// DEPRECATED: new function FindReplacePlaceholders
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

const (
	NewPlaceholders = ""
	CommandOption   = ""
	NewUser         = "user"
	//NewUserMention         = "usermention"
	//NewUserID              = "userid"
	NewServer = "server"
	//NewServerOwner         = "serverowner"
	//NewServerOwnerMention  = "serverownermention"
	//NewJoinDate            = "joindate"
	//NewServerIcon          = "servericon"
	//NewUserIcon            = "usericon"
	//NewMemberCountCurrent  = "currentmembers"
	//NewMemberCountPrevious = "lastmembers"
	//NewCommandName         = "commandname"
)

var (
	NewPlaceholderRegex = regexp.MustCompile(`\{([^{}]+)\}`)
)

func TestThings() {

	stringss := []string{
		"{cmd}} {cmd1.go} {cmd2}",
	}
	place := map[string]any{
		"cmd": "e",
		"cmd1": map[string]any{
			"go": "eee",
		},
		"cmd2": map[string]any{},
	}
	_ = place
	for _, v := range stringss {
		fmt.Println(FindReplacePlaceholders(v, place))
	}
}
func FindReplacePlaceholders(text string, parameters map[string]any) string {
	matches := ExtractPlaceholders(text)
	for _, v := range matches {
		expr, _ := govaluate.NewEvaluableExpression(v[1:len(v)-1], false)
		res, err := expr.Evaluate(parameters)
		if err != nil {
			fmt.Println(err.Error())
		}
		text = strings.ReplaceAll(text, v, fmt.Sprintf("%v", res))
	}
	return text
}
func ExtractPlaceholders(text string) []string {
	return NewPlaceholderRegex.FindAllString(text, -1)
}

//func NewParsePlaceholders(text string, parameters map[string]any) {
//	expr, _ := govaluate.NewEvaluableExpressionWithFunctions(text, map[string]govaluate.ExpressionFunction{}, false)
//	fmt.Println(expr.Evaluate(parameters))
//}
