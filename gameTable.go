package main

import (
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

	dom.assigned = 0

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
	} else if (table.left==dom.left && table.right==dom.right) || (table.right==dom.left && table.left==dom.right){
	//if (dom.left == table.left || dom.left == table.right) && (dom.right == table.left || dom.right == table.right) && (dom.left != dom.right){
		return onBoth
	}else if dom.left == table.left || dom.right == table.left{
		return onLeft
	}else if dom.left == table.right || dom.right == table.right{
		return onRight
	}else{
		return onNone
	}
}

func hasMove(plr *Player, table *gameTable) bool {
	
	for _, dom := range plr.deck {
		if canBeAdded(table, &dom) != onNone && dom.assigned != 0{
			return true
		}
	}
	return false
}


func play(plr *Player, num int, table *gameTable) bool{

	tryAdd := canBeAdded(table, &plr.deck[num])

	if tryAdd == onStartPosition{
		tmpDom := plr.deck[num]
		addDominoOnStart(table, &tmpDom)
		hasADominoFromBank[num] = false
		plr.numOfDominoesInDeck--
		plr.deck[num] = tmpDom //domino changes x and y
		//fmt.Println("on start")
		return true
	}else if tryAdd == onLeft{
		tmpDom := plr.deck[num]
		addDominoOnLeft(table, &tmpDom)
		hasADominoFromBank[num] = false
		plr.numOfDominoesInDeck--
		plr.deck[num] = tmpDom
		//fmt.Println("on left")
		return true
	}else if tryAdd == onRight{
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		hasADominoFromBank[num] = false
		plr.numOfDominoesInDeck--
		plr.deck[num] = tmpDom
		//fmt.Println("on right")
		return true

	}else if tryAdd == onBoth{
		//fmt.Println("on both")
		if leftButtonClicked == 1 { //Left button on screen is clicked
			tmpDom := plr.deck[num]
			addDominoOnLeft(table, &tmpDom)
			plr.deck[num] = tmpDom
			hasADominoFromBank[num] = false
		}else { // Right button on screen is clicked or no button is clicked
			tmpDom := plr.deck[num]
			addDominoOnRight(table, &tmpDom)
			plr.deck[num] = tmpDom
			hasADominoFromBank[num] = false
		}
		plr.numOfDominoesInDeck--

		return true
	}else{
		//fmt.Println("none")
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

func checkComputerMoves(plr *Player, table *gameTable, num int, dom *domino) bool {

	
	
			tryAdd := canBeAdded(table, dom)
			if tryAdd == onStartPosition{
				tmpDom := dom
				addDominoOnStart(table, dom)
				dom = tmpDom //domino changes x and y
				hasComputerDominoFromBank[num] = false

				return true
			}else if tryAdd == onNone{
				return false
			}else if tryAdd == onLeft{
				tmpDom := dom
				addDominoOnLeft(table, dom)
				dom = tmpDom
				hasComputerDominoFromBank[num] = false

				return true
			}else if tryAdd == onRight{
				tmpDom := dom
				addDominoOnRight(table, dom)
				dom = tmpDom
				hasComputerDominoFromBank[num] = false

				return true
			}else if tryAdd == onBoth{  //TODO
				tmpDom := dom
				addDominoOnRight(table, tmpDom)
				dom = tmpDom
				hasComputerDominoFromBank[num] = false

				return true
			}else{
				return true
			}
		
	
	return false
}

func computerPlay(plr *Player, table *gameTable) bool{


	var moves []*domino		//slice of possible moves
	max := -1
	var maxDom *domino		
	var pos int

	for num :=0; num < len(plr.deck); num++{
		if plr.deck[num].assigned==2{
			if canBeAdded(table, &plr.deck[num]) != onNone{
				moves = append(moves, &plr.deck[num])
			}
		}
	}

	

	for i := 0; i < len(moves); i++ {
		if moves[i].left + moves[i].right > max{
			max = moves[i].left+moves[i].right
			maxDom = moves[i]
		
		}
	}

	



	

	if max == -1 {
		addDominoFromBankToComputer(table)
	}else{
		for num :=0; num< len(plr.deck); num++{
			if maxDom == &plr.deck[num]{
				pos = num
			}
		}
		if checkComputerMoves(plr, table, pos, maxDom){
			return true
		}
	}

	for num :=0; num < len(plr.deck); num++{
		if plr.deck[num].assigned==2{
			if canBeAdded(table, &plr.deck[num]) != onNone{
				moves = append(moves, &plr.deck[num])
			}
		}
	}

	for i := 0; i < len(moves); i++ {
		if moves[i].left + moves[i].right > max{
			max = moves[i].left+moves[i].right
			maxDom = moves[i]
		}

		
	}


	if max == -1 {
		return false
	}else{
		for num :=0; num< len(plr.deck); num++{
			if maxDom == &plr.deck[num]{
				pos = num
			}
		}
		if checkComputerMoves(plr, table, pos, maxDom){
			return true
		}
	}

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

func addFromBank(table *gameTable){
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
			if !hasMove(&player1, table) && !hasMove(&player2, table){
				
				if countPoints(&player1, &player2) == 1{
					player1Won = true
				}else if countPoints(&player1, &player2) == 2{
					player2Won = true
				}else{
					//TODO
				}

			}
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

	dominoTmp.assigned = 1
	var foundEmptyPlace = false
	var currentIndex int

	for i := 0; i < len(player1.deck); i++ {
		if 	player1.deck[i].assigned==0 && hasADominoFromBank[i]==false{	// svaki put istu uzima jer nije obelezeno tako u player1 ili dominoMaps
			dominoTmp.x = tablePositionX + (float64)((i+1)*dominoWidth/2)
			dominoTmp.y = tablePositionY
			//	player1.deck[i] = dominoTmp
			foundEmptyPlace = true
			currentIndex = i
			hasADominoFromBank[i] = true
			//dodato
			var dominoOnI domino = player1.deck[currentIndex]
			player1.deck[currentIndex] = dominoTmp
			player1.deck = append(player1.deck, dominoOnI)
			dominoesMap[positionInMap] =dominoTmp
			return
			//kraj dodatog
			break
		}
	}

	if foundEmptyPlace==false{
		dominoTmp.x = tablePositionX + (float64)(dominoCounter*dominoWidth/2)
		dominoTmp.y = tablePositionY
		dominoCounter += 1
		currentIndex = len(player1.deck)
		//indexInDeck := positionInPlayer1(positionInMap)
		hasADominoFromBank[currentIndex] = true
	}
	dominoesMap[positionInMap] =dominoTmp
	player1.deck = append(player1.deck, dominoTmp)

	//fja da na stavi iti na kraj a ovaj dominotmp na i-ti
	/*
	var dominoOnI domino = player1.deck[currentIndex]
	//fmt.Println("dominoI")
	//printDomino(&dominoOnI)
	player1.deck[currentIndex] = dominoTmp
	//fmt.Println("dominoI")
	//printDomino(&dominoOnI)
	player1.deck = append(player1.deck, dominoOnI)
	*/
	//player1.deck = append(player1.deck, dominoTmp)
}

func addDominoFromBankToComputer(table *gameTable){

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
			fmt.Println("Nema domina u banci za komp!!!")
			if !hasMove(&player1, table) && !hasMove(&player2, table){
				if countPoints(&player1, &player2) == 1{
					player1Won = true
				}else if countPoints(&player1, &player2) == 2{
					player2Won = true
				}else{
					//TODO
				}
			}
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
	var currentIndex int

	for i := 0; i < len(player2.deck); i++ {
		if 	player2.deck[i].assigned==0 && hasComputerDominoFromBank[i]==false{	// svaki put istu uzima jer nije obelezeno tako u player1 ili dominoMaps
			dominoTmp.x = tablePositionXOpponent + (float64)((i+1)*dominoWidth/2)
			dominoTmp.y = tablePositionYOpponent
			//	player1.deck[i] = dominoTmp
			foundEmptyPlace = true
			currentIndex = i
			hasComputerDominoFromBank[i] = true
			var dominoOnI domino = player2.deck[currentIndex]
			player2.deck[currentIndex] = dominoTmp
			player2.deck = append(player2.deck, dominoOnI)
			dominoesMap[positionInMap] =dominoTmp
			return
			break
		}
	}

	if foundEmptyPlace==false{
		dominoTmp.x = tablePositionXOpponent + (float64)(dominoCounterOpponent*dominoWidth/2)
		dominoTmp.y = tablePositionYOpponent
		dominoCounterOpponent += 1
		currentIndex = len(player2.deck)
		hasComputerDominoFromBank[currentIndex] = true
	}

	dominoesMap[positionInMap] =dominoTmp

	player2.deck = append(player2.deck, dominoTmp)

	//player2.deck = append(player2.deck, dominoTmp)
}

func countPoints(plr1, plr2 *Player) int {
	points1 := 0
	points2 := 0
	
	for _, dom := range plr1.deck {
		if dom.assigned == 1{
			points1 += dom.left + dom.right
		}
	}
	for _, dom := range plr2.deck {
		if dom.assigned == 2{
			points2 += dom.left + dom.right
		}
	}

	if points1 < points2{
		fmt.Println("Player1 won by countiing points")
		return 1
	}else if points2 < points1{

		fmt.Println("Player2 won by countiing points")
		return 2
	}else{
		return 0
	}
}