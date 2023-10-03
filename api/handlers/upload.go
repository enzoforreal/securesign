package api

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const uploadDirectory = "./static/storage/uploads"

func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("pdf")
	if err != nil {
		c.String(http.StatusBadRequest, "Erreur lors du téléchargement du fichier: %v", err)
		return
	}

	path := filepath.Join(uploadDirectory, file.Filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la sauvegarde du fichier: %v", err)
		return
	}

	c.String(http.StatusOK, "Fichier téléchargé avec succès!")
}
