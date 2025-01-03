package main

import (
	"fmt"
	"sort"
)

func main(){
	// greetings := "hello there friends"

	// fmt.Println(strings.Contains(greetings, "hello"))
	// fmt.Println(strings.ReplaceAll(greetings, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greetings))
	// fmt.Println(strings.Index(greetings, "th"))
	// fmt.Println(strings.Split(greetings, " "))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "bowser"))
}