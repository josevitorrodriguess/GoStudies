package main

import (
	f "fmt"
	"strconv"
)

func main() {
	cardapio := [5]float32{7.00, 10.00, 8.50, 6.00, 6.00}
	var resposta int

	f.Println("Cardápio:")
	f.Printf(" 0-Cachorro-Quente: %.2f\n 1-Pizza: %.2f\n 2-Hamburguer: %.2f\n 3-Guaraná %.2f\n 4-Coca-Cola: %.2f\n", cardapio[0], cardapio[1], cardapio[2], cardapio[3], cardapio[4])
	f.Println("Digite o código dos lanches desejados:")
	f.Scanf("%d", &resposta)

	respostaArray := intToArray(resposta)

	var total float32

	for _, digit := range respostaArray {
		if digit >= 0 && digit < len(cardapio) {
			total += cardapio[digit]
		}
	}

	f.Printf("Total a pagar: %.2f", total)
}

func intToArray(num int) []int {
	numString := strconv.Itoa(num)
	result := make([]int, len(numString))

	for i, char := range numString {
		digit, _ := strconv.Atoi(string(char))
		result[i] = digit
	}

	return result
}
