package app

import (
	"context"
	"fmt"

	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/source"
	"github.com/jae2274/goutils/cchan/pipe"
	"github.com/jae2274/goutils/llog"
)

type Application struct {
	grpcService *crawler_grpc.ReviewGrpcService
	src         source.Source
}

func NewApplication(grpcService *crawler_grpc.ReviewGrpcService, blindSource source.Source) *Application {
	return &Application{
		grpcService: grpcService,
		src:         blindSource,
	}
}

func (a *Application) SetReviewScores(ctx context.Context) error {

	companyNames, err := a.grpcService.GetCrawlingTasks(ctx, a.src.GetSiteName())
	if err != nil {
		return err
	}

	taskChan := make(chan string, len(companyNames))
	for _, task := range companyNames {
		taskChan <- task
	}
	close(taskChan)

	step1 := pipe.NewStep(nil, func(companyName string) (*source.ReviewScoreResult, error) {
		return a.src.GetReviewScore(companyName)
	})

	step2 := pipe.NewStep(nil, func(result *source.ReviewScoreResult) (struct{}, error) {
		var err error = nil
		if result.IsExist {
			err = a.grpcService.SetReviewScore(ctx, a.src.GetSiteName(), result.CompanyName, result.ReviewScore)
		} else {
			err = a.grpcService.SetNotExist(ctx, a.src.GetSiteName(), result.CompanyName)
		}

		return struct{}{}, err
	})

	finishChan, errChan := pipe.Pipeline2(ctx, taskChan, step1, step2)

	go drainChannels(ctx, finishChan)

	return checkErrorCount(ctx, errChan, 10)
}

func (a *Application) SaveReviews(ctx context.Context) error {
	targets, err := a.grpcService.GetCrawlingTargets(ctx, a.src.GetSiteName())
	if err != nil {
		return err
	}

	type targetPage struct {
		companyName string
		page        int32
	}

	targetChan := make(chan targetPage)
	go func() {
		for _, target := range targets {
			for page := target.TotalPageCount; page > 1; page-- {
				targetChan <- targetPage{companyName: target.CompanyName, page: page}
			}
		}
		close(targetChan)
	}()

	step1 := pipe.NewStep(nil, func(target targetPage) (*source.ReviewList, error) {
		return a.src.GetReviews(target.companyName, int(target.page))
	})

	step2 := pipe.NewStep(nil, func(reviews *source.ReviewList) (struct{}, error) {
		err := a.grpcService.SaveCompanyReviews(ctx, a.src.GetSiteName(), reviews)
		if err != nil {
			return struct{}{}, err
		}

		return struct{}{}, nil
	})

	finishChan, errChan := pipe.Pipeline2(ctx, targetChan, step1, step2)

	go drainChannels(ctx, finishChan)

	return checkErrorCount(ctx, errChan, 10)
}

func drainChannels[T any](ctx context.Context, c <-chan T) {
	for {
		select {
		case <-ctx.Done():
			return
		case _, ok := <-c:
			if !ok {
				return
			}
		}
	}
}

func checkErrorCount(ctx context.Context, errChan <-chan error, count int) error {

	for {
		select {
		case <-ctx.Done():
			return nil
		case err, ok := <-errChan:
			if !ok {
				return nil
			}
			llog.LogErr(ctx, err)
			count--

			if count == 0 {
				return fmt.Errorf("too many errors")
			}
		}
	}
}
