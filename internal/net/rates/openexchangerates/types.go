package openexchangerates

type latestRatesResponse struct {
	Rates map[string]float64 `json:"rates"`
}
