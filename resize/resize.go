package resize

import (
	"image"
	"image/color"
	"math/big"
)


type resizedImage struct {
	img image.Image
	scaleFactor big.Rat
}

func (p *resizedImage) Bounds() image.Rectangle {
	old_rect := p.img.Bounds()
	return image.Rect(0, 0, scaleBy(old_rect.Dx(), p.scaleFactor), scaleBy(old_rect.Dy(), p.scaleFactor))
}


func (p *resizedImage) ColorModel() color.Model {
	return p.img.ColorModel()
}


func (p *resizedImage) At(x, y int) color.Color {
	inverseScaleFactor := new(big.Rat)
	inverseScaleFactor.Set(&p.scaleFactor)
	inverseScaleFactor.Inv(inverseScaleFactor)
	return p.img.At(scaleBy(x, *inverseScaleFactor), scaleBy(y, *inverseScaleFactor))
}

func scaleBy(a int, scaleFactor big.Rat) int {
	scaled := new(big.Rat)
	scaled.Mul(big.NewRat(int64(a), 1), &scaleFactor)
	return int(scaled.Num().Int64() / scaled.Denom().Int64())
}

func Resize(p image.Image, scaleFactor *big.Rat) image.Image {
	return &resizedImage{p, *scaleFactor}
}
