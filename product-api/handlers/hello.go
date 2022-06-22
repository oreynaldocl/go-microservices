package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.l.Println("Saying Hello")

	d, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Oops an error", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(writer, "Hello %s", d)
}
