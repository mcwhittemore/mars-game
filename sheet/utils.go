package sheet

import (
	"app/data"

	"image"
	_ "image/png"

	"github.com/faiface/pixel"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := data.Open("/" + path)
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
