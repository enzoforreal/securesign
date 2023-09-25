package main

import (
	"GenSecureSign/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Static("/static", "./web/static")

	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", handlers.HomeHandler)

	r.POST("/send-email", handlers.SendEmailHandler)
	r.POST("/sign-document", handlers.SignDocumentHandler)

	r.Run(":8080")
}
