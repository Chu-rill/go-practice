package main

import "fmt"

func main(){
	// menu := map[string]float64{
	// 	"soup": 4.99,
	// 	"pie": 7.99,
	// 	"salad": 6.99,
	// 	"toffee pudding": 3.55,
	// }

	// fmt.Println(menu)
	// fmt.Println(menu["pie"])

	// //looping through a map

	// for k, v := range menu {
	// 	fmt.Println(k, "-", v)
	// }

	//int as key type
	phonebook := map[int]string{
		1234567890: "James",
		587654321: "Moneypenny",
		425623324:"bob",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[1234567890])

	phonebook[1234567890] = "Bond"
	fmt.Println(phonebook)

}