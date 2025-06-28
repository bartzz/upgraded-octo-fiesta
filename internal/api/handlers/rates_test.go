package handlers_test

import (
	"context"
	"encoding/json"
	"errors"
	"kryptonim-interview/internal/api"
	"kryptonim-interview/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"kryptonim-interview/internal/net/rates"
	"kryptonim-interview/internal/service"

	"github.com/gin-gonic/gin"
)

type mockRatesProvider struct {
	rates map[string]float64
	err   error
}

func (m mockRatesProvider) GetRates(ctx context.Context) (map[string]float64, error) {
	return m.rates, m.err
}

func TestRatesEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type wantPair struct {
		From string
		To   string
		Rate float64
	}

	tests := []struct {
		name      string
		provider  rates.ExchangeRatesProvider
		url       string
		wantCode  int
		wantPairs []wantPair
	}{
		{
			name:     "success USD<->EUR",
			provider: mockRatesProvider{rates: map[string]float64{"USD": 1, "EUR": 2}, err: nil},
			url:      "/rates?currencies=USD,EUR",
			wantCode: http.StatusOK,
			wantPairs: []wantPair{
				{"USD", "EUR", 2.0},
				{"EUR", "USD", 0.5},
			},
		},
		{
			name:     "bad query params",
			provider: mockRatesProvider{rates: map[string]float64{"USD": 1}, err: nil},
			url:      "/rates",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "provider error",
			provider: mockRatesProvider{rates: nil, err: errors.New("down")},
			url:      "/rates?currencies=USD,EUR",
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := service.NewRatesService(tc.provider)
			router := api.SetupRatesRouter(svc)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tc.url, nil)
			router.ServeHTTP(w, req)

			if w.Code != tc.wantCode {
				t.Fatalf("status = %d; want %d", w.Code, tc.wantCode)
			}

			if tc.wantCode != http.StatusOK {
				if len(w.Body.Bytes()) != 0 {
					t.Errorf("body = %q; want empty", w.Body.String())
				}
				return
			}

			// HTTP 200 - decode and compare pairs.
			var got []model.RatePair
			if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if len(got) != len(tc.wantPairs) {
				t.Fatalf("got %d pairs; want %d", len(got), len(tc.wantPairs))
			}
			for i, want := range tc.wantPairs {
				if got[i].From != want.From || got[i].To != want.To || got[i].Rate != want.Rate {
					t.Errorf("pair[%d] = %+v; want %+v", i, got[i], want)
				}
			}
		})
	}
}
