package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	tiff "github.com/chai2010/tiff"
	"github.com/jung-kurt/gofpdf"
)

//main ...
func main() {
	t0 := time.Now()
	ImageFile := "./000004.tif"
	arquivoPDF := ConverteEGeraPdf(ImageFile)
	fmt.Printf("Pdf gerado %s \n", arquivoPDF)
	t1 := time.Now()
	fmt.Printf("Total time %v to run. \n", t1.Sub(t0))
}

//ConverteEGeraPdf ...
func ConverteEGeraPdf(ImageFile string) (PdfFile string) {
	t0 := time.Now()

	dataofimage := ImageRead2(ImageFile)
	ImageFile = Formatpng2(dataofimage)
	fmt.Println("Converting  image to png  ...")

	t1 := time.Now()
	fmt.Printf("Tempo para formatar %v to run. \n", t1.Sub(t0))

	arquivoPdf := FormatarImagemPdf(ImageFile)
	fmt.Println("Converting  image to ppdf ...")
	t1 = time.Now()
	fmt.Printf("Tempo para gerar PDF %v to run. \n", t1.Sub(t0))

	//fmt.Printf("Go launched at %s\n", t0.Local())
	//fmt.Printf("Go terminated at %s\n", t1.Local())

	return arquivoPdf
}

//ImageRead2 ...
func ImageRead2(ImageFile string) (image image.Image) {
	// open "test.jpg"
	file, err := os.Open(ImageFile)
	if err != nil {
		log.Fatal(err)
	}
	// decode jpeg into image.Image
	img, err := tiff.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	return img
}

//Formatpng2 ...
func Formatpng2(img image.Image) (image string) {
	r := rand.New(rand.NewSource(99))

	buffer := bytes.NewBufferString("")
	fmt.Fprint(buffer, "./temp/converterdTOPNG", r.Float32(), ".png")

	out, err := os.Create(buffer.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("\n success... \n ")
	return buffer.String()
}

//FormatarImagemPdf ...
func FormatarImagemPdf(img string) (pdfFile string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetHeaderFunc(func() {
		pdf.Image(img, 0, 0, 210, 297, false, "", 0, "")
		pdf.SetY(5)
	})

	fileStr := img + ".pdf"
	err := pdf.OutputFileAndClose(fileStr)
	Summary(err, fileStr)
	// Output:
	// Successfully generated pdf/Fpdf_AddPage.pdf

	return fileStr
}

//Summary ...
func Summary(err error, fileStr string) {
	if err == nil {
		err = referenceCompare(fileStr)
	}
	if err == nil {
		fileStr = filepath.ToSlash(fileStr)
		fmt.Printf("Successfully generated %s\n", fileStr)
	} else {
		fmt.Println(err)
	}
}

//referenceCompare ...
func referenceCompare(fileStr string) (err error) {
	var refFileStr, refDirStr, dirStr, baseFileStr string
	dirStr, baseFileStr = filepath.Split(fileStr)
	refDirStr = filepath.Join(dirStr, "reference")
	err = os.MkdirAll(refDirStr, 0755)
	if err == nil {
		refFileStr = filepath.Join(refDirStr, baseFileStr)
		err = gofpdf.ComparePDFFiles(fileStr, refFileStr)
	}
	return
}
