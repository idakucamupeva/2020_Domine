package main


import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"time"
)

const (
	screenWidth = 600
	screenHeight = 800
)



type Dominoe struct{
	left int
	right int
}


type Player struct{
	name string
	deck []Dominoe 
}


//player constructor
func newPlayer(name string, deck []Dominoe) *Player{
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

	var dominoesMap = make(map[int]*Dominoe)
	var counter int
	counter = 0
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	broj := r1.Intn(28)
	
	for i:=0; i < 7; i++{
		for j:=i; j < 7; j++{
			dominoesMap[counter] = new (Dominoe)
			dominoesMap[counter].left = i
			dominoesMap[counter].right = j
			counter++
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


		dominoe1 := newDominoe(renderer, "BMPdominoes/6-6.bmp")
		dominoe1.draw(renderer)

		renderer.Present()
	}


	fmt.Print("hahahha")
	fmt.Println(broj)

	//var array []Dominoe
	//var player1 *Player = newPlayer("Stefan", array)

	fmt.Println("")


}