package main


import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

const (
	screenWidth = 600
	screenHeight = 800
)

var dominoesMap = make(map[int]domino, 28)
var player1 = newPlayer("player1", true, nil)
var player2 = newPlayer("bot", false, nil)



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


	flag := 0
	flag1 := 0
	for{
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent(){
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
			//moj kod
			for i:=0; i< len(player1.deck); i++{
				printDomino(&player1.deck[i])
			}

			//fmt.Println(player1.deck)
		}
		if flag1==0 {
			initComputerDomino()
			flag1 = 1
			//moj kod
			for i:=0; i< len(player2.deck); i++{
				printDomino(&player2.deck[i])
			}
			//fmt.Println(player2.deck)
			renderer.Present()
		}

		
	}

}