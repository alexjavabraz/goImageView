package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	tiff "github.com/chai2010/tiff"
)

var files = []string{
	"./000001.tif",
}

func main() {
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

		// Encode tiff
		for i := 0; i < p.ImageNum(); i++ {
			for j := 0; j < p.SubImageNum(i); j++ {

				// is tiled?
				_, isTiled := p.Ifd[i][j].TagGetter().GetTileWidth()
				fmt.Printf("%s(%02d,%02d) isTiled: %v\n", name, i, j, isTiled)

				blocksAcross := p.ImageBlocksAcross(i, j)
				blocksDown := p.ImageBlocksDown(i, j)

				fmt.Printf("BlockAcross %d, BlocksDown %d", blocksAcross, blocksDown)

				for col := 0; col < blocksAcross; col++ {
					for row := 0; row < blocksDown; row++ {
						newname := fmt.Sprintf("%s-%02d-%02d-%02d-%02d.tiff", filepath.Base(name), i, j, col, row)
						fmt.Printf("New name %s", newname)

						m, err := p.DecodeImageBlock(i, j, col, row)
						if err != nil {
							log.Fatal(err)
							return
						}

						var buf bytes.Buffer
						if err = tiff.Encode(&buf, m, nil); err != nil {
							log.Fatal(err)
						}
						if err = ioutil.WriteFile(newname, buf.Bytes(), 0666); err != nil {
							log.Fatal(err)
						}
						fmt.Printf("Save %s ok\n", newname)
					}
				}
			}
		}
	}
}
