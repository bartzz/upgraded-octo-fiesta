package service

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type CryptoToken struct {
	RateToUSD     decimal.Decimal
	DecimalPlaces int32
}

type ExchangeService struct {
	tokens map[string]CryptoToken
}

func NewExchangeService() *ExchangeService {
	return &ExchangeService{
		tokens: map[string]CryptoToken{
			"BEER":  {RateToUSD: decimal.RequireFromString("0.00002461"), DecimalPlaces: 18},
			"FLOKI": {RateToUSD: decimal.RequireFromString("0.0001428"), DecimalPlaces: 18},
			"GATE":  {RateToUSD: decimal.RequireFromString("6.87"), DecimalPlaces: 18},
			"USDT":  {RateToUSD: decimal.RequireFromString("0.999"), DecimalPlaces: 6},
			"WBTC":  {RateToUSD: decimal.RequireFromString("57037.22"), DecimalPlaces: 8},
		},
	}
}

func (s *ExchangeService) Exchange(from, to string, amount float64) (float64, error) {
	a, okA := s.tokens[from]
	b, okB := s.tokens[to]
	if !okA || !okB {
		return 0, fmt.Errorf("unsupported token: %s or %s", from, to)
	}

	amtDec := decimal.NewFromFloat(amount)
	raw := amtDec.Mul(a.RateToUSD).Div(b.RateToUSD)
	rounded := raw.Round(b.DecimalPlaces)

	result, _ := rounded.Float64()
	return result, nil
}
