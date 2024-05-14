package app

import (
	"context"

	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/source/blind"
	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/vars"
	"github.com/jae2274/goutils/apiactor"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/mw/grpcmw"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	SiteName = "blind"
)

func Run(ctx context.Context) {
	envVars, err := vars.Variables()

	checkErr(ctx, err)

	conn, err := grpc.NewClient(envVars.ReviewGrpcEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainStreamInterceptor(grpcmw.SetTraceIdStreamMW()),
		grpc.WithChainUnaryInterceptor(grpcmw.SetTraceIdUnaryMW()),
	)
	checkErr(ctx, err)

	client := crawler_grpc.NewReviewGrpcClient(conn)

	application := NewApplication(client, blind.NewBlindSource(apiactor.NewApiActor(ctx, 3000)))

	err = application.SetReviewScores(ctx)
	checkErr(ctx, err)

	err = application.SaveReviews(ctx)
	checkErr(ctx, err)
}

func checkErr(ctx context.Context, err error) {
	if err != nil {
		llog.LogErr(ctx, err)
		panic(err)
	}
}
