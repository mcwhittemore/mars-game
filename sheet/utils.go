package sheet

import (
	"image"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/markbates/pkger"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := pkger.Open("/" + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
