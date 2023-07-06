package placeholders

import (
	"github.com/gin-gonic/gin"
	"main/cons"
)

/*
GET /api/placeholders/:type
type=command|message
RETURNS 200 []Placeholder
*/
func GetPlaceholders(c *gin.Context) {
	typePlaceholder := c.Param("type")
	data, ok := cons.PlaceholderLists[typePlaceholder]
	if !ok {
		c.JSON(404, gin.H{"error": "invalid placeholder type"})
		return
	}
	c.JSON(200, data)
}
