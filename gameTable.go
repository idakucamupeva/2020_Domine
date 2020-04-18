package main


import (
	//"github.com/veandco/go-sdl2/sdl"
	"fmt"
	"math/rand"
	"time"
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

var startPositionWidth = float64(width/2-dominoWidth/2)/(0.5)
var startPositionHeight = float64(height/2-dominoHeight/2)/ 0.5
var xLeft = startPositionWidth+dominoWidth
var yLeft = startPositionHeight
var xRight = startPositionWidth+2*dominoWidth
var yRight = startPositionHeight

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

	leftDominoCounter += 1

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

//	dom.x = xLeft
//	dom.y = yLeft
	dom.assigned = 0 //on left
//	xLeft -= dominoWidth

	if dom.right != tmpLeft {
		dom.rotation = 180
	}

	if leftDominoCounter<7{
		dom.x = xLeft
		dom.y = yLeft
		xLeft -= dominoWidth

	}else if leftDominoCounter==7{
		dom.rotation += 270
		//xLeft += dominoWidth
		dom.x = (xLeft + dominoWidth)*0.4 -10
		dom.y = yLeft + dominoHeight +10
		xLeft += dominoHeight
		yLeft += dominoWidth/2 + dominoHeight/2

	}else if leftDominoCounter>7{
		dom.rotation += 180
		xLeft += dominoWidth
		dom.x = xLeft
		dom.y = yLeft

	}
}

//adding domino on right side of the deck
func addDominoOnRight(table *gameTable, dom *domino){

	rightDominoCounter += 1

	tmpRight := table.right
	
	if dom.left == table.right{
		table.right = dom.right
	}else{
		table.right = dom.left
	}

	//TODO 
	
	//dom.x = xRight
	//dom.y = yRight
	dom.assigned = 0 
	//xRight += dominoWidth

	if dom.left != tmpRight {
		dom.rotation = 180
	}

	if rightDominoCounter<8{	//10
		dom.x = xRight
		dom.y = yRight
		xRight += dominoWidth
	}else if rightDominoCounter==8{	//10
		dom.x = xRight -  dominoHeight/2
		dom.y = yRight - dominoWidth/4 +10
		dom.rotation -= 90
		xRight = xRight + dominoWidth/2
		yRight = yRight - dominoWidth + dominoHeight/2 + 5
	}else{
		dom.rotation -=180
		xRight -= dominoWidth
		dom.x =xRight
		dom.y = yRight
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

	dom.x = xLeft
	dom.y = yLeft
	dom.assigned = 0 //on left
	xLeft -= dominoWidth
}

func (table gameTable) numOnLeft() int{
	return table.left
}

func (table gameTable) numOnRight() int{
	return table.right
}

func canBeAdded(table *gameTable, dom *domino) addingOnTable{
	if table.left == -1 {
		return onStartPosition
	} else if (dom.left == table.left || dom.left == table.right) && (dom.right == table.left || dom.right == table.right) && (dom.left != dom.right){  
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

	if tryAdd == onStartPosition{
		tmpDom := plr.deck[num]
		addDominoOnStart(table, &tmpDom)
		//hasADominoFromBank = append(hasADominoFromBank, false)

		plr.deck[num] = tmpDom //domino changes x and y

		return true
	}else if tryAdd == onLeft{
		tmpDom := plr.deck[num]
		addDominoOnLeft(table, &tmpDom)
		//hasADominoFromBank = append(hasADominoFromBank, false)

		plr.deck[num] = tmpDom

		return true
	}else if tryAdd == onRight{
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		//hasADominoFromBank = append(hasADominoFromBank, false)

		plr.deck[num] = tmpDom

		return true

	}else if tryAdd == onBoth{
		if leftButtonClicked == 1 { //Left button on screen is clicked
			tmpDom := plr.deck[num]
			addDominoOnLeft(table, &tmpDom)
			plr.deck[num] = tmpDom
		}else { // Right button on screen is clicked or no button is clicked
			tmpDom := plr.deck[num]
			addDominoOnRight(table, &tmpDom)
			plr.deck[num] = tmpDom
		}
		//hasADominoFromBank = append(hasADominoFromBank, false)
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
}

func computerPlay(plr *Player, table *gameTable) bool{

	for num :=0; num< len(plr.deck); num++{
		if plr.deck[num].assigned==2{
			tryAdd := canBeAdded(table, &plr.deck[num])
			if tryAdd == onStartPosition{
				tmpDom := plr.deck[num]
				addDominoOnStart(table, &tmpDom)
				plr.deck[num] = tmpDom //domino changes x and y
				return true
			}else if tryAdd == onNone{
				continue
			}else if tryAdd == onLeft{
				tmpDom := plr.deck[num]
				addDominoOnLeft(table, &tmpDom)
				plr.deck[num] = tmpDom
				return true
			}else if tryAdd == onRight{
				tmpDom := plr.deck[num]
				addDominoOnRight(table, &tmpDom)
				plr.deck[num] = tmpDom
				return true
			}else if tryAdd == onBoth{  //TODO
				tmpDom := plr.deck[num]
				addDominoOnRight(table, &tmpDom)
				plr.deck[num] = tmpDom
				return true
			}else{
				fmt.Println("Nema dominu :(")
				return true
			}
		}
	}
	addDominoFromBankToComputer()
	/*if computerPlay(&player2, table){
		return true
	}*/
	return false
}

func isWon(plr *Player) bool {
	var forCheck = 0
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

func addFromBank(){
	var dominoTmp domino
	var positionInMap = -1

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	bankIsEmpty := true		//if the bank is empty, return from the loop

	for{
		for i:=0;i < 28;i++{
			if dominoesMap[i].assigned == -1{
				bankIsEmpty = false
				
			}
		}

		if bankIsEmpty{
			fmt.Println("Nema domina u banci!!!")
			return
		}

		randNum := r1.Intn(28)
		if dominoesMap[randNum].assigned == -1{
			dominoTmp = dominoesMap[randNum]
			positionInMap = randNum
			break
		}

	}

	if positionInMap == -1{
		return
	}

	dominoTmp.assigned = 1	// TODO promeni na assigned kao argument da moze i player2
	var foundEmptyPlace = false

	for i := 0; i < len(player1.deck); i++ {
		if 	player1.deck[i].assigned==0 && hasADominoFromBank[i]==false{	// svaki put istu uzima jer nije obelezeno tako u player1 ili dominoMaps
			dominoTmp.x = tablePositionX + (float64)((i+1)*dominoWidth/2)
			dominoTmp.y = tablePositionY
			//	player1.deck[i] = dominoTmp
			foundEmptyPlace = true
			hasADominoFromBank[i] = true
			break
		}
	}
	if foundEmptyPlace==false{
		dominoTmp.x = tablePositionX + (float64)(dominoCounter*dominoWidth/2)
		dominoTmp.y = tablePositionY
		dominoCounter += 1
	}
	dominoesMap[positionInMap] =dominoTmp
	player1.deck = append(player1.deck, dominoTmp)
}

func addDominoFromBankToComputer(){

	fmt.Println("Added domino from bank to computer ")
	var dominoTmp domino
	var positionInMap = -1

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	bankIsEmpty := true		//if the bank is empty, return from the loop

	for{
		
		for i:=0;i < 28;i++{
			if dominoesMap[i].assigned == -1{
			bankIsEmpty = false
			
		}
	}

	if bankIsEmpty{
		return
	}
		randNum := r1.Intn(28)
		if dominoesMap[randNum].assigned == -1{
			dominoTmp = dominoesMap[randNum]
			positionInMap = randNum
			break
		}

	}

	if positionInMap == -1{
		return
	}

	dominoTmp.assigned = 2
	var foundEmptyPlace = false
/*
	for i := 0; i < len(player2.deck); i++ {
		if 	player2.deck[i].assigned==0 {	// svaki put istu uzima jer nije obelezeno tako u player1 ili dominoMaps
			dominoTmp.x = tablePositionXOpponent + (float64)((i+1)*dominoWidth/2)
			dominoTmp.y = tablePositionYOpponent
			//	player1.deck[i] = dominoTmp
			foundEmptyPlace = true
			break
		}
	}
*/	if foundEmptyPlace==false{
		dominoTmp.x = tablePositionXOpponent + (float64)(dominoCounterOpponent*dominoWidth/2)
		dominoTmp.y = tablePositionYOpponent
		dominoCounterOpponent += 1
	}
	dominoesMap[positionInMap] =dominoTmp
	hasComputerDominoFromBank = append(hasComputerDominoFromBank, false)
	player2.deck = append(player2.deck, dominoTmp)
}