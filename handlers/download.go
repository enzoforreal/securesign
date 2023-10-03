package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DownloadHandler(c *gin.Context) {
	filename := c.Param("filename")

	if !isValidFilename(filename) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid file name.",
		})

		return
	}

	path := filepath.Join(signedDirectory, filename)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		c.String(http.StatusNotFound, "File not found.")
		return
	}

	c.FileAttachment(path, filename)
}
