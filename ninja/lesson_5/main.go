package main

import (
	"fmt"
	"strings"
)


func getInitials(name string) (string, string){
   s := strings.ToUpper(name)
   names := strings.Split(s, " ")

   var initials []string
   for _, n := range names{
         initials = append(initials, n[:1])
    }

    if len(initials) > 1{
        return initials[0], initials[1]
    }

    return initials[0], "_"
}


func main(){
    fn1,sn1 :=getInitials("john doe")
    fmt.Println(fn1,sn1)

    fn2,sn2 :=getInitials("bob pete")
    fmt.Println(fn2,sn2)

    fn3,sn3 :=getInitials("paul")
    fmt.Println(fn3,sn3)
}