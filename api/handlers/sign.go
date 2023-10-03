package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/digitorus/pdfsign/sign"
	"github.com/gin-gonic/gin"
)

const signedDirectory = "./static/storage/signed"

func SignHandler(c *gin.Context) {
	filename := c.PostForm("filename")
	inputPath := filepath.Join(uploadDirectory, filename)
	outputPath := filepath.Join(signedDirectory, filename)

	// Utilisez ici votre logique pour récupérer la clé privée, le certificat, etc.
	// Je donne un exemple simplifié pour la démonstration.

	//privateKey := // Votre logique pour récupérer la clé privée
	//certificate :=  // Votre logique pour récupérer le certificat

	inputFile, err := os.Open(inputPath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de l'ouverture du fichier: %v", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la création du fichier signé: %v", err)
		return
	}
	defer outputFile.Close()

	err = sign.Sign(inputFile, outputFile, certificate, privateKey, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la signature du fichier: %v", err)
		return
	}

	c.String(http.StatusOK, "Fichier signé avec succès!")
}
