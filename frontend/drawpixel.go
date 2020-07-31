package frontend

import "github.com/veandco/go-sdl2/sdl"

func (draw *draw) drawPixel(x, y int, color uint32) {
	point := sdl.Rect{int32(x), int32(y), draw.size, draw.size}
	draw.FillRect(&point, color)
}
