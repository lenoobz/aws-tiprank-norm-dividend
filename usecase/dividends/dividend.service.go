package dividends

import (
	"context"

	logger "github.com/hthl85/aws-lambda-logger"
	tiprankdividends "github.com/hthl85/aws-tiprank-norm-dividend/usecase/tiprank-dividends"
)

// Service sector
type Service struct {
	dividendRepo           Repo
	tiprankDividendService tiprankdividends.Service
	log                    logger.ContextLog
}

// NewService create new service
func NewService(dividendRepo Repo, distributionService tiprankdividends.Service, log logger.ContextLog) *Service {
	return &Service{
		dividendRepo:           dividendRepo,
		tiprankDividendService: distributionService,
		log:                    log,
	}
}

// PopulateFundDividends populates fund dividends
func (s *Service) PopulateFundDividends(ctx context.Context) error {
	s.log.Info(ctx, "populate fund dividends")

	// distributions, err := s.distributionService.FindTipRankDividends(ctx)
	// if err != nil {
	// 	s.log.Error(ctx, "find all fund distribution failed", "error", err)
	// }

	// for _, distribution := range distributions {
	// 	dividend := distribution.MapFundDistributionToAssetDividend(ctx, s.log)

	// 	if err := s.dividendRepo.InsertAssetDividend(ctx, dividend); err != nil {
	// 		s.log.Error(ctx, "insert asset dividend", "error", err)
	// 		return err
	// 	}
	// }

	return nil
}
