package tiprank

import (
	"context"

	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-tiprank-norm-dividend/entities"
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

// FindTipRankDividendsByTickers finds all TipRank dividends for given tickers
func (s *Service) FindTipRankDividendsByTickers(ctx context.Context, tickers []string) ([]*entities.TipRankDividend, error) {
	s.log.Info(ctx, "finding TipRank dividends by tickers", "tickers", tickers)
	return s.repo.FindTipRankDividendsByTickers(ctx, tickers)
}

// FindTipRankDividends finds all TipRank dividends
func (s *Service) FindTipRankDividends(ctx context.Context) ([]*entities.TipRankDividend, error) {
	s.log.Info(ctx, "finding all TipRank dividends")
	return s.repo.FindTipRankDividends(ctx)
}
