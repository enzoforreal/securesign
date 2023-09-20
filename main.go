package main

import (
	"GenSecureSign/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialise GIN
	r := gin.Default()

	// Définit un répertoire pour les fichiers statiques (CSS, JS)
	r.Static("/static", "./web/static")

	// Charge les modèles HTML du répertoire web/templates
	r.LoadHTMLGlob("web/templates/*")

	// Route pour la page d'accueil
	r.GET("/", handlers.HomeHandler)

	// Routes pour l'envoi d'e-mails et la signature électronique
	r.POST("/send-email", handlers.SendEmailHandler)
	r.POST("/sign-document", handlers.SignDocumentHandler)

	// Lance le serveur sur le port 8080
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
