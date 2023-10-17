package handler

import (
	"context"
	"gRPC/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l: l}
}

func (pr *Product) GetProduct(rw http.ResponseWriter, r *http.Request) {
	pr.l.Println("Handle GET Product")
	lp := data.GetProduct()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, " Unable marshal json", http.StatusInternalServerError)
	}
}

func (pr *Product) PostProduct(rw http.ResponseWriter, r *http.Request) {
	pr.l.Println("Handle POST Product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (pr *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	pr.l.Println("Handle PUT Product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.UpdateProduct(id, &prod)
	if err != nil {
		if err == data.ErrProductNotFound {
			http.Error(rw, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(rw, "Something went wrong", http.StatusBadRequest)
	}
}

type KeyProduct struct{}

func (p Product) MiddlewareProductVallidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Unable decode json", http.StatusInternalServerError)
			return
		}

		//оказывается context - это один из параметров структуры Request и можно в него запихнуть параметр (например из middleware нашу продукцию)
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}
