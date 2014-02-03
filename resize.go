/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
