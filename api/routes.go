package api

import (
	api "GenSecureSign/api/middleware"
	"GenSecureSign/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(api.LoggingMiddleware)

	r.POST("/upload", handlers.UploadHandler)
	r.POST("/sign", handlers.SignHandler)
	r.GET("/download/:filename", handlers.DownloadHandler)

	// Utilisation d'un préfixe spécifique pour les fichiers statiques.
	r.Static("/static", "./static")
	r.GET("/empty-pdf", func(c *gin.Context) {
		c.File("./tools/empty.pdf")
	})

	return r
}
