package functions

import "fmt"

//this func sums two numbers and return the result
func Sum(n1,n2 int) int {
	return n1 + n2
}

//this func type a  string
func TypeString(txt string) string{
	fmt.Println(txt)
	return txt
}

//this func return the sum and subtraciton 
func Calc(n1,n2 int8) (int8, int8)  {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

