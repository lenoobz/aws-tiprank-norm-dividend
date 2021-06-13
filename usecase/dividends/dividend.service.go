package dividends

import (
	"context"

	logger "github.com/hthl85/aws-lambda-logger"
	tiprankdividends "github.com/hthl85/aws-tiprank-norm-dividend/usecase/tiprank-dividends"
)

// Service sector
type Service struct {
	dividendRepo   Repo
	tiprankService tiprankdividends.Service
	log            logger.ContextLog
}

// NewService create new service
func NewService(dividendRepo Repo, tiprankService tiprankdividends.Service, log logger.ContextLog) *Service {
	return &Service{
		dividendRepo:   dividendRepo,
		tiprankService: tiprankService,
		log:            log,
	}
}

// InsertAssetDividends adds new dividend
func (s *Service) InsertAssetDividends(ctx context.Context, tickers []string) error {
	s.log.Info(ctx, "add new asset dividend")

	dividends, err := s.tiprankService.FindTipRankDividends(ctx, tickers)
	if err != nil {
		s.log.Error(ctx, "find all fund distribution failed", "error", err)
	}

	for _, dividend := range dividends {
		dividend := dividend.MapTipRankDividendToAssetDividend(ctx, s.log)

		if err := s.dividendRepo.InsertAssetDividend(ctx, dividend); err != nil {
			s.log.Error(ctx, "insert asset dividend", "error", err)
			return err
		}
	}

	return nil
}
