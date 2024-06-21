package main

import (
	"fmt"
	"module/functions"
)
func main() {

	soma := functions.Sum(10, 14)
	fmt.Println(soma)
	
	text := functions.TypeString("arroz")
	
	fmt.Println(text)

	sum, sub := functions.Calc(10,5)

	fmt.Println(sum,sub)
}
