package blind

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/source/blind"
	"github.com/jae2274/goutils/apiactor"
	"github.com/stretchr/testify/require"
)

func TestBlindSource(t *testing.T) {
	t.Run("return exist company review score", func(t *testing.T) {
		ctx := context.Background()

		blindSource := blind.NewBlindSource(apiactor.NewApiActor(ctx, 10))

		reviewScore, isExisted, err := blindSource.GetReviewScore("구글코리아")
		require.NoError(t, err)
		require.True(t, isExisted)

		require.Equal(t, "blind", reviewScore.Site)
		require.Equal(t, "구글코리아", reviewScore.CompanyName)
		require.Greater(t, reviewScore.AvgScore, int32(0))
		require.Greater(t, reviewScore.ReviewCount, int32(0))
		require.Greater(t, reviewScore.PageCount, int32(0))
	})

	t.Run("return not exist company review score", func(t *testing.T) {
		ctx := context.Background()

		blindSource := blind.NewBlindSource(apiactor.NewApiActor(ctx, 10))

		_, isExisted, err := blindSource.GetReviewScore("구글코리아아")
		require.NoError(t, err)
		require.False(t, isExisted)
	})

	t.Run("return exist company reviews", func(t *testing.T) {
		ctx := context.Background()

		blindSource := blind.NewBlindSource(apiactor.NewApiActor(ctx, 10))

		reviews, err := blindSource.GetReviews("구글코리아", 1)

		require.NoError(t, err)
		require.NotEmpty(t, reviews)
	})

	t.Run("return empty company reviews", func(t *testing.T) {
		ctx := context.Background()

		blindSource := blind.NewBlindSource(apiactor.NewApiActor(ctx, 10))

		reviews, err := blindSource.GetReviews("구글코리아", 100)

		require.NoError(t, err)
		require.Empty(t, reviews)
	})
}
