package blind

import (
	"fmt"

	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/source"
	"github.com/jae2274/goutils/apiactor"
)

type BlindSource struct {
	api *apiactor.ApiActor
}

func NewBlindSource(api *apiactor.ApiActor) *BlindSource {
	return &BlindSource{api: api}
}

type WrappedError struct {
	CompanyName string
	Err         error
}

func newWrappedError(companyName string, err error) *WrappedError {
	return &WrappedError{CompanyName: companyName, Err: err}
}

func (we *WrappedError) Error() string {
	return fmt.Sprintf("company: %s. %s", we.CompanyName, we.Err.Error())
}

func (bs *BlindSource) GetReviewScore(companyName string) (*source.ReviewScoreResult, error) {
	rc, err := bs.api.Call(&apiactor.Request{
		Method: "GET",
		Url:    fmt.Sprintf("https://www.teamblind.com/kr/company/%s/reviews", companyName),
	})

	if err != nil {
		if apiactor.IsHttpErrorWithStatusCode(err, 404) {
			return &source.ReviewScoreResult{CompanyName: companyName, IsExist: false}, nil
		}
		return nil, newWrappedError(companyName, err)
	}

	score, err := ParseScoreReader(rc)
	if err != nil {
		return nil, newWrappedError(companyName, err)
	}

	return &source.ReviewScoreResult{CompanyName: companyName, ReviewScore: score, IsExist: true}, nil
}

func (bs *BlindSource) GetReviews(companyName string, page int) (*source.ReviewList, error) {
	rc, err := bs.api.Call(&apiactor.Request{
		Method: "GET",
		Url:    fmt.Sprintf("https://www.teamblind.com/kr/company/%s/reviews?page=%d", companyName, page),
	})

	if err != nil {
		return nil, newWrappedError(companyName, err)
	}

	reviews, err := ParseReviewsReader(rc)
	if err != nil {
		return nil, newWrappedError(companyName, err)
	}

	return &source.ReviewList{CompanyName: companyName, Page: page, Reviews: reviews}, nil
}
