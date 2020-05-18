package main


import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
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
	rotation float64
}

const (
	dominoWidth = 189
	dominoHeight = 90
)

var tablePositionX = float64(width/16)/0.7
var tablePositionY = float64(height-(height/4))/0.7
var tablePositionXOpponent =  float64(width/16)/0.7
var tablePositionYOpponent = float64(height/16)/0.7

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

	dom.x = tablePositionX
	dom.y = tablePositionY

	dom.assigned = -1
	dom.rotation = 0

	return dom
}

func (dom *domino)draw(renderer *sdl.Renderer, angle float64, a, b int32){
	x := dom.x 
	y := dom.y
	err := renderer.CopyEx(dom.tex, &sdl.Rect{W: dominoWidth, H: dominoHeight},
		&sdl.Rect{X: int32(x), Y: int32(y), W: dominoWidth, H: dominoHeight}, angle,
		&sdl.Point{X: a, Y: b},
		sdl.FLIP_NONE)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}

func (dom *domino)drawHiddenDomino(renderer *sdl.Renderer){
	x := dom.x 
	y := dom.y 
	err := renderer.CopyEx(dom.texHidden, &sdl.Rect{W: dominoWidth, H: dominoHeight},
		&sdl.Rect{X: int32(x), Y: int32(y), W: dominoWidth, H: dominoHeight}, 90,
		&sdl.Point{},
		sdl.FLIP_NONE)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}

func initDomino(){
	
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	
	for i:=1; i<8; i++ {
		randNum := r1.Intn(28)
		if dominoesMap[randNum].assigned == -1{
			dominoTmp := dominoesMap[randNum]
			dominoTmp.x = tablePositionX + float64(i)*dominoWidth/2
			dominoTmp.y = tablePositionY
			dominoTmp.assigned = 1
			dominoesMap[randNum] = dominoTmp
			player1.deck = append(player1.deck, dominoTmp)
		}else {
			i--
		}
	}
	dominoCounter = 8
}

func initComputerDomino(){
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	
	for i:=1; i<8; i++ {
		randNum := r1.Intn(28)
		if dominoesMap[randNum].assigned == -1{
			dominoTmp := dominoesMap[randNum]
			dominoTmp.x = tablePositionXOpponent + (float64)(i*dominoWidth/2)
			dominoTmp.y = tablePositionYOpponent
			dominoTmp.assigned = 2
			dominoesMap[randNum] = dominoTmp
			player2.deck = append(player2.deck, dominoTmp)
		}else{
			i--
		}
	}
	dominoCounterOpponent = 8
}


//function returns who plays first
func firstMove() int{

	max1 := -1
	max2 := -1

	for _, dom := range player1.deck {
		if dom.left == dom.right{
			if dom.left > max1{
				max1 = dom.left
			}
		}

	}


	for _, dom := range player2.deck {
		if dom.left == dom.right{
			if dom.left > max2{
				max2 = dom.left
			}
		}
	}

	if max1 != max2 {
		if max1 > max2{
			return 1
		}else{
			return 2
		}
	}

	max1 = -1
	max2 = -1


	for _, dom := range player1.deck {
		if dom.right == 6{
			if dom.left > max1{
				max1 = dom.left
			}
		}

	}


	for _, dom := range player2.deck {
		if dom.right == 6{
			if dom.left > max2{
				max2 = dom.left
			}
		}
	}

	if max1 != max2 {
		if max1 > max2{
			return 1
		}else{
			return 2
		}
	}

	max1 = -1
	max2 = -1


	for _, dom := range player1.deck {
		if dom.right == 5{
			if dom.left > max1{
				max1 = dom.left
			}
		}

	}


	for _, dom := range player2.deck {
		if dom.right == 5{
			if dom.left > max2{
				max2 = dom.left
			}
		}
	}

	
	if max1 < max2{
		return 2
	}
	
	return 1

}
