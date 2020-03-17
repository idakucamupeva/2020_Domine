package main


import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

type dominoe struct{
	tex *sdl.Texture
	x, y float64
}

const (
	dominoeWidth = 189
	dominoeHeight = 90
)

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture{
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

func newDominoe(renderer *sdl.Renderer, filename string) (dom dominoe){
	dom.tex = textureFromBMP(renderer, filename)

	dom.x = screenWidth/ 2.0
	dom.y = screenHeight - dominoeHeight / 2.0 

		return dom
}

func (dom *dominoe) draw(renderer *sdl.Renderer){
	x := dom.x - dominoeWidth
	y := dom.y - dominoeWidth*2
	renderer.Copy(dom.tex, &sdl.Rect{0, 0, dominoeWidth, dominoeHeight},
		&sdl.Rect{int32(x), int32(y), dominoeWidth, dominoeHeight})

}