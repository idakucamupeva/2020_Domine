package main


import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
	"math/rand"
	"time"
)

type domino struct{
	left, right int
	assigned int
	tex *sdl.Texture
	x, y float64
}

const (
	tablePositionWidth = 150
	tablePositionHeight = 555
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

func newDomino(renderer *sdl.Renderer, filename string, left, right int) (dom* domino){
	dom.tex = textureFromBMP(renderer, filename)

	dom.left = left
	dom.right = right

	dom.x = tablePositionWidth
	dom.y = tablePositionHeight

	dom.assigned = 111

		return dom
}

func (dom *domino)draw(renderer *sdl.Renderer){
	x := dom.x 
	y := dom.y 
	renderer.CopyEx(dom.tex, &sdl.Rect{0, 0, dominoWidth, dominoHeight},
		&sdl.Rect{int32(x), int32(y), dominoWidth, dominoHeight}, 90, 
		&sdl.Point{0, 0},
		sdl.FLIP_NONE)

}

func initDomino(dom *domino, renderer *sdl.Renderer){
	

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	
	for i:=0; i<7; i++ {
		randNum := r1.Intn(28)
		dominoesMap[randNum].assigned = 1

	}

	
}
