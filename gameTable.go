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
	
	tmpLeft := table.left

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

	if (dom.right != tmpLeft){	//domino should be rotate TODO flag
		dom.assigned = 5			
	}
	
} 

//adding domino on right side of the deck
func addDominoOnRight(table *gameTable, dom *domino){
	
	tmpRight := table.right
	
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

	if (dom.left != tmpRight){	//domino should be rotate TODO flag
		dom.assigned = 5		
	}
}

func addDominoOnStart(table *gameTable, dom *domino){
	
	
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



func (table gameTable) numOnLeft() int{
	return table.left
}


func (table gameTable) numOnRight() int{
	return table.right
}

	//if it's possible to add domino on left function returns 1, if on right then 2,
	//if both then 0, if none then -1, if it's start position then -2
func canBeAdded(table *gameTable, dom *domino) int{
	
	//table.rotate = false
	
	if table.left == -1 { //start position
		return -2
	} else if (dom.left == table.left || dom.left == table.right) && (dom.right == table.left || dom.right == table.right) && (dom.left != dom.right){  //TODO choosing side
		return 0 
	}else if dom.left == table.left || dom.right == table.left{
		return 1
	}else if dom.left == table.right || dom.right == table.right{
		return 2
	}else{
		return -1
	}
}

func play(plr *Player, num int, table *gameTable){

	tryAdd := canBeAdded(table, &plr.deck[num])
	fmt.Println(plr.deck[num].assigned)

	if tryAdd == -2{
		tmpDom := plr.deck[num]
		addDominoOnStart(table, &tmpDom)
		plr.deck[num] = tmpDom //domino changes x and y
		fmt.Println("-2 start")
	}else if tryAdd == 1{
		tmpDom := plr.deck[num]
		addDominoOnLeft(table, &tmpDom)
		plr.deck[num] = tmpDom 
		fmt.Println("1 left", plr.deck[num].assigned)
	}else if tryAdd == 2{
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		plr.deck[num] = tmpDom
		fmt.Println("2 right", plr.deck[num].assigned)
	}else if tryAdd == 0{  //TODO
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		plr.deck[num] = tmpDom
		fmt.Println("0 both", plr.deck[num].assigned)
	}else{
		fmt.Println("-1 none", plr.deck[num].assigned)
		return
	}

}
