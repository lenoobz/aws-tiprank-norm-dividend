package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TipRankAssetModel struct
type TipRankAssetModel struct {
	ID              *primitive.ObjectID     `bson:"_id,omitempty"`
	IsActive        bool                    `bson:"isActive,omitempty"`
	CreatedAt       int64                   `bson:"createdAt,omitempty"`
	ModifiedAt      int64                   `bson:"modifiedAt,omitempty"`
	Schema          string                  `bson:"schema,omitempty"`
	Ticker          string                  `bson:"ticker,omitempty"`
	Name            string                  `bson:"name,omitempty"`
	Yield           float64                 `bson:"yield,omitempty"`
	DividendHistory []*TipRankDividendModel `bson:"dividendHistory,omitempty"`
}

// TipRankDividendModel struct
type TipRankDividendModel struct {
	Dividend       float64    `bson:"dividend,omitempty"`
	ExDividendDate *time.Time `bson:"exDividendDate,omitempty"`
	RecordDate     *time.Time `bson:"recordDate,omitempty"`
	DividendDate   *time.Time `bson:"payoutDate,omitempty"`
}
