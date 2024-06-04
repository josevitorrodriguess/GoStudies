package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type elemento struct {
	pessoa
	esquerda *elemento
	direita  *elemento
}

type arvoreBinaria struct {
	raiz *elemento
}

func (arvore *arvoreBinaria) adicionar(nome string, idade int) {
	novaPessoa := pessoa{nome, idade}
	novoElemento := &elemento{pessoa: novaPessoa}

	if arvore.raiz == nil {
		arvore.raiz = novoElemento
	} else {
		arvore.adicionarElemento(arvore.raiz, novoElemento)
	}
}

func (arvore *arvoreBinaria) adicionarElemento(raiz *elemento, novo *elemento) {
	if novo.idade < raiz.idade {
		if raiz.esquerda == nil {
			raiz.esquerda = novo
		} else {
			arvore.adicionarElemento(raiz.esquerda, novo)
		}
	} else {
		if raiz.direita == nil {
			raiz.direita = novo
		} else {
			arvore.adicionarElemento(raiz.direita, novo)
		}
	}
}

func (arvore *arvoreBinaria) imprimir(raiz *elemento) {
	if raiz != nil {
		arvore.imprimir(raiz.esquerda)
		fmt.Printf("Nome: %s, Idade: %d\n", raiz.pessoa.nome, raiz.pessoa.idade)
		arvore.imprimir(raiz.direita)
	}
}

func main() {
	arvore := &arvoreBinaria{}

	arvore.adicionar("Paulo", 20)
	arvore.adicionar("Lucas", 2)
	arvore.adicionar("João", 15)

	arvore.imprimir(arvore.raiz)
}
