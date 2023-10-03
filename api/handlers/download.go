package api

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DownloadHandler(c *gin.Context) {
	filename := c.Param("filename")
	path := filepath.Join(signedDirectory, filename)

	c.FileAttachment(path, filename)
}
