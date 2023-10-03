package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(c *gin.Context) {

	startTime := time.Now()

	c.Next()

	duration := time.Since(startTime)

	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	method := c.Request.Method
	statusCode := c.Writer.Status()

	logMessage := method + " " + path
	if raw != "" {
		logMessage = logMessage + "?" + raw
	}
	logMessage = logMessage + " - " + c.ClientIP() + " - " + c.Request.UserAgent() + " - " + duration.String()
	logMessage = logMessage + " - " + http.StatusText(statusCode)

	// Enregistre le message
	log.Println(logMessage)
}
