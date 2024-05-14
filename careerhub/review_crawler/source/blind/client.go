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

func (bs *BlindSource) GetReviewScore(companyName string) (*source.ReviewScoreResult, error) {
	rc, err := bs.api.Call(&apiactor.Request{
		Method: "GET",
		Url:    fmt.Sprintf("https://www.teamblind.com/kr/company/%s/reviews", companyName),
	})

	if err != nil {
		if apiactor.IsHttpErrorWithStatusCode(err, 404) {
			return &source.ReviewScoreResult{IsExist: false}, nil
		}
		return nil, err
	}

	score, err := ParseScoreReader(rc)
	if err != nil {
		return nil, err
	}

	return &source.ReviewScoreResult{ReviewScore: score, IsExist: true}, nil
}

func (bs *BlindSource) GetReviews(companyName string, page int) (*source.ReviewList, error) {
	rc, err := bs.api.Call(&apiactor.Request{
		Method: "GET",
		Url:    fmt.Sprintf("https://www.teamblind.com/kr/company/%s/reviews?page=%d", companyName, page),
	})

	if err != nil {
		return nil, err
	}

	reviews, err := ParseReviewsReader(rc)
	if err != nil {
		return nil, err
	}

	return &source.ReviewList{CompanyName: companyName, Page: page, Reviews: reviews}, nil
}
