package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello QR Code")

	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = GenerateQRCode(file, "555-2368")
	if err != nil {
		log.Fatal(err)
	}
}

//GenerateQRCode receive a string and return a slice of strings
func GenerateQRCode(w io.Writer, code string) error {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	return png.Encode(w, img)
}
