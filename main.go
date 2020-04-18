package main


import (
	//"flag"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
	//"time"
)

//Mouse events related structures and functions
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
	result.leftButton = !(leftButton==0) //left mouse button is pressed
	result.rightButton = !(rightButton==0) //right mouse button is pressed
	return result
}


//Global variables
var leftButtonClicked = 0
var dominoesMap = make(map[int]domino, 28)
var player1 = newPlayer("player1", true, nil)
var player2 = newPlayer("bot", false, nil)
var player1Active = true
var dominoCounter = 0
var dominoCounterOpponent = 0
//var width =int32(1000)
var width =int32(1200)
var height =int32(700)
var hasADominoFromBank [] bool
var hasComputerDominoFromBank [] bool
var scaleSize = 0.7
var leftDominoCounter =0
var rightDominoCounter =0


func main(){
	//Initialization
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil{
		fmt.Println("initializing SDL:", err)
		return
	}

	var window *sdl.Window
	window, err = sdl.CreateWindow(
		"MATF Dominoes",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height,
		sdl.WINDOW_OPENGL |  sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE) // | sdl.WINDOW_FULLSCREEN_DESKTOP)
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
	var counter = 0
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

	//BUTTONS-textures
	leftBtn := newButton(renderer, "img/leftBtn.bmp", float64(width)*0.85, float64(height)*0.9)
	rightBtn := newButton(renderer, "img/rightBtn.bmp", float64(width)*0.85+leftAndRightSize+20, float64(height)*0.9)
	//bankBtn := newButton(renderer, "img/bank.bmp", (float64(width)/6*5)/0.7, (float64(height)/2-float64(bankSize)/2)/0.7)
	bankBtn := newButton(renderer, "img/bank.bmp", (float64(width)/6*5)/0.7, float64(height)/7)

	//TODO  bank is empty
	//emptyBankBtn := newButton(renderer, "img/bankEmpty.bmp", (float64(width)/6*5)/0.7, (float64(height)/2-float64(bankSize)/2)/0.7)

	//final scene
	player1WonScene := newFinalScene(renderer, "img/youWon.bmp", 200, 150)
	trophyScene := newFinalScene(renderer, "img/trophy.bmp", 500, -150)

	
	initDomino()
	initComputerDomino()
	table := newGameTable()

	for i:=0; i< 21; i++{
		hasADominoFromBank = append(hasADominoFromBank, false)
	}

	currentMouseState := getMouseState()
	previousMouseState := currentMouseState

	var tmpX = float64(width)/6*5 //0.7
	var tmpY = (float64(height)/7)*0.7
	var leftBtnTmpX = float64(width)*0.85
	var leftBtnTmpY = float64(height)*0.9
	var rightBtnTmpX = float64(width)*0.85+leftAndRightSize+20
	var rightBtnTmpY = float64(height)*0.9
	//previousMouseState := currentMouseState


	var player1Won = false  //player1 won signal
	var player2Won = false	//player2 won signal

//	var changeCoordinate bool =true

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					return
				}
			}

			currentMouseState = getMouseState()

			if player1Active {
				switch event.(type) {
				case *sdl.MouseButtonEvent:

					mouseX := currentMouseState.x
					mouseY := currentMouseState.y

	//				fmt.Println(mouseX,mouseY)

					//bank clicked
					if !currentMouseState.leftButton && previousMouseState.leftButton{
						if float64(mouseX) >= tmpX && float64(mouseX) <= tmpX+140 && float64(mouseY) >= tmpY && float64(mouseY) <= tmpY+140{
							fmt.Println("Bank touched")
							addFromBank()
						}
					}
					//left arrow on screen clicked
					if !currentMouseState.leftButton && previousMouseState.leftButton{
						if float64(mouseX) >= leftBtnTmpX && float64(mouseX) <= leftBtnTmpX+leftAndRightSize && float64(mouseY) >= leftBtnTmpY && float64(mouseY) <= leftBtnTmpY+leftAndRightSize{
							fmt.Println("levi button touched")
							leftButtonClicked = 1
						}
					}
					//right arrow on screen clicked
					if !currentMouseState.leftButton && previousMouseState.leftButton{
						if float64(mouseX) >= rightBtnTmpX && float64(mouseX) <= rightBtnTmpX+leftAndRightSize && float64(mouseY) >= rightBtnTmpY && float64(mouseY) <= rightBtnTmpY+leftAndRightSize{
							fmt.Println("desni button touched")
							leftButtonClicked = 0
						}
					}

					for i := 0; i < len(player1.deck); i++ {
						x := player1.deck[i].x
						y := player1.deck[i].y

						if float64(mouseX) >= (x-dominoHeight)*scaleSize && float64(mouseX) <= x*scaleSize && float64(mouseY) <= (y+dominoWidth)*scaleSize && float64(mouseY) >= y*scaleSize {
							//fmt.Println("Domino hit", i)
							if player1.deck[i].assigned==1{
								if play(&player1, i, &table) {
									if isWon(&player1) {
										fmt.Println("Player1 won")
										player1Won = true

									}

									player1Active = !player1Active //false

									if computerPlay(&player2, &table) {
										if isWon(&player2) {
											fmt.Println("Player2 won")
											player2Won = true
										}
										player1Active = !player1Active //true
									} else {
										computerPlay(&player2, &table)
										player1Active = !player1Active //true

									}

									/*else{
										if computerPlay(&player2,&table){
											if isWon(&player2){
												fmt.Println("Player2 won")
											}
											player1Active = !player1Active	//true

										}
									}*/
								}
							}
								//player1.deck = append(player1.deck[:i], player1.deck[i+1:]...)
						}
					}

				}	//switch

			}	//if active

			renderer.SetDrawColor(255, 255, 255, 255) 	//background
			renderer.Clear()

			if player1Won{
					player1WonScene.drawScene(renderer, plr1WonWidth, plr1WonHeight)
					trophyScene.drawScene(renderer, trophyWidth, trophyHeight)	
					trophyScene.updateScene(5, 550, 220)
				
			}else if player2Won{

			}else{

				renderer.SetDrawColor(128, 0, 0, 0)			//rectangles
				//renderer.FillRect(&sdl.Rect{X: width/16, Y: height-(height/4+50), W: width/4*3, H: height/4})
				renderer.FillRect(&sdl.Rect{X: width/16, Y: height-(height/4), W: width/4*3, H: height/4})
				renderer.FillRect(&sdl.Rect{X: width/16, Y: height/16, W: width/4*3, H: height/4}) //W: width-100
				//renderer.FillRect(&sdl.Rect{50, 550, width/4*3, height/4})

				leftBtn.drawButton(renderer, leftAndRightSize)
				rightBtn.drawButton(renderer, leftAndRightSize)

				renderer.SetScale(0.7, 0.7)
				//renderer.SetScale(float32(scaleSize), float32(scaleSize))

				bankBtn.drawButton(renderer, bankSize)

				for _, dom := range player1.deck {
					if dom.assigned == 1 {
						dom.draw(renderer, 90.0, 0, 0)
					}
					if dom.assigned == 0 { //TODO rotation
						if leftDominoCounter>6 || rightDominoCounter>5{
							dom.y = dom.y / 0.7
							//dom.y = dom.y / 0.8
							renderer.SetScale(0.4, 0.4)
						}else{
							renderer.SetScale(0.5, 0.5)
						}
						dom.draw(renderer, dom.rotation, dominoWidth/2, dominoHeight/2)
						renderer.SetScale(0.7, 0.7)
					}
				}

				for _, dom := range player2.deck {
					if dom.assigned == 2{
							//dom.drawHiddenDomino(renderer)
						dom.draw(renderer,90, 0, 0)
					}
					if dom.assigned == 0 { //TODO rotation
						if leftDominoCounter>6 || rightDominoCounter>5 {
							//dom.y = dom.y / 0.8
							dom.y = dom.y / 0.7
							renderer.SetScale(0.4, 0.4)
						}else{
							renderer.SetScale(0.5, 0.5)
						}
						dom.draw(renderer, dom.rotation, dominoWidth/2, dominoHeight/2)
						renderer.SetScale(0.7, 0.7)
					}
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

			previousMouseState = currentMouseState
		}
	}
}
