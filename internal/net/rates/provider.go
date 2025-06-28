package rates

import "context"

type ExchangeRatesProvider interface {
	GetRates(ctx context.Context) (map[string]float64, error)
}
