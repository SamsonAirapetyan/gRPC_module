package server

import (
	"context"
	protos "github.com/SamsonAirapetyan/gRPC_module/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	l hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l: l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.l.Info("[GET RATE] Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
