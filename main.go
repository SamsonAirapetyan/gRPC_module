package main

import (
	"github.com/SamsonAirapetyan/gRPC_module/data"
	protos "github.com/SamsonAirapetyan/gRPC_module/protos/currency"
	"github.com/SamsonAirapetyan/gRPC_module/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	ex, err := data.NewExchangeRates(log)
	if err != nil {
		log.Error("Unable to create Exchange Rates", "error", err)
		os.Exit(1)
	}
	cs := server.NewCurrency(ex, log)

	protos.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)
	log.Info("Server is starting...")
	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	//Like listen and Serve in HTTP
	gs.Serve(l)
}
