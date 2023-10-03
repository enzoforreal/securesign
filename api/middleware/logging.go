package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware est un middleware qui consigne chaque requête à l'API.
func LoggingMiddleware(c *gin.Context) {
	// Horodatage de début pour mesurer la durée de la requête
	startTime := time.Now()

	// Exécute la requête
	c.Next()

	// Calcule la durée de la requête
	duration := time.Since(startTime)

	// Obtient des détails sur la requête
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	method := c.Request.Method
	statusCode := c.Writer.Status()

	// Formatte et consigne la requête
	logMessage := method + " " + path
	if raw != "" {
		logMessage = logMessage + "?" + raw
	}
	logMessage = logMessage + " - " + c.ClientIP() + " - " + c.Request.UserAgent() + " - " + duration.String()
	logMessage = logMessage + " - " + http.StatusText(statusCode)

	// Enregistre le message
	log.Println(logMessage)
}
