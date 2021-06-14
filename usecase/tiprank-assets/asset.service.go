package tiprank

import (
	"context"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-tiprank-norm-dividend/entities"
)

// Service sector
type Service struct {
	repo Repo
	log  logger.ContextLog
}

// NewService create new service
func NewService(repo Repo, log logger.ContextLog) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}

// FindTipRankAssets finds all tip rank assets
func (s *Service) FindTipRankAssets(ctx context.Context, tickers []string) ([]*entities.TipRankAsset, error) {
	s.log.Info(ctx, "find all TipRank assets")
	return s.repo.FindTipRankAssets(ctx, tickers)
}
