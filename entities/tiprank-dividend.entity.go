package entities

import (
	"time"
)

// TipRankDividend struct
type TipRankDividend struct {
	Ticker          string                   `json:"ticker,omitempty"`
	Name            string                   `json:"name,omitempty"`
	Yield           float64                  `json:"yield,omitempty"`
	DividendHistory map[int64]*DividendModel `json:"dividendHistory,omitempty"`
}

// DividendModel struct
type DividendModel struct {
	Dividend       float64    `json:"dividend,omitempty"`
	ExDividendDate *time.Time `json:"exDividendDate,omitempty"`
	RecordDate     *time.Time `json:"recordDate,omitempty"`
	DividendDate   *time.Time `json:"payoutDate,omitempty"`
}
