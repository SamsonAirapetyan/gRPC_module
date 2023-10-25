package handler

import (
	"context"
	"gRPC/data"
	protos "github.com/SamsonAirapetyan/gRPC_module/protos/currency"
	"net/http"
)

func (pr *Product) GetProduct(rw http.ResponseWriter, r *http.Request) {
	pr.l.Println("Handle GET Product")
	lp := data.GetProduct()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, " Unable marshal json", http.StatusInternalServerError)
	}
}

// Write here Ew are working with gRPC
func (pr *Product) ListSingle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	id := getProductID(r)
	pr.l.Println("[DEBUG] ger record id ", id)

	prod, err := data.GetProductByID(id)
	//get exchange

	//Here we have to create our Request (in this example its just stuff because result every time the same)
	rr := &protos.RateRequest{
		Base:        protos.Currencies_EUR,
		Destination: protos.Currencies_JPY,
	}
	//And here we are using The function in other repository
	resp, err := pr.cc.GetRate(context.Background(), rr)
	if err != nil {
		pr.l.Println("[Error] error with new Rate ", err)
	}
	//Here we are just changing the price in current Product
	prod.Price = prod.Price * resp.Rate
	err = data.ToJSON(prod, rw)
	if err != nil {
		pr.l.Println("[ERROR] with decoding product", err)
	}
}
