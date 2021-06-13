package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	corid "github.com/hthl85/aws-lambda-corid"
	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-tiprank-norm-dividend/config"
	"github.com/hthl85/aws-tiprank-norm-dividend/infrastructure/repositories/mongodb/repos"
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

	// create new service
	tiprankService := tiprankdividends.NewService(repo, zap)

	// try correlation context
	id, _ := uuid.NewRandom()
	coridCtx := corid.NewContext(ctx, id)
	dividends, err := tiprankService.FindTipRankDividends(coridCtx, req.Tickers)
	zap.Info(coridCtx, "TipRank Dividends", "dividends", dividends)
}
