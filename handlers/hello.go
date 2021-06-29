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

func (h *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.l.Println("Hello World")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Error!", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "Hello %s, ", data)
	rw.Write([]byte("Byeeee"))
}