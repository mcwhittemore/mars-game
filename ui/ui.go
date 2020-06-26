package ui

import (
	"app/characters"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type uiCacheItem struct {
	Canvas      *pixelgl.Canvas
	Pos         pixel.Vec
	Etag        string
	Generator   func(characters.MindInput) (*pixelgl.Canvas, pixel.Vec)
	ComputeETag func(characters.MindInput) string
}

var uiCacheStore []*uiCacheItem

type UIElement int

const (
	ItemsUI UIElement = iota
)

func Draw(id UIElement, mi characters.MindInput) {
	ele := uiCacheStore[id]

	etag := ele.ComputeETag(mi)

	if etag != ele.Etag {
		ele.Canvas, ele.Pos = ele.Generator(mi)
	}

	mi.AddCanvasStatic(ele.Canvas, ele.Pos)
}

func init() {
	uiCacheStore = make([]*uiCacheItem, 1)

	uiCacheStore[ItemsUI] = &uiCacheItem{
		Generator:   drawHeroItemsUI,
		ComputeETag: eTagHeroItemUI,
	}
}
