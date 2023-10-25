package handler

import (
	"context"
	"fmt"
	"gRPC/data"
	"net/http"
)

type KeyProduct struct{}

func (p Product) MiddlewareProductVallidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := data.FromJSON(prod, r.Body)
		if err != nil {
			http.Error(w, "Unable decode json", http.StatusInternalServerError)
			return
		}
		//Check validation
		if err := prod.Validate(); err != nil {
			http.Error(w, fmt.Sprintf("Problem with validation %s", err.Error()), http.StatusBadRequest)
			return
		}

		//оказывается context - это один из параметров структуры Request и можно в него запихнуть параметр (например из middleware нашу продукцию)
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}
