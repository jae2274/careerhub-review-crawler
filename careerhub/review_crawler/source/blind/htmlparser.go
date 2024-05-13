package blind

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/source"
	"github.com/jae2274/goutils/terr"
)

func ParseScoreHtml(htmlStr string) (*source.ReviewScore, error) {

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

	return &source.ReviewScore{
		Site:        "blind",
		CompanyName: companyName,
		AvgScore:    score,
		ReviewCount: reviewCount,
	}, nil
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

func ParseReviews(html string) ([]*source.Review, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	review_items := doc.Find(".review_item")
	reviews := make([]*source.Review, 0)
	errs := make([]error, 0)

	review_items.Each(func(i int, s *goquery.Selection) {
		review, err := parseReview(s)
		if err != nil {
			errs = append(errs, err)
			return
		}

		reviews = append(reviews, review)
	})

	return reviews, errors.Join(errs...)
}

func parseReview(doc *goquery.Selection) (*source.Review, error) {
	score, err := findScore(doc)

	if err != nil {
		return nil, err
	}

	summary, err := findSummary(doc)
	if err != nil {
		return nil, err
	}

	authEle := doc.Find(".auth")
	employmentStatus, err := findEmploymentStatus(authEle)
	if err != nil {
		return nil, err
	}

	name, roles, date, err := findReviewInfo(authEle)

	return &source.Review{
		Score:            score,
		Summary:          summary,
		EmploymentStatus: employmentStatus,
		ReviewUserId:     name,
		JobType:          roles,
		UnixMilli:        date.UnixMilli(),
	}, nil
}

func findScore(doc *goquery.Selection) (int32, error) {
	nodes := doc.Find(".rating > .num").Nodes

	if len(nodes) == 0 {
		return 0, terr.New("score not found")
	}

	scoreStr := strings.TrimSpace(nodes[0].LastChild.Data)
	score, err := ParseReviewScore(scoreStr)
	if err != nil {
		return 0, err
	}

	return score, nil
}

func findSummary(doc *goquery.Selection) (string, error) {
	nodes := doc.Find(".rvtit a").Nodes

	if len(nodes) == 0 {
		return "", terr.New("summary not found")
	}

	return strings.TrimSpace(nodes[0].LastChild.Data), nil
}

func findEmploymentStatus(doc *goquery.Selection) (bool, error) {
	nodes := doc.Find(".vrf").Nodes

	if len(nodes) == 0 {
		return false, nil
	}

	return strings.TrimSpace(nodes[0].LastChild.Data) == "현직원", nil
}

func findReviewInfo(doc *goquery.Selection) (string, string, time.Time, error) {
	if len(doc.Nodes) == 0 {
		return "", "", time.Time{}, terr.New("review info not found")
	}

	str := strings.TrimSpace(doc.Nodes[0].LastChild.Data)

	return ParseReviewUser(str)
}
