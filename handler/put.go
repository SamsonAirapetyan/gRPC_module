package handler

import (
	"gRPC/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

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
