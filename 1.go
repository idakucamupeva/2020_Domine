package main

import(
	"fmt"
	"math/rand"
	"time"
)



type Dominoe struct{
	left int
	right int
}


type Player struct{
	name string
	deck []Dominoe 
}


//player constructor
func newPlayer(name string, deck []Dominoe) *Player{
	return &Player{
		name: name,
		deck: deck,

	}

}


func main(){
	var dominoesMap = make(map[int]*Dominoe)
	var counter int
	counter = 0
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	broj := r1.Intn(28)
	
	for i:=0; i < 7; i++{
		for j:=i; j < 7; j++{
			dominoesMap[counter] = new (Dominoe)
			dominoesMap[counter].left = i
			dominoesMap[counter].right = j
			counter++
		}
	}
	

	var array []Dominoe
	var player1 *Player = newPlayer("Stefan", array)

	fmt.Println("")


}
