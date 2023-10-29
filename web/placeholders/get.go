package placeholders

import (
	"github.com/gin-gonic/gin"
	"main/cons"
	"strings"
)

// GetPlaceholders
/*
GET /api/placeholders/:type

type=command|message

RETURNS 200 []Placeholder
*/
//DEPRECATED
func GetPlaceholders(c *gin.Context) {
	typePlaceholder := c.Param("type")
	data, ok := cons.PlaceholderLists[typePlaceholder]
	if !ok {
		c.JSON(404, gin.H{"error": "invalid placeholder type"})
		return
	}
	c.JSON(200, data)
}

// NewPlaceholders /*
/*
GET /api/placeholders/:type

type=guild|member|user|command

RETURNS 200 map[string]any

SYNTAX:

type={option1}

type={option1},{option2},{option3}
*/
func NewPlaceholders(c *gin.Context) {
	listTypes := strings.Split(c.Param("type"), ",")
	mapsPlace := map[string]func() map[string]any{
		"guild":   cons.GenerateAPIGuild,
		"member":  cons.GenerateAPIMember,
		"user":    cons.GenerateAPIUser,
		"command": cons.GenerateAPICommand,
	}
	typesDontExist := make([]string, 0)
	data := map[string]any{}
	for _, v := range listTypes {
		gen, ok := mapsPlace[strings.Replace(v, " ", "", -1)]
		if !ok {
			typesDontExist = append(typesDontExist, v)
			continue
		}
		data[v] = gen()

	}
	if len(typesDontExist) != 0 {
		c.JSON(404, gin.H{
			"error": "Invalid types: " + strings.Join(typesDontExist, ", "),
		})
		return
	}
	c.JSON(200, data)
}
