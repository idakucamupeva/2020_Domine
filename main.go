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

type mouseState struct{
	leftButton bool
	rightButton bool
	x,y	int
}

func getMouseState() mouseState{
	mouseX, mouseY, mouseButtonState := sdl.GetMouseState()
	leftButton := mouseButtonState & sdl.ButtonLMask()
	rightButton := mouseButtonState & sdl.ButtonRMask()
	var result mouseState
	result.x = int(mouseX)
	result.y = int(mouseY)
	result.leftButton = !(leftButton==0) //left button is pressed
	result.rightButton = !(rightButton==0) //right button is pressed

	return result
}

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
			dominoTmp.texHidden = textureFromBMP(renderer, "BMPdominoes/7-7.bmp")
			dominoesMap[counter] = dominoTmp
			counter++
			startStr = "BMPdominoes/"
		}
	}


	initDomino()
	initComputerDomino()

	currentMouseState := getMouseState()
	previousMouseState := currentMouseState
	for{
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent(){
			switch event.(type) {
			case *sdl.QuitEvent:
				return
				
			}
		}
		currentMouseState = getMouseState()

		if !previousMouseState.leftButton && currentMouseState.leftButton {
			fmt.Println("Left click")
		}

/*
			mouseX := currentMouseState.x
			mouseY := currentMouseState.y
			fmt.Println(mouseX,mouseY)


		for i:=0; i< len(player1.deck); i++{
			x := player1.deck[i].x*0.7
			y := player1.deck[i].y*0.7

			if float64(mouseX) >= x && float64(mouseX)<=(x+189) && float64(mouseY)<=y && float64(mouseY)>=(y+90){
				fmt.Println("Domino hit")
			}
		}

*/
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.SetDrawColor(128, 0, 0, 0)
		renderer.FillRect(&sdl.Rect{50, 550, 500, 200})
		renderer.FillRect(&sdl.Rect{50, 50, 500, 200})

			
			for i:=0; i< len(player1.deck); i++{
				printDomino(&player1.deck[i])
			}


		renderer.SetScale(0.7, 0.7)
		
		for _, dom := range player1.deck{
			dom.draw(renderer)
		}	

		for _, dom := range player2.deck{
			dom.drawHiddenDomino(renderer)
		}	

		renderer.SetScale(1, 1)
		
		
		
		renderer.Present()

		previousMouseState = currentMouseState
	}
}
