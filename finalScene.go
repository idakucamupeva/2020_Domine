package main

import(
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	
)


type FinalScene struct{
	tex *sdl.Texture
	x, y float64
}

const(
	finalSceneWidth = 800
	finalSceneHeight = 600
)

func textureFromBMPFinal(renderer *sdl.Renderer, filename string) *sdl.Texture{
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


//final scene constructor
func newFinalScene(renderer *sdl.Renderer, filename string, x, y float64) (scene FinalScene){
	scene.tex = textureFromBMPFinal(renderer, filename)
	scene.x = x
	scene.y = y

	return scene
}

func (scene *FinalScene)drawScene(renderer *sdl.Renderer, width, height int){
	x := scene.x 
	y := scene.y 
	renderer.Copy(scene.tex, &sdl.Rect{W: int32(width), H: int32(height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(width), H: int32(height)})
}

