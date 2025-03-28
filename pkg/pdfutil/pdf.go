package pdfutil

import (
	"os"

	"github.com/ledongthuc/pdf"
)

// GetPDFPageCount returns the number of pages in a PDF file
func GetPDFPageCount(filePath string) (int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		return 0, err
	}

	pdfReader, err := pdf.NewReader(f, fileInfo.Size())
	if err != nil {
		return 0, err
	}

	return pdfReader.NumPage(), nil
}
