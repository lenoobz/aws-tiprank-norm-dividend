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
	FindTipRankAssets(context.Context, []string) ([]*entities.TipRankAsset, error)
}

// Writer interface
type Writer interface{}

// Repo interface
type Repo interface {
	Reader
	Writer
}
