package main


import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)


//Button struct describes all buttons used in game
type Button struct{
	tex *sdl.Texture
	x, y float64
}

const(
	leftAndRightSize = 50
	bankSize = 200
)



//button constructor
func newButton(renderer *sdl.Renderer, filename string, x, y float64) (but Button){
	but.tex = textureFromBMP(renderer, filename)
	but.x = x
	but.y = y

	return but
}

func (but *Button)drawButton(renderer *sdl.Renderer, size int){
	x := but.x 
	y := but.y 
	err := renderer.Copy(but.tex, &sdl.Rect{W: int32(size), H: int32(size)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(size), H: int32(size)})
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}