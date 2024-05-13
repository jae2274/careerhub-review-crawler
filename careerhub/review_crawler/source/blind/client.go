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

func (bs *BlindSource) GetReviewScore(companyName string) (*source.ReviewScore, bool, error) {
	rc, err := bs.api.Call(&apiactor.Request{
		Method: "GET",
		Url:    fmt.Sprintf("https://www.teamblind.com/kr/company/%s/reviews", companyName),
	})

	if err != nil {
		if apiactor.IsHttpErrorWithStatusCode(err, 404) {
			return nil, false, nil
		} else {
			return nil, false, err
		}
	}

	score, err := ParseScoreReader(rc)
	if err != nil {
		return nil, false, err
	}

	return score, true, nil
}

func (bs *BlindSource) GetReviews(companyName string, page int) ([]*source.Review, error) {
	rc, err := bs.api.Call(&apiactor.Request{
		Method: "GET",
		Url:    fmt.Sprintf("https://www.teamblind.com/kr/company/%s/reviews?page=%d", companyName, page),
	})

	if err != nil {
		return nil, err
	}

	return ParseReviewsReader(rc)
}
