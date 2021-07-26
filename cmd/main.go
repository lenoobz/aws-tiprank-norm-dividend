package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	corid "github.com/lenoobz/aws-lambda-corid"
	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-tiprank-norm-dividend/config"
	"github.com/lenoobz/aws-tiprank-norm-dividend/infrastructure/repositories/mongodb/repos"
	assets "github.com/lenoobz/aws-tiprank-norm-dividend/usecase/asset-dividends"
	"github.com/lenoobz/aws-tiprank-norm-dividend/usecase/tiprank-dividends"
)

func main() {
	appConf := config.AppConf

	// create new logger
	zap, err := logger.NewZapLogger()
	if err != nil {
		log.Fatal("create app logger failed")
	}
	defer zap.Close()

	// create new repository
	tiprankRepo, err := repos.NewTipRankAssetMongo(nil, zap, &appConf.Mongo)
	if err != nil {
		log.Fatalf("create TipRank dividend mongo failed: %v\n", err)
	}
	defer tiprankRepo.Close()

	// create new repository
	dividendRepo, err := repos.NewAssetDividendMongo(nil, zap, &appConf.Mongo)
	if err != nil {
		log.Fatal("create dividend mongo failed")
	}
	defer dividendRepo.Close()

	// create new service
	tiprankService := tiprank.NewService(tiprankRepo, zap)
	dividendService := assets.NewService(dividendRepo, *tiprankService, zap)

	// try correlation context
	id, _ := uuid.NewRandom()
	ctx := corid.NewContext(context.Background(), id)

	// try correlation context
	if err := dividendService.AddAssetDividends(ctx); err != nil {
		log.Fatal("add asset dividends failed")
	}

	// try correlation context
	// if err := dividendService.InsertAssetDividendsByTickers(ctx, []string{"TSE:FAP"}); err != nil {
	// 	log.Fatal("add asset dividends by tickers failed")
	// }
}
