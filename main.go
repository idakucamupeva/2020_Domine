package main


import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
	"math/rand"
	"time"
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

	var dominoesMap = make(map[int]domino)
	
	//domino's number
	var counter int
	counter = 0
	
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	broj := r1.Intn(28)
	
	startStr := "BMPdominoes/"
	
	for i:=0; i < 7; i++{
		for j:=i; j < 7; j++{
			startStr += strconv.Itoa(i) + "-" + strconv.Itoa(j) + ".bmp"
			dominoesMap[counter] = newDomino(renderer, startStr, i, j)
			counter++
			startStr = "BMPdominoes/"
		}
	}

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


		
		for i:=0; i<7; i++ {
			randNum := r1.Intn(28)
			domino1 := dominoesMap[randNum]
 			domino1.draw(renderer)
		}

		

		renderer.Present()
	}



	fmt.Println(broj)

	//var array []Domino
	//var player1 *Player = newPlayer("Stefan", array)

	fmt.Println("")


}