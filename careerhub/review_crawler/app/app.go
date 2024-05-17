package app

import (
	"context"
	"os"

	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/source/blind"
	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/vars"
	"github.com/jae2274/goutils/apiactor"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/mw"
	"github.com/jae2274/goutils/mw/grpcmw"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	appName     = "review-crawler"
	serviceName = "careerhub"

	ctxKeyTraceID = string(mw.CtxKeyTraceID)

	// needRole = "ROLE_CAREERHUB_USER"
)

func initLogger(ctx context.Context) error {
	llog.SetMetadata("service", serviceName)
	llog.SetMetadata("app", appName)
	llog.SetDefaultContextData(ctxKeyTraceID)

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	llog.SetMetadata("hostname", hostname)

	return nil
}

func Run(ctx context.Context) {
	envVars, err := vars.Variables()
	checkErr(ctx, err)

	err = initLogger(ctx)
	checkErr(ctx, err)

	llog.Info(ctx, "Start application")

	conn, err := grpc.NewClient(envVars.ReviewGrpcEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainStreamInterceptor(grpcmw.SetTraceIdStreamMW()),
		grpc.WithChainUnaryInterceptor(grpcmw.SetTraceIdUnaryMW()),
	)
	checkErr(ctx, err)

	svc := crawler_grpc.NewReviewGrpcService(crawler_grpc.NewReviewGrpcClient(conn))

	application := NewApplication(svc, blind.NewBlindSource(apiactor.NewApiActor(ctx, 3000)))

	llog.Info(ctx, "Start SetReviewScores")
	err = application.SetReviewScores(ctx)
	checkErr(ctx, err)

	llog.Info(ctx, "Start SaveReviews")
	err = application.SaveReviews(ctx)
	checkErr(ctx, err)

	llog.Info(ctx, "Finish application")
}

func checkErr(ctx context.Context, err error) {
	if err != nil {
		llog.LogErr(ctx, err)
		os.Exit(1)
	}
}
