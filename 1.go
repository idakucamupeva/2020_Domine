package main

import "fmt"

type Dominoe struct{
	left int
	right int
}

func main(){
	var dominoesMap = make(map[int]*Dominoe)
	var counter int
	counter = 0
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
}