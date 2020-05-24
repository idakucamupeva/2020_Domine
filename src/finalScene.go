package main

import(
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	
)

//FinalScene struct describes all final scenes (won and lost) used in game
type FinalScene struct{
	tex *sdl.Texture
	x, y float64
}

const(
	finalSceneWidth = 800
	finalSceneHeight = 600
)


//final scene constructor
func newFinalScene(renderer *sdl.Renderer, filename string, x, y float64) (scene FinalScene){
	scene.tex = textureFromBMP(renderer, filename)
	scene.x = x
	scene.y = y

	return scene
}

func (scene *FinalScene)drawScene(renderer *sdl.Renderer, width, height int){
	x := scene.x 
	y := scene.y 
	err := renderer.Copy(scene.tex, &sdl.Rect{W: int32(width), H: int32(height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(width), H: int32(height)})
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}

