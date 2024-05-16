package app

import (
	"fmt"
	"time"

	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/source"
)

const (
	existedCompanyName       = "existedCompany"
	onePageReviewCompanyName = "onePageReviewCompany"
	zeroReviewCompanyName    = "zeroReviewCompany"
	nonExistedCompanyName    = "nonExistedCompany"
	mockSiteName             = "mockSite"
)

type MockSource struct {
}

func (m *MockSource) GetSiteName() string {
	return mockSiteName
}

func (m *MockSource) GetReviewScore(companyName string) (*source.ReviewScoreResult, error) {
	if existedCompanyName == companyName {
		return &source.ReviewScoreResult{
			IsExist:     true,
			CompanyName: companyName,
			ReviewScore: &source.ReviewScore{
				Site:        mockSiteName,
				AvgScore:    40,
				ReviewCount: 7,
				PageCount:   4,
				PageReviews: []*source.Review{
					newReview(newSummary(1, 2, 1)), newReview(newSummary(1, 2, 2)),
				},
			},
		}, nil
	} else if zeroReviewCompanyName == companyName {
		return &source.ReviewScoreResult{
			IsExist:     true,
			CompanyName: companyName,
			ReviewScore: &source.ReviewScore{
				Site:        mockSiteName,
				AvgScore:    0,
				ReviewCount: 0,
				PageCount:   0,
				PageReviews: []*source.Review{},
			},
		}, nil
	} else if onePageReviewCompanyName == companyName {
		return &source.ReviewScoreResult{
			IsExist:     true,
			CompanyName: companyName,
			ReviewScore: &source.ReviewScore{
				Site:        mockSiteName,
				AvgScore:    40,
				ReviewCount: 2,
				PageCount:   1,
				PageReviews: []*source.Review{
					newReview(newSummary(1, 2, 1)), newReview(newSummary(1, 2, 2)),
				},
			},
		}, nil
	}

	return &source.ReviewScoreResult{
		IsExist:     false,
		CompanyName: companyName,
	}, nil
}

func (m *MockSource) GetReviews(companyName string, page int) (*source.ReviewList, error) {
	if existedCompanyName == companyName {
		return &source.ReviewList{
			CompanyName: companyName,
			Page:        page,
			Reviews: []*source.Review{
				newReview(newSummary(page, 2, 1)), newReview(newSummary(page, 2, 2)),
			},
		}, nil
	}

	return &source.ReviewList{
		CompanyName: companyName,
		Page:        page,
		Reviews:     []*source.Review{},
	}, nil
}

func newReview(summary string) *source.Review {
	return &source.Review{
		Score:            40,
		Summary:          summary,
		EmploymentStatus: true,
		ReviewUserId:     "reviewUserId",
		JobType:          "jobType",
		UnixMilli:        time.Now().UnixMilli(),
	}
}

func newSummary(page int, pageSize int, add int) string {
	num := (page-1)*pageSize + add
	return fmt.Sprintf("summary%d", num)
}
