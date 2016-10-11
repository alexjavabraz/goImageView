package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"time"

	tiff "github.com/chai2010/tiff"
)

func main() {
	t0 := time.Now()

	ImageFile := "./000005.tif"
	dataofimage := ImageRead2(ImageFile)
	Formatpng2(dataofimage)
	fmt.Println("Converting  image to gif  ...")

	t1 := time.Now()
	fmt.Printf("The call toook %v to run. \n", t1.Sub(t0))
	fmt.Printf("Go launched at %s\n", t0.Local())
	fmt.Printf("Go terminated at %s\n", t1.Local())
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
func Formatpng2(img image.Image) {
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

	fmt.Println("\n success... \n ")

}
