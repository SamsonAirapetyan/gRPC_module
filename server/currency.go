package server

import (
	"context"
	"github.com/SamsonAirapetyan/gRPC_module/data"
	protos "github.com/SamsonAirapetyan/gRPC_module/protos/currency"
	"github.com/hashicorp/go-hclog"
	"io"
	"time"
)

type Currency struct {
	l             hclog.Logger
	rates         *data.ExchangeRates
	subscriptions map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest
}

func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	c := &Currency{l, r, make(map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest)}
	go c.handleUpdates()
	return c
}

func (c *Currency) handleUpdates() {
	ru := c.rates.MonitorRates(5 * time.Second)
	for range ru {
		c.l.Info("Got Updated rates")
		for k, v := range c.subscriptions {
			for _, rr := range v {
				r, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
				if err != nil {
					c.l.Error("Unable to get update rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
				}
				err = k.Send(&protos.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: r})
				if err != nil {
					c.l.Error("Unable to send updates rate")
				}
			}
		}
	}
}
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.l.Info("[GET RATE] Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	Rate, err := c.rates.GetRate(rr.Base.String(), rr.Destination.String())
	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: Rate}, nil
}

func (c *Currency) SubscribeRates(srv protos.Currency_SubscribeRatesServer) error {
	for {
		//Получение сообщения от клиента
		rr, err := srv.Recv()
		if err == io.EOF {
			c.l.Info("Client has closed connection")
			break
		}
		if err != nil {
			c.l.Error("Unable to read from client", "error", err)
			break
		}

		c.l.Info("Handle client request", "request", rr)
		rrs, ok := c.subscriptions[srv]
		if !ok {
			rrs = []*protos.RateRequest{}
		}
		rrs = append(rrs, rr)
		c.subscriptions[srv] = rrs
	}

	return nil
}
