package entities

import (
	"context"
	"time"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-tiprank-norm-dividend/utils/ticker"
)

// TipRankDividend struct
type TipRankDividend struct {
	Ticker          string                     `json:"ticker,omitempty"`
	Name            string                     `json:"name,omitempty"`
	Yield           float64                    `json:"yield,omitempty"`
	Currency        string                     `json:"currency,omitempty"`
	DividendHistory map[int64]*DividendHistory `json:"dividendHistory,omitempty"`
}

// DividendHistory struct
type DividendHistory struct {
	Dividend       float64    `json:"dividend,omitempty"`
	ExDividendDate *time.Time `json:"exDividendDate,omitempty"`
	RecordDate     *time.Time `json:"recordDate,omitempty"`
	PayoutDate     *time.Time `json:"payoutDate,omitempty"`
}

// MapTipRankDividendToAssetDividend map TipRank dividend to asset dividend
func (f *TipRankDividend) MapTipRankDividendToAssetDividend(ctx context.Context, log logger.ContextLog) *AssetDividend {
	assetDividend := &AssetDividend{
		Ticker:    ticker.GenYahooTickerFromTipRankTicker(f.Ticker),
		Dividends: make(map[int64]*DividendDetails),
	}

	for key, val := range f.DividendHistory {
		dividendDetails := &DividendDetails{
			Amount:         val.Dividend,
			ExDividendDate: val.ExDividendDate,
			RecordDate:     val.RecordDate,
			PayableDate:    val.PayoutDate,
		}

		assetDividend.Dividends[key] = dividendDetails
	}

	return assetDividend
}
