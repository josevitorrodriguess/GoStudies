package main

import (
	"net/http"
)

func main() {
	http.Handle("/", Message{message: "Test"})
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func (m Message) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(m.message))
}

type Message struct {
	message string
}
