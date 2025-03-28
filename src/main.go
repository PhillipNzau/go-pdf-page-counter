package main

import (
	"fmt"
	"log"

	"github.com/phillipnzau/pageExtractor/pkg/pdfutil"
)

func main() {
	filePath := "../Pawawise feature.pdf" 

	pageCount, err := pdfutil.GetPDFPageCount(filePath)
	if err != nil {
		log.Fatalf("Error getting page count: %v", err)
	}
	fmt.Printf("The PDF has %d pages.\n", pageCount)
}
