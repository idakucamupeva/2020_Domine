package main


import (
	//"github.com/veandco/go-sdl2/sdl"
//	"fmt"
)



type gameTable struct{
	left, right int
	deck [] domino
}

func newGameTable()  gameTable{
	left := -1
	right := -1

	return gameTable{
		left:left,
		right:right,
	}
}

const(
	startPositionWidth = 0+dominoWidth
	startPositionHeight = 700
)

var x_left = startPositionWidth+dominoWidth
var y_left = startPositionHeight
var x_right = startPositionWidth+2*dominoWidth
var y_right = startPositionHeight
/*
var x_left = width+dominoWidth
var y_left = startPositionHeight
var x_right = width+2*dominoWidth
var y_right = startPositionHeight
*/

type addingOnTable int

const(
	onStartPosition addingOnTable = -2
	onLeft addingOnTable = 1
	onRight addingOnTable = 2
	onBoth addingOnTable = 0
	onNone addingOnTable = -1
)

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
func canBeAdded(table *gameTable, dom *domino) addingOnTable{

	if table.left == -1 { //start position
		return onStartPosition
	} else if (dom.left == table.left || dom.left == table.right) && (dom.right == table.left || dom.right == table.right) && (dom.left != dom.right){  //TODO choosing side
		return onBoth
	}else if dom.left == table.left || dom.right == table.left{
		return onLeft
	}else if dom.left == table.right || dom.right == table.right{
		return onRight
	}else{
		return onNone
	}
}


func play(plr *Player, num int, table *gameTable) bool{

	tryAdd := canBeAdded(table, &plr.deck[num])
//	fmt.Println(plr.deck[num].assigned)

	if tryAdd == onStartPosition{
		tmpDom := plr.deck[num]
		addDominoOnStart(table, &tmpDom)
		plr.deck[num] = tmpDom //domino changes x and y
		table.deck = append(table.deck, tmpDom)

		return true
	}else if tryAdd == onLeft{

		tmpDom := plr.deck[num]
		addDominoOnLeft(table, &tmpDom)
		plr.deck[num] = tmpDom
		table.deck = append(table.deck, tmpDom)

		return true

	}else if tryAdd == onRight{
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		plr.deck[num] = tmpDom
		table.deck = append(table.deck, tmpDom)

		return true

	}else if tryAdd == onBoth{  //TODO
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		plr.deck[num] = tmpDom
		table.deck = append(table.deck, tmpDom)
		return true

	}else{
	/*	var dominoTmp domino
		for _, element := range dominoesMap {
			if element.assigned ==-1{
				dominoTmp = element
				break
			}
		}
		dominoTmp.assigned = plr.deck[num].assigned
		plr.deck = append(plr.deck, dominoTmp)
		if play(plr, len(plr.deck)-1, table){
			return true
		}
	*/	return false
	}
	return false
}

func computerPlay(plr *Player, table *gameTable) bool{

	//for num :=0; num< len(plr.deck); num++{

	for num :=0; num< len(player2.deck); num++{
		tryAdd := canBeAdded(table, &player2.deck[num])
		//fmt.Println(plr.deck[num].assigned)

		if tryAdd == onStartPosition{
			tmpDom := plr.deck[num]
			addDominoOnStart(table, &tmpDom)
			plr.deck[num] = tmpDom //domino changes x and y
			table.deck = append(table.deck, tmpDom)
			return true
		}else if tryAdd == onNone{
			continue
		}else if tryAdd == onLeft{
			tmpDom := plr.deck[num]
			addDominoOnLeft(table, &tmpDom)
			plr.deck[num] = tmpDom
			table.deck = append(table.deck, tmpDom)
			return true

		}else if tryAdd == onRight{
			tmpDom := plr.deck[num]
			addDominoOnRight(table, &tmpDom)
			plr.deck[num] = tmpDom
			table.deck = append(table.deck, tmpDom)
			return true

		}else if tryAdd == onBoth{  //TODO
			tmpDom := plr.deck[num]
			addDominoOnRight(table, &tmpDom)
			plr.deck[num] = tmpDom
			table.deck = append(table.deck, tmpDom)
			return true

		}else{
		//	fmt.Println("-1 none", plr.deck[num].assigned)
			return false
		}
	}
	return false
}


func isWon(plr *Player) bool {
	var forCheck int = 0
	if plr==&player1{
		forCheck = 1
	}
	if plr==&player2{
		forCheck=2
	}

	for _, dom := range plr.deck {
		if dom.assigned == forCheck{
			return false
		}
	}
	return true
}

//


/*
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

func play(plr *Player, num int, table *gameTable) bool{

	tryAdd := canBeAdded(table, &plr.deck[num])
	fmt.Println(plr.deck[num].assigned)

	if tryAdd == -2{
		tmpDom := plr.deck[num]
		addDominoOnStart(table, &tmpDom)
		plr.deck[num] = tmpDom //domino changes x and y
		return true
		//fmt.Println("-2 start")
	}else if tryAdd == 1{
		tmpDom := plr.deck[num]
		addDominoOnLeft(table, &tmpDom)
		plr.deck[num] = tmpDom 
//		fmt.Println("1 left", plr.deck[num].assigned)
		return true

	}else if tryAdd == 2{
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		plr.deck[num] = tmpDom
//		fmt.Println("2 right", plr.deck[num].assigned)
		return true

	}else if tryAdd == 0{  //TODO
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		plr.deck[num] = tmpDom
//		fmt.Println("0 both", plr.deck[num].assigned)
		return true

	}else{
		fmt.Println("-1 none", plr.deck[num].assigned)
		return false
	}
	return false
}
//

func computerPlay(plr *Player, table *gameTable) bool{

	for num :=0; num< len(plr.deck); num++{
		tryAdd := canBeAdded(table, &plr.deck[num])
	//fmt.Println(plr.deck[num].assigned)

	if tryAdd == -2{
		tmpDom := plr.deck[num]
		addDominoOnStart(table, &tmpDom)
		plr.deck[num] = tmpDom //domino changes x and y
		return true
		//fmt.Println("-2 start")
	}else if tryAdd == -1{
		continue;
	}else if tryAdd == 1{
		tmpDom := plr.deck[num]
		addDominoOnLeft(table, &tmpDom)
		plr.deck[num] = tmpDom
		//		fmt.Println("1 left", plr.deck[num].assigned)
		return true

	}else if tryAdd == 2{
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		plr.deck[num] = tmpDom
		//		fmt.Println("2 right", plr.deck[num].assigned)
		return true

	}else if tryAdd == 0{  //TODO
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		plr.deck[num] = tmpDom
		//		fmt.Println("0 both", plr.deck[num].assigned)
		return true

	}else{
		fmt.Println("-1 none", plr.deck[num].assigned)
		return false
	}
	}
	return false
}
*/
