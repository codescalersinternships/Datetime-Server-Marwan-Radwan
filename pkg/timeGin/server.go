package timegin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	routes := r.Group("/datetime")
	{
		routes.GET("/json", getTimeJsonHandler)
		routes.GET("/plain", getTimePlainHandler)
	}
	// r.GET("/datetime/json", getTimeJsonHandler)

	return r

}

// getTimeJsonHandler returns an HTTP handler function that responds with the current time in json format.
func getTimeJsonHandler(ctx *gin.Context) {
	currentTime := time.Now().Format(time.RFC822)

	ctx.JSON(http.StatusOK, gin.H{
		"DateTime": currentTime,
	})
}

// getTimePlainHandler returns an HTTP handler function that responds with the current time in plain text.
func getTimePlainHandler(ctx *gin.Context) {
	currentTime := time.Now().Format(time.RFC822)

	ctx.String(http.StatusOK, currentTime)
}
