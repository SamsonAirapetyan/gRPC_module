package handler

import (
	"gRPC/data"
	"net/http"
)

func (pr *Product) PostProduct(rw http.ResponseWriter, r *http.Request) {
	pr.l.Println("Handle POST Product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
