package main


import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	//"math/rand"
	//"time"
	"strconv"
)

const (
	screenWidth = 600
	screenHeight = 800
)



type Player struct{
	name string
	deck []domino 
}


//player constructor
func newPlayer(name string, deck []domino) *Player{
	return &Player{
		name: name,
		deck: deck,
	}
}

var dominoesMap = make(map[int]domino, 28)


func main(){

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil{
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(  
		"MATF Dominoes",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}

	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}

	defer renderer.Destroy()

	//domino's number
	var counter int
	counter = 0	

	
	startStr := "BMPdominoes/"
	
	for i:=0; i < 7; i++{
		for j:=i; j < 7; j++{
			startStr += strconv.Itoa(i) + "-" + strconv.Itoa(j) + ".bmp"
			dominoTmp := newDomino(renderer, startStr, i, j)
			dominoesMap[counter] = dominoTmp
			//fmt.Println("heeeej")
			counter++
			startStr = "BMPdominoes/"
		}
	}

	//fmt.Println(&dominoesMap[21].assigned)
	//dominoesMap[21].assigned = 222
	//fmt.Println(dominoesMap[21].assigned)
	var flag = 0
	for{
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent(){
			//fmt.Print(flag)
			switch event.(type) {
			case *sdl.QuitEvent:
				return
				
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.SetDrawColor(128, 0, 0, 0)
		renderer.FillRect(&sdl.Rect{50, 550, 500, 200})
		if flag==0 {
			initDomino(renderer)
			flag = 1
			renderer.Present()
		}
	}

}