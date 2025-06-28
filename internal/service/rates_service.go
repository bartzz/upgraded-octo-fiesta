package service

import (
	"context"
	"fmt"
	"kryptonim-interview/internal/model"

	"kryptonim-interview/internal/net/rates"
)

type RatesService struct {
	provider rates.ExchangeRatesProvider
}

func NewRatesService(provider rates.ExchangeRatesProvider) *RatesService {
	return &RatesService{provider: provider}
}

func (s *RatesService) GetAllPairs(ctx context.Context, currencies []string) ([]model.RatePair, error) {
	if len(currencies) < 2 {
		return nil, fmt.Errorf("at least two currencies required")
	}

	ratesMap, err := s.provider.GetRates(ctx)
	if err != nil {
		return nil, err
	}

	var pairs []model.RatePair
	for _, a := range currencies {
		for _, b := range currencies {
			if a == b {
				continue
			}
			ra, okA := ratesMap[a]
			rb, okB := ratesMap[b]
			if !okA || !okB {
				return nil, fmt.Errorf("unknown currency: %s or %s", a, b)
			}
			pairs = append(pairs, model.RatePair{
				From: a,
				To:   b,
				Rate: rb / ra,
			})
		}
	}

	return pairs, nil
}
