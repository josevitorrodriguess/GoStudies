package main

import (
	"module/competition"
)

func main() {
	go competition.Write("Hello")
	competition.Write("Olá")
}
