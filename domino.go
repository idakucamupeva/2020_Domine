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
	filename string
	tex *sdl.Texture
	x, y float64
}

const (
	tablePositionWidth = 200
	tablePositionHeight = 830
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
	dom.filename = filename

	dom.left = left
	dom.right = right

	dom.x = tablePositionWidth
	dom.y = tablePositionHeight

	dom.assigned = 0

		return dom
}

func (dom *domino)draw(renderer *sdl.Renderer){
	x := dom.x 
	y := dom.y 
	renderer.SetScale(0.7, 0.7)
	renderer.CopyEx(dom.tex, &sdl.Rect{0, 0, dominoWidth, dominoHeight},
		&sdl.Rect{int32(x), int32(y), dominoWidth, dominoHeight}, 90, 
		&sdl.Point{0, 0},
		sdl.FLIP_NONE)

}

func initDomino(renderer *sdl.Renderer){
	
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	
	for i:=0; i<7; i++ {
		randNum := r1.Intn(28)
		if dominoesMap[randNum].assigned == 0{
			dominoTmp := dominoesMap[randNum]
			dominoTmp.x = (float64)(tablePositionWidth + i*dominoWidth/2)
			dominoTmp.y = tablePositionHeight
			dominoTmp.assigned = 1
			dominoesMap[randNum] = dominoTmp
			dominoTmp.draw(renderer)

			player1.deck = append(player1.deck, dominoTmp)
		
		}else{
			i--
		}
		
	}
	
}

func initComputerDomino(){
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	
	for i:=0; i<7; i++ {
		randNum := r1.Intn(28)
		if dominoesMap[randNum].assigned == 0{
			dominoTmp := dominoesMap[randNum]
			dominoTmp.assigned = 2
			dominoesMap[randNum] = dominoTmp

			player2.deck = append(player2.deck, dominoTmp)
					
		}else{
			i--
		}
		
	}
}
