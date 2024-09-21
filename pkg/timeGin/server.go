package timegin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetTimeHandler returns an HTTP handler function that responds with the current time formatted depending on accept header.
func GetTimeHandler(ctx *gin.Context) {
	currentTime := time.Now().Format(time.RFC822)

	acceptHeader := ctx.GetHeader("Accept")
	if acceptHeader == "application/json" {
		ctx.JSON(http.StatusOK, gin.H{
			"DateTime": currentTime,
		})
		return
	}

	ctx.String(http.StatusOK, currentTime)
}

// StartServer initializes and returns a new Gin engine with predefined routes.
func StartServer() *gin.Engine {
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	r.GET("/datetime", GetTimeHandler)

	return r

}
