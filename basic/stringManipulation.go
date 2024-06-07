package main

import (
	"fmt"
)

func reverse(str string ) (result string){
	for _, v := range str{
		result = string(v) + result
	}
	return
}

// func main(){
// 	var s string

// 	fmt.Println("type something(not count after backspace):")
// 	fmt.Scan(&s)

// 	fmt.Printf("\nNumber off charcters: %d\n\n",len(s))

// 	fmt.Printf("String reversed: %s", reverse(s))

// }