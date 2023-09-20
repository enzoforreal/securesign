package main

import (
	"GenSecureSign/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()

	
	r.Static("/static", "./web/static")

	
	r.LoadHTMLGlob("web/templates/*")

	
	r.GET("/", handlers.HomeHandler)

	
	r.POST("/send-email", handlers.SendEmailHandler)
	r.POST("/sign-document", handlers.SignDocumentHandler)

	
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
