package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func main() {
	fp, err := ioutil.TempFile("", "ScreenQRLoader-*.jpg")
	if err != nil {
		log.Println("Temporary file could not be generated")
		os.Exit(1)
	}

	fpath := fp.Name()
	err = exec.Command("import", "-monochrome", fpath).Run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// open and decode image file
	file, err := os.Open(fpath)
	if err != nil {
		log.Panicln(err)
		os.Exit(1)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}
