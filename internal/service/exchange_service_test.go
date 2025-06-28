package service

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestExchangeService(t *testing.T) {
	svc := NewExchangeService()

	tests := []struct {
		from    string
		to      string
		amount  float64
		want    decimal.Decimal
		wantErr bool
	}{
		{from: "BEER", to: "USDT", amount: 1.0, want: decimal.RequireFromString("0.000025"), wantErr: false},
		{from: "WBTC", to: "USDT", amount: 1.0, want: decimal.RequireFromString("57094.314314"), wantErr: false},
		{from: "INVALID", to: "USDT", amount: 1.0, want: decimal.Zero, wantErr: true},
		{from: "USDT", to: "USDT", amount: 123.456, want: decimal.RequireFromString("123.456"), wantErr: false},
		{from: "BEER", to: "BEER", amount: 999.0, want: decimal.RequireFromString("999.0"), wantErr: false},
		{from: "FLOKI", to: "USDT", amount: 0.0, want: decimal.Zero, wantErr: false},
		{from: "USDT", to: "WBTC", amount: 1.0, want: decimal.RequireFromString("0.00001751"), wantErr: false},
		{from: "WBTC", to: "USDT", amount: -1.0, want: decimal.RequireFromString("-57094.314314"), wantErr: false},
	}

	for _, tt := range tests {
		gotFloat, err := svc.Exchange(tt.from, tt.to, tt.amount)

		if tt.wantErr {
			if err == nil {
				t.Errorf("Exchange(%q,%q,%f) expected error, got none", tt.from, tt.to, tt.amount)
			}
			continue
		}
		if err != nil {
			t.Errorf("Exchange(%q,%q,%f) unexpected error: %v", tt.from, tt.to, tt.amount, err)
			continue
		}

		gotDec := decimal.NewFromFloat(gotFloat)
		if !gotDec.Equal(tt.want) {
			t.Errorf(
				"Exchange(%q,%q,%f) = %s, want %s",
				tt.from, tt.to, tt.amount,
				gotDec.String(), tt.want.String(),
			)
		}
	}
}
