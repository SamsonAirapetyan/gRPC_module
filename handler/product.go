package handler

import (
	"gRPC/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	if r.Method == http.MethodPut {
		//we have ti get id from URL
		regex := regexp.MustCompile("/([0-9]+)")
		g := regex.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, _ := strconv.Atoi(idString)
		pr.l.Println("got id", id)

		pr.updateProduct(id, rw, r)
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

func (pr *Product) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	pr.l.Println("Handle PUT Product")
	lp := &data.Product{}
	err := lp.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable decode json", http.StatusInternalServerError)
	}
	err = data.UpdateProduct(id, lp)
	if err != nil {
		if err == data.ErrProductNotFound {
			http.Error(rw, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(rw, "Something went wrong", http.StatusBadRequest)
	}
}
