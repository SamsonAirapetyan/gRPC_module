package handler

import (
	"gRPC/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (pr *Product) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	pr.l.Println("Delete product")
	err := data.DeleteProduct(id)
	if err != nil {
		if err == data.ErrProductNotFound {
			http.Error(rw, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(rw, "Something went wrong", http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusNoContent)
}
