package fonts

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

var Basic *text.Atlas

func init() {
	Basic = text.NewAtlas(basicfont.Face7x13, text.ASCII)
}

func NewText(content string, pos pixel.Vec) *text.Text {
	txt := text.New(pos, Basic)
	fmt.Fprintln(txt, content)
	return txt
}
