package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/phillipnzau/pageExtractor/pkg/pdfutil"
)

const uploadDir = "uploads"

func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("pdf")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	// Ensure upload directory exists
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	// Save the uploaded file
	filePath := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Get PDF page count
	pageCount, err := pdfutil.GetPDFPageCount(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read PDF"})
		return
	}

	// Return JSON response
	c.JSON(http.StatusOK, gin.H{"pageCount": pageCount})
}

func main() {
	r := gin.Default()

	// Set up file upload route
	r.POST("/upload", uploadHandler)

	port := ":8080"
	fmt.Println("Server started on http://localhost" + port)
	log.Fatal(r.Run(port))
}
