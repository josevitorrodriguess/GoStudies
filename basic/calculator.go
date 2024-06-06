package main

import f "fmt"

func main(){
	var num1,num2 float32
	var operation string

	f.Println("First number:")
	f.Scan(&num1)

	f.Println("Operation:(+,-,/,*):")
	f.Scan(&operation)
	
	f.Println("Scond number:")
	f.Scan(&num2)

	switch operation {
	case "+":
		f.Printf("Resultado: %.2f:", num1+num2)

	case "-": 
		f.Printf("Result: %.2f:", num1-num2)
	
	case "/": 
		f.Printf("Result: %.2f:", num1/num2)
	
	case "*": 
		f.Printf("Result: %.2f:", num1*num2)
	}
}
	


