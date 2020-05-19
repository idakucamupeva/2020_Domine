package main


//struct for Players who play game
type Player struct{
	name string
	isHuman bool
	deck []domino 
	numOfDominoesInDeck int
}


//player constructor
func newPlayer(name string, isHuman bool, deck []domino) Player{
	num := 7
	return Player{
		name: name,
		isHuman: isHuman,
		deck: deck,
		numOfDominoesInDeck : num,
	}
}