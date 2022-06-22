package server

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"working/currency/protos"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle Get Rate", "base", rr.GetBase(), "destination", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}

func (c *Currency) mustEmbedUnimplementedCurrencyServer() {
	//TODO implement me
	panic("implement me")
}
