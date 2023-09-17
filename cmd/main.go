package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Initialise GIN
    r := gin.Default()

    // Définit un répertoire pour les fichiers statiques (CSS, JS)
    r.Static("/static", "./web/static")

    // Charge les modèles HTML du répertoire web/templates
    r.LoadHTMLGlob("web/templates/*")

    // Route pour la page d'accueil
    r.GET("/", homeHandler)

    // Routes pour l'envoi d'e-mails et la signature électronique
    r.POST("/send-email", sendEmailHandler)
    r.POST("/sign-document", signDocumentHandler)

    // Lance le serveur sur le port 8080
    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
