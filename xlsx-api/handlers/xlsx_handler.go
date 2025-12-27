package handlers

import (
	"github.com/gin-gonic/gin"
	"xlsx/models"
	"xlsx/services"
	"os"
	"github.com/google/uuid"
	"fmt"
)

const downloadDir = "./downloads"

func ensureDir() {
	_ = os.MkdirAll(downloadDir, 0755)
}

func XLSXHandler(c *gin.Context) {
	var req models.XLSXRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error":"Invalid JSON format"})
		return
	}
	f, err := services.BuildXLSX(req)
	if err != nil {
		c.JSON(503, gin.H{
			"error":"Something went wrong. Please try again.\n"+err.Error(),
		})
		return
	}
	defer f.Close()

	ensureDir()
	filename := uuid.New().String() + ".xlsx"
	path := downloadDir + "/" + filename

	if err := f.SaveAs(path); err != nil {
		c.JSON(500, gin.H{"error":"File save failed"})
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	host := c.Request.Host

	fullURL := fmt.Sprintf("%s://%s/downloads/%s", scheme, host, filename)

	c.JSON(200, gin.H{
		"url": fullURL,
	})
}