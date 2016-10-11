package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"time"

	tiff "github.com/chai2010/tiff"
	"github.com/disintegration/imaging"
	"github.com/jung-kurt/gofpdf"
)

var files = []string{
	"./000001.tif",
}

func main() {
	t0 := time.Now()
	pdf()
	convertToTiffSimple()
	t1 := time.Now()

	fmt.Printf("The call toook %v to run. \n", t1.Sub(t0))

	ImageFile := "./000005.tif"
	dataofimage := ImageRead(ImageFile)
	Formatpng(dataofimage)
	fmt.Println("Converting  image to gif  ...")

	t1 = time.Now()
	fmt.Printf("The call toook %v to run. \n", t1.Sub(t0))
	fmt.Printf("Go launched at %s\n", t0.Local())
	fmt.Printf("Go terminated at %s\n", t1.Local())
}

//ImageRead ...
func ImageRead(ImageFile string) (image image.Image) {
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

func Formatpng(img image.Image) {
	out, err := os.Create("converterdTOPNG.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(img, "\n success... \n ")

}

func convertToJPG() {
	// load images and make 100x100 thumbnails of them
	var thumbnails []image.Image
	for _, file := range files {
		img, err := imaging.Open(file)

		if err != nil {
			panic(err)
		}
		thumb := imaging.Thumbnail(img, 800, 600, imaging.CatmullRom)
		thumbnails = append(thumbnails, thumb)
	}

	fmt.Println("Length ", len(thumbnails))

	// create a new blank image
	dst := imaging.New(800*len(thumbnails), 600, color.NRGBA{0, 0, 0, 0})

	// paste thumbnails into the new image side by side
	for i, thumb := range thumbnails {
		dst = imaging.Paste(dst, thumb, image.Pt(i*100, 0))
	}

	// save the combined image to file
	err := imaging.Save(dst, "dst.jpg")
	if err != nil {
		panic(err)
	}
}

func convertToTiffSimple() {
	for _, name := range files {
		// Load file data
		f, err := os.Open(name)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// Open tiff reader
		p, err := tiff.OpenReader(f)
		if err != nil {
			log.Println(err)
		}
		defer p.Close()

	}
}

func ImageFile(fileStr string) string {
	return filepath.Join("", "", fileStr)
}

func pdf() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetHeaderFunc(func() {
		pdf.Image("dst.jpg", 0, 0, 200, 200, false, "", 0, "")
		pdf.SetY(5)
	})

	fileStr := "./pdf/Fpdf_AddPage.pdf"
	err := pdf.OutputFileAndClose(fileStr)
	Summary(err, fileStr)
	// Output:
	// Successfully generated pdf/Fpdf_AddPage.pdf
}

func Filename(baseStr string) string {
	return PdfFile("./")
}

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

//asdfasdf
func PdfDir() string {
	return filepath.Join("./", "pdf")
}

// PdfFile returns a qualified filename in which the path to the PDF output
// directory is prepended to the specified filename.
func PdfFile(fileStr string) string {
	return filepath.Join(PdfDir(), fileStr)
}

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
