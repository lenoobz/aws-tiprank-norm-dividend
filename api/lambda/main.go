package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	logger "github.com/hthl85/aws-lambda-logger"
)

type TipRankDividendRequest struct {
	Tickers []string `json:"tickers"`
}

func main() {
	// create new logger
	zap, err := logger.NewZapLogger()
	if err != nil {
		log.Fatal("create app logger failed")
	}
	defer zap.Close()

	lambda.Start(lambdaHandler)
}

func lambdaHandler(ctx context.Context, req TipRankDividendRequest) {
	log.Println("lambda handler is called")
	log.Println("request", req.Tickers)
}
