package htmlparser

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/crawler_grpc"
	"github.com/jae2274/goutils/terr"
)

func ParseScoreHtml(htmlStr string) (*crawler_grpc.SetReviewScoreRequest, error) {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))
	if err != nil {
		return nil, err
	}
	companyEle := doc.Find(".company")

	companyName, err := findCompanyName(companyEle)
	if err != nil {
		return nil, err
	}

	ratingNoEle := companyEle.Find(".rating_no")
	score, err := findReviewScore(ratingNoEle)
	if err != nil {
		return nil, err
	}
	reviewCount, err := findReviewCount(ratingNoEle)
	if err != nil {
		return nil, err
	}

	return &crawler_grpc.SetReviewScoreRequest{
		Site:        "blind",
		CompanyName: companyName,
		AvgScore:    score,
		ReviewCount: reviewCount,
	}, nil
}

func ParseReviews(html string) []*crawler_grpc.Review {
	return []*crawler_grpc.Review{}
}

func findCompanyName(doc *goquery.Selection) (string, error) {
	nameEle := doc.Find(".where .name")

	if len(nameEle.Nodes) == 0 {
		return "", terr.New("company name not found")
	}

	return strings.TrimSpace(nameEle.Nodes[0].LastChild.Data), nil
}

var scoreRegex = regexp.MustCompile(`^([0-4]\.\d|5\.0)$`)

func findReviewScore(doc *goquery.Selection) (int32, error) {
	scoreEle := doc.ChildrenFiltered(".rate")
	if len(scoreEle.Nodes) == 0 {
		return 0, terr.New("score not found")
	}

	scoreStr := strings.TrimSpace(scoreEle.Nodes[0].LastChild.Data)

	return ParseReviewScore(scoreStr)
}

func findReviewCount(doc *goquery.Selection) (int32, error) {
	countEle := doc.ChildrenFiltered(".count")
	if len(countEle.Nodes) == 0 {
		return 0, terr.New("review count not found")
	}

	countStr := strings.TrimSpace(countEle.Nodes[0].LastChild.Data)

	return ParseReviewCount(countStr)
}
