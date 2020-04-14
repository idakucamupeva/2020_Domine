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
	texHidden *sdl.Texture	
	x, y float64
}
//ispraviti tablePositions
const (
	tablePositionWidth = 190
	//tablePositionHeight = 830
	tablePositionHeight = 800
	tablePositionWidthOpponent = 190
	tablePositionHeightOpponent = 120
	dominoWidth = 189
	dominoHeight = 90
)
/*

var tablePositionWidth = float64(50)
var tablePositionHeight = float64(height-(height/4+50))
var tablePositionWidthOpponent =  float64(50)
var tablePositionHeightOpponent = float64(50)
*/


//Making texture from bmp picture
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

	dom.assigned = -1

		return dom
}

func (dom *domino)draw(renderer *sdl.Renderer, angle float64, a, b int32){
	x := dom.x 
	y := dom.y 
	renderer.CopyEx(dom.tex, &sdl.Rect{0, 0, dominoWidth, dominoHeight},
		&sdl.Rect{int32(x), int32(y), dominoWidth, dominoHeight}, angle, 
		&sdl.Point{a, b},
		sdl.FLIP_NONE)

}

func (dom *domino)drawHiddenDomino(renderer *sdl.Renderer){
	x := dom.x 
	y := dom.y 
	renderer.CopyEx(dom.texHidden, &sdl.Rect{0, 0, dominoWidth, dominoHeight},
		&sdl.Rect{int32(x), int32(y), dominoWidth, dominoHeight}, 90, 
		&sdl.Point{0, 0},
		sdl.FLIP_NONE)

}



func initDomino(){
	
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	
	for i:=0; i<7; i++ {
		randNum := r1.Intn(28)
		if dominoesMap[randNum].assigned == -1{
			dominoTmp := dominoesMap[randNum]
			dominoTmp.x = (float64)(tablePositionWidth + i*dominoWidth/2)
			dominoTmp.y = tablePositionHeight
			dominoTmp.assigned = 1
			dominoesMap[randNum] = dominoTmp
			

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
		if dominoesMap[randNum].assigned == -1{
			dominoTmp := dominoesMap[randNum]
			dominoTmp.x = (float64)(tablePositionWidthOpponent + i*dominoWidth/2)
			dominoTmp.y = tablePositionHeightOpponent
			dominoTmp.assigned = 2
			dominoesMap[randNum] = dominoTmp

			player2.deck = append(player2.deck, dominoTmp)
					
		}else{
			i--
		}
		
	}
}

func printDomino(dom *domino){
	fmt.Println(dom.left," ,",dom.right)
}
