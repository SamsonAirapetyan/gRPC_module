package handler

import (
	"gRPC/data"
	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l: l}
}

func (pr *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		pr.getProdsuct(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (pr *Product) getProdsuct(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProduct()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, " Unable marshal json", http.StatusInternalServerError)
	}
}
