package entities

import "time"

// AssetDividend struct
type AssetDividend struct {
	Ticker    string                     `json:"ticker,omitempty"`
	Dividends map[int64]*DividendDetails `json:"dividends,omitempty"`
}

// DividendDetails struct
type DividendDetails struct {
	DistDesc       string     `json:"distDesc,omitempty"`
	DistCode       string     `json:"distCode,omitempty"`
	Amount         float64    `json:"amount,omitempty"`
	ExDividendDate *time.Time `json:"exDividendDate,omitempty"`
	RecordDate     *time.Time `json:"recordDate,omitempty"`
	PayableDate    *time.Time `json:"payableDate,omitempty"`
}
