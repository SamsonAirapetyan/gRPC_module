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
		pr.getProduct(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		pr.postProduct(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (pr *Product) getProduct(rw http.ResponseWriter, r *http.Request) {
	pr.l.Println("Handle GET Product")
	lp := data.GetProduct()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, " Unable marshal json", http.StatusInternalServerError)
	}
}

func (pr *Product) postProduct(rw http.ResponseWriter, r *http.Request) {
	pr.l.Println("Handle POST Product")
	lp := &data.Product{}
	err := lp.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable decode json", http.StatusInternalServerError)
	}
	data.AddProduct(lp)
}
