package dividends

import (
	"context"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-tiprank-norm-dividend/usecase/tiprank-assets"
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

// InsertAssetDividends adds new dividend
func (s *Service) InsertAssetDividends(ctx context.Context, tickers []string) error {
	s.log.Info(ctx, "add new asset dividend")

	assets, err := s.tiprankService.FindTipRankAssets(ctx, tickers)
	if err != nil {
		s.log.Error(ctx, "find all TipRank dividend failed", "error", err)
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
