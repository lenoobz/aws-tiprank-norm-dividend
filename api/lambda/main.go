package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-tiprank-norm-dividend/config"
	"github.com/hthl85/aws-tiprank-norm-dividend/infrastructure/repositories/mongodb/repos"
	dividends "github.com/hthl85/aws-tiprank-norm-dividend/usecase/asset-dividends"
	"github.com/hthl85/aws-tiprank-norm-dividend/usecase/tiprank-dividends"
)

type TipRankDividendRequest struct {
	Tickers []string `json:"tickers"`
}

func main() {
	lambda.Start(lambdaHandler)
}

func lambdaHandler(ctx context.Context, req TipRankDividendRequest) {
	log.Println("lambda handler is called")

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
		log.Fatalf("create TipRank asset mongo failed: %v\n", err)
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
	dividendService := dividends.NewService(dividendRepo, *tiprankService, zap)

	// try correlation context
	if err := dividendService.AddAssetDividendsByTickers(ctx, req.Tickers); err != nil {
		log.Fatal("add asset dividends by tickers failed")
	}
}
