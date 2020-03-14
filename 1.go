package main

import "fmt"

type Domina struct{
	levi int
	desni int

}

func main(){
	var mapaDomina = make(map[int]*Domina)
	
	var brojac int
	brojac = 0

	for i:=0; i < 7; i++{
		for j:=i; j < 7; j++{
			mapaDomina[brojac] = new (Domina)
			mapaDomina[brojac].levi = i
			mapaDomina[brojac].desni = j
			brojac++
		}
	} 
	
	//mapaDomina[0] = new (Domina)
	//mapaDomina[0].levi = 5

	for i:=0; i < brojac; i++{
		fmt.Print(*mapaDomina[i], "   ")

	}
}