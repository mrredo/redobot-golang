package guilds

import "reflect"

func StripDiscordMessageOfUnwantedInformation(a map[string]any) {
	delete(a, "tts")
	delete(a, "message_refrence")
	delete(a, "components")
	delete(a, "files")
	delete(a, "payload_json")
	delete(a, "attachments")
	delete(a, "flags")
	if va, ok := a["embeds"]; ok && reflect.TypeOf(va).Kind() == reflect.Slice {
		newEmbed := []map[string]any{}
		for _, v := range va.([]any) {
			if reflect.TypeOf(v).Kind() == reflect.Map {
				delete(v.(map[string]any), "type")
				delete(v.(map[string]any), "video")
				delete(v.(map[string]any), "provider")
				newEmbed = append(newEmbed, v.(map[string]any))
			}
		}
		a["embeds"] = newEmbed

	}
}

//type DiscordMessage struct {
//	Content string
//	Embeds  []Embed
//}
//type Embed struct {
//	Title       string
//	Description string
//	Url         string
//	Color       int
//	Timestamp   string
//	Footer      Footer
//	Image       Image
//	Thumbnail   Thumbnail
//	Author      Author
//	Fields      []Field
//}
//type Field struct {
//	Name   string `json:"name,omitempty"`
//	Value  string `json:"value,omitempty"`
//	Inline bool   `json:"inline,omitempty"`
//}
//type Author struct {
//}
//type Image struct {
//}
//type Thumbnail struct {
//}
//type Footer struct {
//	Text    string
//	IconUrl string
//}
