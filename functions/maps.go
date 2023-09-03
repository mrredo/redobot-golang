package functions

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func MapToJSONString(m map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	jsonString := string(jsonBytes)
	return jsonString, nil
}
func Error(msg string) gin.H {
	return gin.H{"error": msg}
}
