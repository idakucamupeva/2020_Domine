package main


import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

type domino struct{
	left, right int
	tex *sdl.Texture
	x, y float64
}

const (
	dominoWidth = 189
	dominoHeight = 90
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

func newDomino(renderer *sdl.Renderer, filename string, left, right int) (dom domino){
	dom.tex = textureFromBMP(renderer, filename)

	dom.x = screenWidth/ 2.0
	dom.y = screenHeight - dominoHeight / 2.0 

		return dom
}

func (dom *domino) draw(renderer *sdl.Renderer){
	x := dom.x - dominoWidth
	y := dom.y - dominoWidth*2
	renderer.Copy(dom.tex, &sdl.Rect{0, 0, dominoWidth, dominoHeight},
		&sdl.Rect{int32(x), int32(y), dominoWidth, dominoHeight})

}