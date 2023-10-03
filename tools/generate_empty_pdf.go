package main

import (
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	err := pdf.OutputFileAndClose("empty.pdf")
	if err != nil {
		panic(err)
	}
}
