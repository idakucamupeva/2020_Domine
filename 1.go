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

func main(){
	var dominoesMap = make(map[int]*Dominoe)
	var counter int
	counter = 0
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	broj := r1.Intn(28)
	var player1 []int
	player1 = append(player1,broj)
	for i:=0; i < 7; i++{
		for j:=i; j < 7; j++{
			dominoesMap[counter] = new (Dominoe)
			dominoesMap[counter].left = i
			dominoesMap[counter].right = j
			counter++
		}
	}


	for i:=0; i < counter; i++{
		fmt.Print(*dominoesMap[i], "   ")
	}
	fmt.Print(player1[0])
}