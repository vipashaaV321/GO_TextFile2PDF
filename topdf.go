package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	file := "test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("%s file not found", file)
	}
	pdf := gofpdf.New("P", "mm", "A4","")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.MultiCell(190, 5,string(content),"0","0",false)
	_ = pdf.OutputFileAndClose("test.pdf")
	fmt.Println("Pdf created!")
}
