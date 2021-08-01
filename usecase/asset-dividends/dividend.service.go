package assets

import (
	"context"

	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-tiprank-norm-dividend/usecase/tiprank-dividends"
)

// Service sector
type Service struct {
	dividendRepo   Repo
	tiprankService tiprank.Service
	log            logger.ContextLog
}

// NewService create new service
func NewService(dividendRepo Repo, tiprankService tiprank.Service, log logger.ContextLog) *Service {
	return &Service{
		dividendRepo:   dividendRepo,
		tiprankService: tiprankService,
		log:            log,
	}
}

// AddAssetDividendsByTickers adds new asset dividends for given tickers
func (s *Service) AddAssetDividendsByTickers(ctx context.Context, tickers []string) error {
	s.log.Info(ctx, "adding new asset dividend by tickers", "tickers", tickers)

	assets, err := s.tiprankService.FindTipRankDividendsByTickers(ctx, tickers)
	if err != nil {
		s.log.Error(ctx, "find all TipRank dividends by tickers failed", "error", err)
	}

	for _, asset := range assets {
		dividend := asset.MapTipRankDividendToAssetDividend(ctx, s.log)

		if err := s.dividendRepo.InsertAssetDividend(ctx, dividend); err != nil {
			s.log.Error(ctx, "insert asset dividend failed", "error", err)
			return err
		}
	}

	return nil
}

// AddAssetDividends adds new asset dividends
func (s *Service) AddAssetDividends(ctx context.Context) error {
	s.log.Info(ctx, "adding new asset dividends")

	assets, err := s.tiprankService.FindTipRankDividends(ctx)
	if err != nil {
		s.log.Error(ctx, "find all TipRank assets failed", "error", err)
	}

	for _, asset := range assets {
		dividend := asset.MapTipRankDividendToAssetDividend(ctx, s.log)

		if err := s.dividendRepo.InsertAssetDividend(ctx, dividend); err != nil {
			s.log.Error(ctx, "insert asset dividend", "error", err)
			return err
		}
	}

	return nil
}
