package main


import (
	//"github.com/veandco/go-sdl2/sdl"
	"fmt"
)



type gameTable struct{
	left, right int

}

const(
	startPositionWidth = 0+dominoWidth
	startPositionHeight = 700 
)

var x_left = startPositionWidth+dominoWidth
var y_left = startPositionHeight
var x_right = startPositionWidth+2*dominoWidth
var y_right = startPositionHeight


func newGameTable()  gameTable{
	left := -1
	right := -1

	return gameTable{
		left:left,
		right:right,
	}
}


//adding domino on left side of the deck
func addDominoOnLeft(table *gameTable, dom *domino) {
	if table.left == -1{
		table.left = dom.left
		table.right = dom.right
	}else if dom.left == table.left{
		table.left = dom.right
	}else{
		table.left = dom.left
	}

	
	//TODO

	dom.x = float64(x_left) 
	dom.y = float64(y_left)
	dom.assigned = -2 //on left
	x_left -= dominoWidth
	
} 

//adding domino on right side of the deck
func addDominoOnRight(table *gameTable, dom *domino){
	if dom.left == table.right{
		table.right = dom.right
	}else{
		table.right = dom.left
	}

	//TODO 
	
	dom.x = float64(x_right)
	dom.y = float64(y_right)
	dom.assigned = 0 
	x_right += dominoWidth

}


func (table gameTable) numOnLeft() int{
	return table.left
}


func (table gameTable) numOnRight() int{
	return table.right
}

	//if it's possible to add domino on left function returns 1, if on right then 2,
	//if both then 0, if none then -1, if it's start position then -2
func (table gameTable) canBeAdded (dom domino) int{
	if table.left == -1 { //start position
		return -2
	} else if (dom.left == table.left || dom.left == table.right) && (dom.right == table.left || dom.right == table.right){
		return 0
	}else if dom.left == table.left || dom.right == table.left{
		return 1
	}else if dom.left == table.right || dom.right == table.right{
		return 2
	}else{
		return -1
	}
}

func play(plr Player, num int, table *gameTable){
	
	tryAdd := table.canBeAdded(plr.deck[num])

	if tryAdd == -2{
		tmpDom := plr.deck[num]
		addDominoOnLeft(table, &tmpDom)
		plr.deck[num] = tmpDom //domino changes x and y
		fmt.Println("-2")
	}else if tryAdd == 1{
		tmpDom := plr.deck[num]
		addDominoOnLeft(table, &tmpDom)
		plr.deck[num] = tmpDom 
		fmt.Println("1")
	}else if tryAdd == 2{
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		plr.deck[num] = tmpDom
		fmt.Println("2")
	}else if tryAdd == 0{  //TODO
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		plr.deck[num] = tmpDom
		fmt.Println("0")
	}else{
		fmt.Println("-1")
		return
	}

}
