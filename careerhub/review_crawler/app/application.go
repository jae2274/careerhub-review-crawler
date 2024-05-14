package app

import (
	"context"
	"fmt"

	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/source"
	"github.com/jae2274/goutils/cchan/pipe"
	"github.com/jae2274/goutils/llog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type application struct {
	grpcClient  crawler_grpc.ReviewGrpcClient
	blindSource source.Source
}

func NewApplication(grpcClient crawler_grpc.ReviewGrpcClient, blindSource source.Source) *application {
	return &application{
		grpcClient:  grpcClient,
		blindSource: blindSource,
	}
}

func (a *application) SetReviewScores(ctx context.Context) error {
	tasks, err := a.grpcClient.GetCrawlingTasks(ctx, &crawler_grpc.GetCrawlingTasksRequest{Site: SiteName})
	if err != nil {
		return err
	}

	taskChan := make(chan string, len(tasks.CompanyNames))
	for _, task := range tasks.CompanyNames {
		taskChan <- task
	}
	close(taskChan)

	step1 := pipe.NewStep(nil, func(companyName string) (*source.ReviewScoreResult, error) {
		return a.blindSource.GetReviewScore(companyName)
	})

	step2 := pipe.NewStep(nil, func(result *source.ReviewScoreResult) (*emptypb.Empty, error) {
		if result.IsExist {
			return a.grpcClient.SetReviewScore(ctx, &crawler_grpc.SetReviewScoreRequest{
				Site:        SiteName,
				CompanyName: result.ReviewScore.CompanyName,
				AvgScore:    result.ReviewScore.AvgScore,
				ReviewCount: result.ReviewScore.ReviewCount,
			})
		} else {
			return a.grpcClient.SetNotExist(ctx, &crawler_grpc.SetNotExistRequest{
				Site:        SiteName,
				CompanyName: result.ReviewScore.CompanyName,
			})
		}
	})

	finishChan, errChan := pipe.Pipeline2(ctx, taskChan, step1, step2)

	go drainChannels(ctx, finishChan)

	checkErrorCount(ctx, errChan, 10)

	return nil
}

func (a *application) SaveReviews(ctx context.Context) error {
	targets, err := a.grpcClient.GetCrawlingTargets(ctx, &crawler_grpc.GetCrawlingTargetsRequest{Site: SiteName})
	if err != nil {
		return err
	}

	type targetPage struct {
		companyName string
		page        int32
	}

	targetChan := make(chan targetPage)
	go func() {
		for _, target := range targets.Targets {
			for page := target.TotalPageCount; page > 0; page-- {
				targetChan <- targetPage{companyName: target.CompanyName, page: page}
			}
		}
		close(targetChan)
	}()

	step1 := pipe.NewStep(nil, func(target targetPage) (*source.ReviewList, error) {
		return a.blindSource.GetReviews(target.companyName, int(target.page))
	})

	step2 := pipe.NewStep(nil, func(reviews *source.ReviewList) (*emptypb.Empty, error) {
		_, err := a.grpcClient.SaveCompanyReviews(ctx, convertReviewListToGrpc(reviews))
		if err != nil {
			return nil, err
		}

		if reviews.Page == 1 {
			_, err := a.grpcClient.FinishCrawlingTask(ctx, &crawler_grpc.FinishCrawlingTaskRequest{
				Site:        SiteName,
				CompanyName: reviews.CompanyName,
			})

			if err != nil {
				return nil, err
			}
		}

		return &emptypb.Empty{}, nil
	})

	finishChan, errChan := pipe.Pipeline2(ctx, targetChan, step1, step2)

	go drainChannels(ctx, finishChan)

	checkErrorCount(ctx, errChan, 10)

	return nil
}

func convertReviewListToGrpc(reviews *source.ReviewList) *crawler_grpc.SaveCompanyReviewsRequest {
	grpcReviews := make([]*crawler_grpc.Review, len(reviews.Reviews))
	for i, review := range reviews.Reviews {
		grpcReviews[i] = &crawler_grpc.Review{
			Score:            review.Score,
			Summary:          review.Summary,
			EmploymentStatus: review.EmploymentStatus,
			ReviewUserId:     review.ReviewUserId,
			JobType:          review.JobType,
			UnixMilli:        review.UnixMilli,
		}
	}

	return &crawler_grpc.SaveCompanyReviewsRequest{
		CompanyName: reviews.CompanyName,
		Reviews:     grpcReviews,
	}
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
