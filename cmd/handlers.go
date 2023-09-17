package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "GenSecureSign/pkg/email" 
    "GenSecureSign/pkg/sign"   
)

// Handler pour l'envoi d'e-mails
func sendEmailHandler(c *gin.Context) {
    var emailData map[string]interface{}

    // Récupérez les données de la requête POST pour configurer l'e-mail
    if err := c.BindJSON(&emailData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Données JSON invalides",
        })
        return
    }

    // Appelez la fonction d'envoi d'e-mails du package email
    success, err := email.SendMail(emailData)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Erreur lors de l'envoi de l'e-mail",
        })
        return
    }

    if success {
        c.JSON(http.StatusOK, gin.H{
            "message": "E-mail envoyé avec succès.",
        })
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Erreur lors de l'envoi de l'e-mail",
        })
    }
}
    
// Handler pour la signature électronique
func signDocumentHandler(c *gin.Context) {
    var signData map[string]interface{}

    // Récupérez les données de la requête POST pour configurer la signature
    if err := c.BindJSON(&signData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Données JSON invalides",
        })
        return
    }

    // Appelez la fonction de signature électronique du package sign
    documentID := int(signData["document_id"].(float64)) // Assurez-vous de convertir le document_id en int
    signer := signData["signer"].(string)
    document, err := sign.GetDocumentByID(documentID)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Erreur lors de la récupération du document",
        })
        return
    }

    if document == nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Document non trouvé",
        })
        return
    }

    // Vérifiez la signature électronique avant de l'appliquer
    if !sign.VerifySignature(signer) {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Signature électronique invalide",
        })
        return
    }

    // Appliquez la signature électronique
    err = sign.ApplySignature(document, signer)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Erreur lors de la signature électronique",
        })
        return
    }

    // Vous pouvez renvoyer les détails du document signé en réponse si nécessaire
    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("Signature électronique réussie pour le document %d par %s", document.ID, signer),
    })
}


// Handler pour la page d'accueil
func homeHandler(c *gin.Context) {
    // Ici, vous pouvez renvoyer la page d'accueil HTML située dans le répertoire web/templates
    // Assurez-vous que le modèle HTML affiche les fonctionnalités de votre application

    // Exemple de rendu de modèle HTML
    c.HTML(http.StatusOK, "index.html", gin.H{
        // Ajoutez des données que vous souhaitez afficher dans le modèle ici
    })
}


