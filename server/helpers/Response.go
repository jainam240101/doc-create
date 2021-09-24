package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccessResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "Ok",
	})
}

func SendErrorResponse(c *gin.Context, status int, message interface{}) {
	finalStatus := http.StatusNotAcceptable
	if status != 0 {
		finalStatus = status
	}
	c.JSON(finalStatus, gin.H{
		"data":    nil,
		"message": message,
	})
}
