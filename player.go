package main


import (
	//"github.com/veandco/go-sdl2/sdl"
	
)


type Player struct{
	name string
	isHuman bool
	deck []domino 
}


//player constructor
func newPlayer(name string, isHuman bool, deck []domino) Player{
	return Player{
		name: name,
		isHuman: isHuman,
		deck: deck,
	}
}

