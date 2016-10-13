package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	ImageFile := "g://000001.tif"
	ImageFileDest := criar(ImageFile)

	delete(ImageFileDest)

}

func criar(ImageFile string) (ImageName string) {
	fmt.Printf("Criando")

	r := rand.New(rand.NewSource(99))

	ImageFileDest := bytes.NewBufferString("")
	fmt.Fprint(ImageFileDest, "./temp/copiedLocal", r.Float32(), ".tif")

	fmt.Printf("Buscando %s \n", ImageFile)

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

	return ImageFileDest.String()
}

func delete(ImageName string) {
	fmt.Printf("Removendo %s \n", ImageName)
	err := os.Remove(ImageName)

	if err != nil {
		fmt.Println(err)
	}
}
