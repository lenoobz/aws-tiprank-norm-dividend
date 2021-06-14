package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	corid "github.com/hthl85/aws-lambda-corid"
	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-tiprank-norm-dividend/config"
	"github.com/hthl85/aws-tiprank-norm-dividend/infrastructure/repositories/mongodb/repos"
	"github.com/hthl85/aws-tiprank-norm-dividend/usecase/tiprank-assets"
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

	// create new service
	tiprankService := tiprank.NewService(tiprankRepo, zap)

	// try correlation context
	id, _ := uuid.NewRandom()
	ctx := corid.NewContext(context.Background(), id)
	dividends, err := tiprankService.FindTipRankAssets(ctx, []string{"TSE:LGT.A", "TSE:LGT.B"})
	zap.Info(ctx, "TipRank Dividends", "dividends", dividends)
}
