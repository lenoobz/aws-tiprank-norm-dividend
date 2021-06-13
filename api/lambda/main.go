package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-tiprank-norm-dividend/config"
	"github.com/hthl85/aws-tiprank-norm-dividend/infrastructure/repositories/mongodb/repos"
	"github.com/hthl85/aws-tiprank-norm-dividend/usecase/dividends"
	tiprankdividends "github.com/hthl85/aws-tiprank-norm-dividend/usecase/tiprank-dividends"
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
	repo, err := repos.NewTipRankDividendMongo(nil, zap, &appConf.Mongo)
	if err != nil {
		log.Fatalf("create TipRank dividend mongo repo failed: %v\n", err)
	}
	defer repo.Close()

	// create new repository
	dividendRepo, err := repos.NewDividendMongo(nil, zap, &appConf.Mongo)
	if err != nil {
		log.Fatal("create dividend mongo repo failed")
	}
	defer dividendRepo.Close()

	// create new service
	tiprankService := tiprankdividends.NewService(repo, zap)
	dividendService := dividends.NewService(dividendRepo, *tiprankService, zap)

	// try correlation context
	if err := dividendService.InsertAssetDividends(ctx, req.Tickers); err != nil {
		log.Fatal("insert asset dividends failed")
	}
}
