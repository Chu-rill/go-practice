package main

import (
	"errors"
	"fmt"
)

func main(){
	// var printValue string = "Hello churchill"
	var a int = 10
	var b int = 0
	var result,remainder,error  = intDivision(a, b)
	//if/else
	if error != nil{
		fmt.Println(error.Error())
	}else if remainder == 0{
		fmt.Printf("The result of the integer division is %v ", result)
	}else{
		fmt.Printf("The result of the integer is %v with remainder %v", result, remainder)
	}
	//swith case
	switch{
		case error != nil:
			fmt.Println(error.Error())
		case remainder == 0:
			fmt.Printf("The result of the integer division is %v ", result)
		default:
			fmt.Printf("The result of the integer is %v with remainder %v", result, remainder)
	}
	
	// printMe(printValue)
}

// func printMe(printValue string){
// 	fmt.Println(printValue);	
// }
func intDivision(a int, b int) (int,int,error){
	var err error
	if b==0{
		err = errors.New("division by zero is not allowed")
		return 0,0,err
	}
	var result int = a/b
	var remainder int = a%b
	return result, remainder,err
}