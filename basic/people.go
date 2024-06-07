package main

import  "fmt"


// func main(){
// 	pessoa :=  Pessoa{}

// 	fmt.Println("Digite o nome:")
// 	fmt.Scan(&pessoa.nome)

// 	fmt.Println("Digite a idade:")
// 	fmt.Scan(&pessoa.idade)

// 	introduce(pessoa)

// 	if verififyAge(pessoa.idade) == true {
// 		fmt.Println("Maior de idade!")
// 	}else {
// 		fmt.Println("Menor de idade!")
// 	}
	
// }

type Pessoa struct{
	nome string
	idade int
}

func verififyAge(age int) bool{
	return age>= 18
}

func introduce( pessoa Pessoa){
	fmt.Printf("Olá, meu nome é %s e eu tenho %d anos!\n",pessoa.nome,pessoa.idade)
}