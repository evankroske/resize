package main

import (
	"math/big"
	"fmt"
	"os"
	"image/png"
	"image"
	"log"
	"github.com/evankroske/resize/resize"
)

func main() {
	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	scaleFactor := new(big.Rat)
	_, err = fmt.Sscan(os.Args[1], scaleFactor)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("scaleFactor: %v\n", scaleFactor)
	}
	resizedImage := resize.Resize(img, scaleFactor)
	log.Println(resizedImage.Bounds())
	png.Encode(os.Stdout, resizedImage)
}
