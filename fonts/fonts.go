package fonts

import (
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

var Basic *text.Atlas

func init() {
	Basic = text.NewAtlas(basicfont.Face7x13, text.ASCII)
}
