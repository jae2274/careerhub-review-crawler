package crawler_grpc

import (
	"context"

	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/source"
)

type ReviewGrpcService struct {
	grpcClient ReviewGrpcClient
}

func NewReviewGrpcService(grpcClient ReviewGrpcClient) *ReviewGrpcService {
	return &ReviewGrpcService{
		grpcClient: grpcClient,
	}
}

func (r *ReviewGrpcService) GetCrawlingTasks(ctx context.Context, siteName string) ([]string, error) {
	res, err := r.grpcClient.GetCrawlingTasks(ctx, &GetCrawlingTasksRequest{
		Site: siteName,
	})

	if err != nil {
		return nil, err
	}

	return res.CompanyNames, nil
}

func (r *ReviewGrpcService) SetReviewScore(ctx context.Context, siteName, companyName string, score *source.ReviewScore) error {
	_, err := r.grpcClient.SetReviewScore(ctx, &SetReviewScoreRequest{
		Site:           siteName,
		CompanyName:    companyName,
		AvgScore:       score.AvgScore,
		ReviewCount:    score.ReviewCount,
		TotalPageCount: score.PageCount,
	})
	if err != nil {
		return err
	}

	if len(score.PageReviews) > 0 {
		_, err := r.grpcClient.SaveCompanyReviews(ctx, convertReviewListToGrpc(siteName, &source.ReviewList{
			CompanyName: companyName,
			Page:        1,
			Reviews:     score.PageReviews,
		}))
		if err != nil {
			return err
		}
	}

	if score.PageCount <= 1 {
		_, err := r.grpcClient.FinishCrawlingTask(ctx, &FinishCrawlingTaskRequest{
			Site:        siteName,
			CompanyName: companyName,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ReviewGrpcService) SetNotExist(ctx context.Context, siteName, companyName string) error {
	_, err := r.grpcClient.SetNotExist(ctx, &SetNotExistRequest{
		Site:        siteName,
		CompanyName: companyName,
	})

	return err
}

func (r *ReviewGrpcService) GetCrawlingTargets(ctx context.Context, SiteName string) ([]*CrawlingTarget, error) {
	res, err := r.grpcClient.GetCrawlingTargets(ctx, &GetCrawlingTargetsRequest{Site: SiteName})
	if err != nil {
		return nil, err
	}

	return res.Targets, nil
}

func (r *ReviewGrpcService) SaveCompanyReviews(ctx context.Context, siteName string, reviewList *source.ReviewList) error {
	_, err := r.grpcClient.SaveCompanyReviews(ctx, convertReviewListToGrpc(siteName, reviewList))
	if err != nil {
		return err
	}

	if reviewList.Page == 2 {
		_, err := r.grpcClient.FinishCrawlingTask(ctx, &FinishCrawlingTaskRequest{
			Site:        siteName,
			CompanyName: reviewList.CompanyName,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func convertReviewListToGrpc(siteName string, reviews *source.ReviewList) *SaveCompanyReviewsRequest {
	grpcReviews := make([]*Review, len(reviews.Reviews))
	for i, review := range reviews.Reviews {
		grpcReviews[i] = &Review{
			Score:            review.Score,
			Summary:          review.Summary,
			EmploymentStatus: review.EmploymentStatus,
			ReviewUserId:     review.ReviewUserId,
			JobType:          review.JobType,
			UnixMilli:        review.UnixMilli,
		}
	}

	return &SaveCompanyReviewsRequest{
		Site:        siteName,
		CompanyName: reviews.CompanyName,
		Reviews:     grpcReviews,
	}
}
