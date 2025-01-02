package main

import "fmt"

//stuct and interface

type gasEngine struct{
	mpg uint8
	gallons uint8
}

func main(){
	var myEngine gasEngine = gasEngine{mpg: 30, gallons: 10}
	fmt.Println(myEngine.mpg,myEngine.gallons)
}