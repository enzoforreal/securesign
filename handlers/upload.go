package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("pdf")
	if err != nil {
		c.String(http.StatusBadRequest, "Erreur lors du téléchargement du fichier: %v", err)
		return
	}

	if !isValidFilename(file.Filename) {
		c.String(http.StatusBadRequest, "Nom de fichier non valide.")
		return
	}

	path := filepath.Join(uploadDirectory, file.Filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la sauvegarde du fichier: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Fichier téléchargé avec succès!",
	})

}
