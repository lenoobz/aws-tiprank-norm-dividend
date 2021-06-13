package tiprankdividends

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

// FindTipRankDividends finds all tip rank dividends
func (s *Service) FindTipRankDividends(ctx context.Context, tickers []string) ([]*entities.TipRankDividend, error) {
	s.log.Info(ctx, "find all tip rank dividends")
	return s.repo.FindTipRankDividends(ctx, tickers)
}
