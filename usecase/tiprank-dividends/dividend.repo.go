package tiprank

import (
	"context"

	"github.com/hthl85/aws-tiprank-norm-dividend/entities"
)

///////////////////////////////////////////////////////////
// Fund Repository Interface
///////////////////////////////////////////////////////////

// Reader interface
type Reader interface {
	FindTipRankDividends(context.Context) ([]*entities.TipRankDividend, error)
	FindTipRankDividendsByTickers(context.Context, []string) ([]*entities.TipRankDividend, error)
}

// Writer interface
type Writer interface{}

// Repo interface
type Repo interface {
	Reader
	Writer
}
