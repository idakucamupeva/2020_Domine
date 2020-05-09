package main


import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Button struct{
	tex *sdl.Texture
	x, y float64
}

const(
	leftAndRightSize = 50
	bankSize = 200
)

func textureFromBMPBtn(renderer *sdl.Renderer, filename string) *sdl.Texture{
	img, err := sdl.LoadBMP(filename)
		if err != nil {
			panic(fmt.Errorf("loading %v: %v", filename, err))
		}

	defer img.Free()
	
	tex, err := renderer.CreateTextureFromSurface(img)
		if err != nil{
			panic(fmt.Errorf("creating texture from %v: %v", filename, err))
		}

	return tex
}

//button constructor
func newButton(renderer *sdl.Renderer, filename string, x, y float64) (but Button){
	but.tex = textureFromBMPBtn(renderer, filename)
	but.x = x
	but.y = y

	return but
}

func (but *Button)drawButton(renderer *sdl.Renderer, size int){
	x := but.x 
	y := but.y 
	renderer.Copy(but.tex, &sdl.Rect{W: int32(size), H: int32(size)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(size), H: int32(size)})
}