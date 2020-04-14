package main


import (
//	"flag"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
	//"time"
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

var player1_active = true

var width =int32(1000)
var height =int32(700)

func main(){

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil{
	//if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil{
		fmt.Println("initializing SDL:", err)
		return
	}

	var window *sdl.Window


	window, err = sdl.CreateWindow(
		//	window, err := sdl.CreateWindow(
		"MATF Dominoes",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height,
		sdl.WINDOW_OPENGL |  sdl.WINDOW_SHOWN )//sdl.WINDOW_RESIZABLE | sdl.WINDOW_FULLSCREEN_DESKTOP)
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
	table := newGameTable()


	currentMouseState := getMouseState()
	//previousMouseState := currentMouseState
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE{
					return
				}
			}
			currentMouseState = getMouseState()

			if player1_active {
				switch event.(type) {
				case *sdl.MouseButtonEvent:

					mouseX := currentMouseState.x
					mouseY := currentMouseState.y

//					fmt.Println(mouseX,mouseY)


					for i := 0; i < len(player1.deck); i++ {
						x := player1.deck[i].x
						y := player1.deck[i].y

						if float64(mouseX) >= (x-dominoHeight)*0.7 && float64(mouseX) <= x*0.7 && float64(mouseY) <= (y+dominoWidth)*0.7 && float64(mouseY) >= y*0.7 {
							//fmt.Println("Domino hit", i)
							if player1.deck[i].assigned==1{
							if play(&player1, i, &table){
									if isWon(&player1){
										fmt.Println("Player1 won")
									}

									player1_active = !player1_active

									if computerPlay(&player2,&table){
										if isWon(&player2){
											fmt.Println("Player2 won")
										}
										player1_active = !player1_active

									}
																}
							}
								//player1.deck = append(player1.deck[:i], player1.deck[i+1:]...)
						}
					}

				}

			}

			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.Clear()

			renderer.SetDrawColor(128, 0, 0, 0)
			//renderer.FillRect(&sdl.Rect{50, 550, width/4*3, height/4})
			renderer.FillRect(&sdl.Rect{width/16, height-(height/4+50), width/4*3, height/4})
			renderer.FillRect(&sdl.Rect{width/16, height/16, width/4*3, height/4}) //W: width-100

			renderer.SetDrawColor(50, 0, 128, 0)

			//renderer.FillRect(&sdl.Rect{width-200, height/4+100, 150, 200})
			//TODO iscrtati dominu u banci

			renderer.SetScale(0.7, 0.7)

			for _, dom := range player1.deck {
				if dom.assigned == 1 {
					dom.draw(renderer, 90.0, 0, 0)
				}
				if dom.assigned == 0 { //TODO rotation
					renderer.SetScale(0.5, 0.5)
					dom.draw(renderer, dom.rotation, dominoWidth/2, dominoHeight/2)
					renderer.SetScale(0.7, 0.7)

				}
			}
			for _, dom := range player2.deck {
				if dom.assigned == 2{
				//	dom.drawHiddenDomino(renderer)
				dom.draw(renderer,90, 0, 0)
				}
				if dom.assigned == 0 { //TODO rotation
					renderer.SetScale(0.5, 0.5)
					dom.draw(renderer, dom.rotation, dominoWidth/2, dominoHeight/2)
					renderer.SetScale(0.7, 0.7)

				}

			}
			/*
				for _, dom := range player2.deck {
					if dom.assigned == 2{
					dom.drawHiddenDomino(renderer)
					}
				}
*/
			renderer.SetScale(1, 1)

			renderer.Present()

			//previousMouseState = currentMouseState
		}
	}
}
