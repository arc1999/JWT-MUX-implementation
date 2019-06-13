package main

import (
//	"log"
	//"net/http"
	"RestApi/Muteex"
)


func main(){
	Muteex.S1 = Muteex.E1{
		Muteex.Employee{Name: "ayushi",Age:20},
		Muteex.Employee{Name: "ayush",Age:10},
	}
	Muteex.Handlerequests()
}
