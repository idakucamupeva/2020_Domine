package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
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
var width =int32(1200)
var height =int32(700)
var hasADominoFromBank [] bool
var hasComputerDominoFromBank [] bool
var scaleSize = 0.7
var leftDominoCounter =0
var rightDominoCounter =0
var oneDominoOnly = true
var player1Won = false  //player1 won signal
var player2Won = false	//player2 won signal


func start(){
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
		sdl.WINDOW_OPENGL |  sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}

	//domino's number
	var counter = 0
	startStr := "../BMPdominoes/"
	for i:=0; i < 7; i++{
		for j:=i; j < 7; j++{
			startStr += strconv.Itoa(i) + "-" + strconv.Itoa(j) + ".bmp"
			dominoTmp := newDomino(renderer, startStr, i, j)
			dominoTmp.texHidden = textureFromBMP(renderer, "../BMPdominoes/7-7.bmp")
			dominoesMap[counter] = dominoTmp
			counter++
			startStr = "../BMPdominoes/"
		}
	}

	//BUTTONS-textures
	leftBtn := newButton(renderer, "../img/leftBtn.bmp", float64(width)*0.85, float64(height)*0.9)
	rightBtn := newButton(renderer, "../img/rightBtn.bmp", float64(width)*0.85+leftAndRightSize+20, float64(height)*0.9)
	nextBtn := newButton(renderer, "../img/next.bmp", float64(width)*0.85+leftAndRightSize-17, 550)
	bankBtn := newButton(renderer, "../img/bank.bmp", (float64(width)/6*5)/0.7, float64(height)/7)

	//final scene
	player1WonScene1 := newFinalScene(renderer, "../img/youWon1.bmp", 200, 50)
	player1WonScene2 := newFinalScene(renderer, "../img/youWon2.bmp", 200, 50)
	
	player2WonScene1 := newFinalScene(renderer, "../img/youLost1.bmp", 200, 50)
	player2WonScene2 := newFinalScene(renderer, "../img/youLost2.bmp", 200, 50)
	


	initDomino()
	initComputerDomino()
	table := newGameTable()

	for i:=0; i< 7; i++{
		hasADominoFromBank = append(hasADominoFromBank, true)
	}

	for i:=7; i< 21; i++{
		hasADominoFromBank = append(hasADominoFromBank, false)
	}


	for i:=0; i< 7; i++{
		hasComputerDominoFromBank = append(hasComputerDominoFromBank, true)
	}

	for i:=7; i< 21; i++{
		hasComputerDominoFromBank = append(hasComputerDominoFromBank, false)
	}


	currentMouseState := getMouseState()
	previousMouseState := currentMouseState

	var tmpX = float64(width)/6*5
	var tmpY = (float64(height)/7)*0.7
	var leftBtnTmpX = float64(width)*0.85
	var leftBtnTmpY = float64(height)*0.9
	var rightBtnTmpX = float64(width)*0.85+leftAndRightSize+20
	var rightBtnTmpY = float64(height)*0.9
	var nextBtnTmpX = float64(width)*0.85+leftAndRightSize-17
	var nextBtnTmpY = float64(550)

	var playFirst = firstMove()


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

			if playFirst == 2{
				computerPlay(&player2, &table)
				playFirst = -1
			}

			if player1Active {
				switch event.(type) {
				case *sdl.MouseButtonEvent:

					mouseX := currentMouseState.x
					mouseY := currentMouseState.y


					//bank clicked
					if !currentMouseState.leftButton && previousMouseState.leftButton{
						if float64(mouseX) >= tmpX && float64(mouseX) <= tmpX+140 && float64(mouseY) >= tmpY && float64(mouseY) <= tmpY+140{
							if oneDominoOnly {
								addFromBank(&table)
								player1.numOfDominoesInDeck++
								oneDominoOnly = false
							}
						}
					}
					//left arrow on screen clicked
					if !currentMouseState.leftButton && previousMouseState.leftButton{
						if float64(mouseX) >= leftBtnTmpX && float64(mouseX) <= leftBtnTmpX+leftAndRightSize && float64(mouseY) >= leftBtnTmpY && float64(mouseY) <= leftBtnTmpY+leftAndRightSize{
							leftButtonClicked = 1
						}
					}
					//right arrow on screen clicked
					if !currentMouseState.leftButton && previousMouseState.leftButton{
						if float64(mouseX) >= rightBtnTmpX && float64(mouseX) <= rightBtnTmpX+leftAndRightSize && float64(mouseY) >= rightBtnTmpY && float64(mouseY) <= rightBtnTmpY+leftAndRightSize{
							leftButtonClicked = 0
						}
					}
					// next button clicked
					if !currentMouseState.leftButton && previousMouseState.leftButton{
						if float64(mouseX) >= nextBtnTmpX && float64(mouseX) <= nextBtnTmpX+leftAndRightSize && float64(mouseY) >= nextBtnTmpY && float64(mouseY) <= nextBtnTmpY+leftAndRightSize{
							if oneDominoOnly == false {
								oneDominoOnly = true
								if computerPlay(&player2, &table) {
									if isWon(&player2) {
										player2Won = true
									}
									player1Active = true
								} else {
									player1Active = true

								}
							}
						}
					}

					for i := 0; i < len(player1.deck); i++ {
						x := player1.deck[i].x
						y := player1.deck[i].y

						if float64(mouseX) >= (x-dominoHeight)*scaleSize && float64(mouseX) <= x*scaleSize && float64(mouseY) <= (y+dominoWidth)*scaleSize && float64(mouseY) >= y*scaleSize {
							if player1.deck[i].assigned==1{
								if play(&player1, i, &table) {
									if isWon(&player1) {
										player1Won = true
									}
									player1Active = !player1Active
									if computerPlay(&player2, &table) {
										oneDominoOnly = true
										if isWon(&player2) {
											player2Won = true
										}
										player1Active = !player1Active
									} else {
										player1Active = !player1Active
										oneDominoOnly = true
									}
								}
							}
						}
					}
				}
			}

			err = renderer.SetDrawColor(192, 192, 192, 192) 	//background
			if err != nil {
				fmt.Println("error: ", err)
				return
			}
			err = renderer.Clear()
			if err != nil {
				fmt.Println("error: ", err)
				return
			}

			if player1Won {
				if countingPoints{
					player1WonScene2.drawScene(renderer, finalSceneWidth, finalSceneHeight)
				}else{
					player1WonScene1.drawScene(renderer, finalSceneWidth, finalSceneHeight)
				}

			}else if player2Won {
				if countingPoints{
					player2WonScene2.drawScene(renderer, finalSceneWidth, finalSceneHeight)
				}else{
					player2WonScene1.drawScene(renderer, finalSceneWidth, finalSceneHeight)
				}


			}else{
				err = renderer.SetDrawColor(128, 0, 0, 0)//rectangles
				if err != nil {
					fmt.Println("error: ", err)
					return
				}
				err = renderer.FillRect(&sdl.Rect{X: width/16-20, Y: height-(height/4)-20, W: width/4*3, H: height/4})
				if err != nil {
					fmt.Println("error: ", err)
					return
				}
				err = renderer.FillRect(&sdl.Rect{X: width/16-20, Y: height/16-20, W: width/4*3, H: height/4})
				if err != nil {
					fmt.Println("error: ", err)
					return
				}

				leftBtn.drawButton(renderer, leftAndRightSize)
				rightBtn.drawButton(renderer, leftAndRightSize)
				nextBtn.drawButton(renderer, leftAndRightSize)
				err = renderer.SetScale(0.7, 0.7)
				if err != nil {
					fmt.Println("error: ", err)
					return
				}

				bankBtn.drawButton(renderer, bankSize)

				for _, dom := range player1.deck {
					if dom.assigned == 1 {
						dom.draw(renderer, 90.0, 0, 0)
					}
					if dom.assigned == 0 {
						if leftDominoCounter>6 || rightDominoCounter>5{
							dom.y = dom.y / 0.7
							err = renderer.SetScale(0.4, 0.4)
							if err != nil {
								fmt.Println("error: ", err)
								return
							}
						}else{
							err = renderer.SetScale(0.5, 0.5)
							if err != nil {
								fmt.Println("error: ", err)
								return
							}
						}
						dom.draw(renderer, dom.rotation, dominoWidth/2, dominoHeight/2)
						err = renderer.SetScale(0.7, 0.7)
						if err != nil {
							fmt.Println("error: ", err)
							return
						}
					}
				}

				for _, dom := range player2.deck {
					if dom.assigned == 2{
						dom.drawHiddenDomino(renderer)
						dom.draw(renderer,90, 0, 0)
					}
					if dom.assigned == 0 {
						if leftDominoCounter>6 || rightDominoCounter>5 {
							dom.y = dom.y / 0.7
							err = renderer.SetScale(0.4, 0.4)
							if err != nil {
								fmt.Println("error: ", err)
								return
							}
						}else{
							err = renderer.SetScale(0.5, 0.5)
							if err != nil {
								fmt.Println("error: ", err)
								return
							}
						}
						dom.draw(renderer, dom.rotation, dominoWidth/2, dominoHeight/2)
						err = renderer.SetScale(0.7, 0.7)
						if err != nil {
							fmt.Println("error: ", err)
							return
						}
					}
				}

				for _, dom := range player2.deck {
					if dom.assigned == 2{
						dom.drawHiddenDomino(renderer)
						
					}
				}
			}


			err = renderer.SetScale(1, 1)
			if err != nil {
				fmt.Println("error: ", err)
				return
			}
			renderer.Present()

			previousMouseState = currentMouseState
		}
	}
}

