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
	plr1WonWidth = 600
	plr1WonHeight = 400
	trophyWidth = 203
	trophyHeight = 273
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

func (scene *FinalScene)updateScene(k, x_fin, y_fin float64){
	if scene.x <= x_fin{
		scene. x += k
	}
	if scene.y <= y_fin {
		scene.y += k
	}
}
