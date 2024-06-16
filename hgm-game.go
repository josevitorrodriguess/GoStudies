package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var resp string
	word := sortWord()
	hiddenWord := strings.Repeat("_", len(word))
	fmt.Println("HANGMAN GAME")

	for i := 0; i < 8; i++ {
		if !strings.Contains(hiddenWord, "_") {
			fmt.Println("YOU WIN! The correct word is:", word)
			return
		}

		fmt.Println("Type a char:")
		fmt.Scan(&resp)
		hiddenWord = game(resp, word, hiddenWord)
		fmt.Println(hiddenWord)
	}

	fmt.Println("GAME OVER! The correct word was:", word)
}

func game(sc string, word string, hiddenWord string) string {
	sc = strings.ToLower(sc)
	word = strings.ToLower(word)
	for i, c := range word {
		if string(c) == sc {
			hiddenWord = hiddenWord[:i] + sc + hiddenWord[i+1:]
		}
	}
	if !strings.Contains(word, sc) {
		fmt.Println("ERROR")
	} else {
		fmt.Println("CORRECT")
	}
	return hiddenWord
}

func verifyWord(c string, word string) bool {
	word = strings.ToLower(word)
	word = removeAccents(word)
	return strings.Contains(word, c)
}

func sortWord() string {
	word := randomWord(words)
	return word
}

func randomWord(slice []string) string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(slice))
	return slice[randomIndex]
}

func removeAccents(str string) string {
	return replacer.Replace(str)
}

var replacer = strings.NewReplacer(
	"á", "a", "à", "a", "ã", "a", "â", "a", "ä", "a",
	"é", "e", "è", "e", "ê", "e", "ë", "e",
	"í", "i", "ì", "i", "î", "i", "ï", "i",
	"ó", "o", "ò", "o", "õ", "o", "ô", "o", "ö", "o",
	"ú", "u", "ù", "u", "û", "u", "ü", "u",
	"ç", "c",
	"ñ", "n",
)

var words = []string{
	// words with letter A
	"Amigo", "Abacaxi", "Avião", "Anel", "Aventura", "Animal", "Atleta", "Arco-íris", "Alegria", "Alface",
	"Abelha", "Amor", "Alvo", "Arte", "Ação", "Agulha", "Aula", "Azeite", "Asa", "Aparência", "Ator",
	"Areia", "Armário", "Assado", "Anjo", "Água", "Almofada", "Avó", "Amendoim", "Árvore", "Aniversário",
	"Açúcar", "Acessório", "Aluno", "Agricultor", "Autêntico", "Azul", "Andar", "Amassar", "Amuleto",

	// words with letter B
	"Bola", "Brinquedo", "Barco", "Banana", "Bebida", "Bicicleta", "Bonito", "Bolo", "Botão", "Bolsa",
	"Beleza", "Borboleta", "Brilho", "Brócolis", "Beijo", "Banho", "Brisa", "Bússola", "Bairro", "Briga",
	"Baú", "Barulho", "Balão", "Branco", "Breve", "Beisebol", "Bebê", "Bandido", "Brasa", "Bola de neve",
	"Barba", "Barraca", "Bruxa", "Baixar", "Batom", "Batida", "Beira", "Biblioteca", "Beira-mar", "Beethoven",

	// words with letter C
	"Cachorro", "Cadeira", "Caneta", "Carro", "Computador", "Coração", "Coelho", "Chocolate", "Cachecol",
	"Cebola", "Cidade", "Casa", "Criança", "Cinto", "Camiseta", "Coruja", "Céu", "Chave", "Castelo",
	"Caneca", "Caminho", "Cacto", "Canto", "Cozinha", "Cachoeira", "Cogumelo", "Cenoura", "Ciclismo",
	"Clima", "Cama", "Crocodilo", "Colher", "Colina", "Caranguejo", "Corrida", "Celular", "Circo", "Copo", "Canário",

	// words with letter D
	"Dente", "Dado", "Dia", "Doce", "Dança", "Dedo", "Desenho", "Diamante", "Duna", "Desejo",
	"Direção", "Decisão", "Divertido", "Dente-de-leão", "Dose", "Dourado", "Donut", "Dieta",
	"Dançarino", "Dardo", "Dicionário", "Dólar", "Diamante",

	// words with letter E
	"Elefante", "Escola", "Estrela", "Espelho", "Esporte", "Elegante", "Esmalte", "Emoção", "Estante",
	"Esfera", "Energia", "Estudo", "Escada", "Escrito", "Entusiasmo", "Ervilha", "Engrenagem", "Enigma",
	"Equilíbrio", "Eucalipto", "Escorregador", "Efeito", "Esquadro", "Erro", "Esquilo", "Encontro",
	"Erosão", "Esqueleto", "Elogio", "Escrever", "Equipamento", "Escultura", "Enfermeira", "Espaço",
	"Elástico", "Eucaristia", "Edifício", "Esmeralda", "Embaixada",

	// words with letter F
	"Flor", "Feliz", "Fogo", "Fruta", "Folha", "Filme", "Frio", "Fada", "Festa", "Futebol",
	"Foca", "Feira", "Fantoche", "Fone", "Falar", "Faca", "Fazenda", "Fotografia", "Forno", "Fonte",
	"Fusca", "Feriado", "Formiga", "Ferro", "Filhote", "Fã", "Fechadura", "Fazer", "Folclore",
	"Ferrovia", "Felicidade", "Forte", "Fiação", "Fundo", "Flauta", "Feitiço", "Fivela", "Futuro",

	// word hard
	"Abstruso", "Acípite", "Ábaco", "Anacrônico", "Âmbito", "Antítese", "Aplacar", "Arcaico",
}
