package main

import "fmt"

func main(){
	//array declaration "fixed length"
	// var age[3]int = [3]int{20, 25, 30}
	var age = [3]int{20, 25, 30}

	names := [4]string{"John", "Doe", "Smith", "Alex"}
	names[1]= "Bob"
	fmt.Println(names,len(names))
	fmt.Println(age,len(age))

	//slices declaration "dynamic length"
	var scores = []int{100, 50, 60}
	scores = append(scores, 90)

	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[0:3]
	rangeTwo :=names[:len(names)-1]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
}