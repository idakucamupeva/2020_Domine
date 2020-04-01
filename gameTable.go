package main


import (
	//"github.com/veandco/go-sdl2/sdl"
	//"fmt"
)



type gameTable struct{
	left, right int

}


func newGameTable()  gameTable{
	left := -1
	right := -1
	return gameTable{
		left:left,
		right:right,
	}
}


//adding domino on left side of the deck
func (table gameTable) addDominoOnLeft(dom domino){
	if dom.left == table.left{
		table.left = dom.right
	}else{
		table.left = dom.left
	}
	
	//TODO changing domino position 
} 

//adding domino on right side of the deck
func (table gameTable) addDominoOnRight(dom domino){
	if dom.left == table.right{
		table.right = dom.right
	}else{
		table.right = dom.left
	}

	//TODO changing domino position
}


func (table gameTable) numOnLeft() int{
	return table.left
}


func (table gameTable) numOnRight() int{
	return table.right
}

	//if it's possible to add domino on left function returns 1, if on right then 2,
	//if both then 0, if none then -1
func (table gameTable) canBeAdded (dom domino) int{
	if (dom.left == table.left && dom.left == table.right) || (dom.right == table.left && dom.right == table.right){
		return 0
	}else if dom.left == table.left || dom.right == table.left{
		return 1
	}else if dom.left == table.right || dom.right == table.right{
		return 2
	}else{
		return -1
	}
}