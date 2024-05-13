package htmlparser

import (
	_ "embed"
	"testing"

	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/htmlparser"
	"github.com/stretchr/testify/require"
)

//go:embed reviewPage.html
var reviewPage string

func TestHtmlParser(t *testing.T) {

	t.Run("return review score from html", func(t *testing.T) {
		scoreReq, err := htmlparser.ParseScoreHtml(reviewPage)
		require.NoError(t, err)

		expected := &crawler_grpc.SetReviewScoreRequest{
			Site:        "blind",
			CompanyName: "구글코리아",
			AvgScore:    46,
			ReviewCount: 507,
		}

		// require.Equal(t, expected, scoreReq)
		require.Equal(t, expected.Site, scoreReq.Site)
		require.Equal(t, expected.CompanyName, scoreReq.CompanyName)
		require.Equal(t, expected.AvgScore, scoreReq.AvgScore)
		require.Equal(t, expected.ReviewCount, scoreReq.ReviewCount)
	})

	t.Run("return score from score string", func(t *testing.T) {
		t.Run("success case", func(t *testing.T) {
			type testCase struct {
				ScoreStr string
				Expected int32
			}

			testCases := []testCase{
				{"5.0", 50},
				{"4.5", 45},
				{"5.0", 50},
				{"3.3", 33},
				{"0.0", 0},
			}

			for _, tc := range testCases {
				score, err := htmlparser.ParseReviewScore(tc.ScoreStr)
				require.NoError(t, err)
				require.Equal(t, tc.Expected, score)
			}
		})

		t.Run("error case", func(t *testing.T) {
			type testCase struct {
				ScoreStr string
			}

			testCases := []testCase{
				{"0.01"},
				{"00.0"},
				{"5.1"},
				{"6.0"},
				{"10.0"},
				{"5"},
				{"0"},
				{"-1.0"},
			}

			for _, tc := range testCases {
				_, err := htmlparser.ParseReviewScore(tc.ScoreStr)
				require.Error(t, err)
			}
		})
	})

	t.Run("return review count from count string", func(t *testing.T) {
		t.Run("success case", func(t *testing.T) {
			type testCase struct {
				CountStr string
				Expected int32
			}

			testCases := []testCase{
				{"390명", 390},
				{"507개 리뷰", 507},
				{"3,393개 리뷰", 3393},
				{"3,393명", 3393},
			}

			for _, tc := range testCases {
				count, err := htmlparser.ParseReviewCount(tc.CountStr)
				require.NoError(t, err)
				require.Equal(t, tc.Expected, count)
			}
		})
	})
}
