package models

import (
	"context"
	"time"

	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-tiprank-norm-dividend/consts"
	"github.com/lenoobz/aws-tiprank-norm-dividend/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AssetDividendModel struct
type AssetDividendModel struct {
	ID         *primitive.ObjectID             `bson:"_id,omitempty"`
	CreatedAt  int64                           `bson:"createdAt,omitempty"`
	ModifiedAt int64                           `bson:"modifiedAt,omitempty"`
	Enabled    bool                            `bson:"enabled"`
	Deleted    bool                            `bson:"deleted"`
	Schema     string                          `bson:"schema,omitempty"`
	Source     string                          `bson:"source,omitempty"`
	Ticker     string                          `bson:"ticker,omitempty"`
	Dividends  map[int64]*DividendDetailsModel `bson:"dividends,omitempty"`
}

// DividendDetailsModel struct
type DividendDetailsModel struct {
	DistDesc       string     `bson:"distDesc,omitempty"`
	DistCode       string     `bson:"distCode,omitempty"`
	Amount         float64    `bson:"amount,omitempty"`
	ExDividendDate *time.Time `bson:"exDividendDate,omitempty"`
	RecordDate     *time.Time `bson:"recordDate,omitempty"`
	PayableDate    *time.Time `bson:"payableDate,omitempty"`
}

// NewAssetDividendModel create asset dividend model
func NewAssetDividendModel(ctx context.Context, log logger.ContextLog, assetDividend *entities.AssetDividend, schemaVersion string) (*AssetDividendModel, error) {
	var assetDividendModel = &AssetDividendModel{
		ModifiedAt: time.Now().UTC().Unix(),
		Enabled:    true,
		Deleted:    false,
		Schema:     schemaVersion,
		Source:     consts.DATA_SOURCE,
		Ticker:     assetDividend.Ticker,
		Dividends:  map[int64]*DividendDetailsModel{},
	}

	for key, value := range assetDividend.Dividends {
		assetDividendModel.Dividends[key] = &DividendDetailsModel{
			DistDesc:       value.DistDesc,
			DistCode:       value.DistCode,
			Amount:         value.Amount,
			ExDividendDate: value.ExDividendDate,
			RecordDate:     value.RecordDate,
			PayableDate:    value.PayableDate,
		}
	}

	return assetDividendModel, nil
}
