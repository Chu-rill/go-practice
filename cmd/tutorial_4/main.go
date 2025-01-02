package main

import "fmt"

func main(){
	//aray 
	// var intArr [3]int32
	// var intArr [3]int32 = [3]int32{1, 2, 3}
	// intArr := [3]int32{1, 2, 3}
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr[2])

	//slice
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Printf("intSlice with append: %v\n", intSlice)

	//map basically an object
	// var myMap map[string]int32 = make(map[string]int32)
	var myMap2 = map[string]uint8{"Adam":23,"Sarah":25}
	// delete(myMap2, "Adam")
	fmt.Println(myMap2)

	//standard for loop
	for i:=0; i<10; i++{
		fmt.Println(i)
	}

	//loop through map
	for name,age:= range myMap2{
		fmt.Printf("Name: %v, Age:%v\n", name,age)
	}

	//loop through array
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
}