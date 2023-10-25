// Package classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version:1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// -application/json
// swagger: meta

package handler

import (
	protos "github.com/SamsonAirapetyan/gRPC_module/protos/currency"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	l  *log.Logger
	cc protos.CurrencyClient
}

func NewProduct(l *log.Logger, cc protos.CurrencyClient) *Product {
	return &Product{l: l, cc: cc}
}

func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
