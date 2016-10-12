package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
)

func main() {
	r := rand.New(rand.NewSource(99))

	ImageFileDest := bytes.NewBufferString("")
	fmt.Fprint(ImageFileDest, "./temp/copiedLocal", r.Float32(), ".tif")

	ImageFile := "z://000004.tif"

	sFile, err := os.Open(ImageFile)
	if err != nil {
		log.Fatal(err)
	}
	defer sFile.Close()

	eFile, err := os.Create(ImageFileDest.String())
	if err != nil {
		log.Fatal(err)
	}
	defer eFile.Close()

	_, err = io.Copy(eFile, sFile) // first var shows number of bytes
	if err != nil {
		log.Fatal(err)
	}

	err = eFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

	delete(ImageFileDest.String())

}

func delete(ImageName string) {
	err := os.Remove(ImageName)

	if err != nil {
		fmt.Println(err)
	}
}
