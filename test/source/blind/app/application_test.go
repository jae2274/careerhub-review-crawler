package app

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/app"
	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/crawler_grpc"
	"github.com/stretchr/testify/require"
)

func TestApplication(t *testing.T) {
	mockGrpcClient := NewMockReviewGrpcClient()
	grpcService := crawler_grpc.NewReviewGrpcService(mockGrpcClient)
	a := app.NewApplication(grpcService, &MockSource{})

	ctx := context.Background()
	err := a.SetReviewScores(ctx)
	require.NoError(t, err)

	require.Equal(t, 3, len(mockGrpcClient.SetReviewScoreReqs))
	require.Equal(t, 2, len(mockGrpcClient.SaveCompanyReviewsReqs))
	require.Equal(t, 1, len(mockGrpcClient.SetNotExistReqs))
	require.Equal(t, 2, len(mockGrpcClient.FinishCrawlingTaskReqs))

	err = a.SaveReviews(ctx)
	require.NoError(t, err)

	require.Equal(t, 3, len(mockGrpcClient.SetReviewScoreReqs))
	require.Equal(t, 5, len(mockGrpcClient.SaveCompanyReviewsReqs))
	require.Equal(t, 1, len(mockGrpcClient.SetNotExistReqs))
	require.Equal(t, 3, len(mockGrpcClient.FinishCrawlingTaskReqs))
}
