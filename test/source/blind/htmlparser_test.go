package blind

import (
	_ "embed"
	"testing"

	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/source"
	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/source/blind"
	"github.com/stretchr/testify/require"
)

//go:embed reviewPage.html
var reviewPage string

//go:embed noneReviewPage.html
var noneReviewPage string

//go:embed oneReviewPage.html
var oneReviewPage string

var pageReviews = []*source.Review{
	{Score: 30, Summary: "케바케인회사", EmploymentStatus: true, ReviewUserId: "J*****", JobType: "마케팅·홍보 전문가", UnixMilli: 1621123200000}, {Score: 40, Summary: "구글리는 마케팅용 단어일 뿐", EmploymentStatus: true, ReviewUserId: "p****", JobType: "해외영업·사업개발·마케팅", UnixMilli: 1683590400000}, {Score: 50, Summary: "국내 직장이라는 범주 안에서 대체불가한 회사", EmploymentStatus: true, ReviewUserId: "K****", JobType: "영업 전문가", UnixMilli: 1596758400000}, {Score: 20, Summary: "고인물들 대정치파티. 영업을 위해 구글 이미지를 사용할뿐 회사는 헬조선.", EmploymentStatus: true, ReviewUserId: "k*****", JobType: "해외영업·사업개발·마케팅", UnixMilli: 1677801600000}, {Score: 10, Summary: "언제든 짤릴수 있는곳. 무능한 경영진. 팀원들은 좋음. ", EmploymentStatus: true, ReviewUserId: "!*******", JobType: "IT 엔지니어", UnixMilli: 1677801600000}, {Score: 20, Summary: "많이 변함 특히 비개발은 헬한국패치화 진행중…", EmploymentStatus: true, ReviewUserId: "q******", JobType: "IT 디자이너", UnixMilli: 1668729600000}, {Score: 30, Summary: "고인물들과 말만 잘하는 사람들의 파티", EmploymentStatus: false, ReviewUserId: "v*****", JobType: "마케팅·홍보 전문가", UnixMilli: 1680220800000}, {Score: 30, Summary: "옛날만큼은 아니지만 그래도 괜찮은 회사", EmploymentStatus: true, ReviewUserId: "루**", JobType: "마케팅·홍보 전문가", UnixMilli: 1659225600000}, {Score: 40, Summary: "보상이 확실하고 몸은 편하나 멘탈이 갈릴 수 있음", EmploymentStatus: true, ReviewUserId: "P*****", JobType: "IT 디자이너", UnixMilli: 1623974400000}, {Score: 30, Summary: "과거에 좋았던 회사", EmploymentStatus: true, ReviewUserId: "u****", JobType: "IT 엔지니어", UnixMilli: 1694476800000}, {Score: 40, Summary: "늘 그렇듯, 실상은 별로인 곳", EmploymentStatus: true, ReviewUserId: "m*****", JobType: "해외영업·사업개발·마케팅", UnixMilli: 1692576000000}, {Score: 30, Summary: "한국회사", EmploymentStatus: true, ReviewUserId: "f*****", JobType: "마케팅·홍보 전문가", UnixMilli: 1687305600000}, {Score: 30, Summary: "분위기 많이 안좋아지고 있음", EmploymentStatus: true, ReviewUserId: "M*****", JobType: "마케팅·홍보 전문가", UnixMilli: 1670630400000}, {Score: 40, Summary: "신분세탁에 최고인 회사", EmploymentStatus: true, ReviewUserId: "F****", JobType: "영업 전문가", UnixMilli: 1630886400000}, {Score: 50, Summary: "기업 총점은 높으나 그만큼 부담과 실적을 내야 하는 회사", EmploymentStatus: false, ReviewUserId: "q*****", JobType: "마케팅·홍보 전문가", UnixMilli: 1652918400000}, {Score: 50, Summary: "공돌이 천국", EmploymentStatus: true, ReviewUserId: "K*****", JobType: "IT 엔지니어", UnixMilli: 1649980800000}, {Score: 50, Summary: "국내에서 갈데가 없게 만드는 회사", EmploymentStatus: true, ReviewUserId: "L*****", JobType: "마케팅·홍보 전문가", UnixMilli: 1606176000000}, {Score: 40, Summary: "한국적 글로벌적 특징이 모두 있음", EmploymentStatus: true, ReviewUserId: "v*****", JobType: "마케팅·홍보 전문가", UnixMilli: 1614902400000}, {Score: 30, Summary: "황금족쇄", EmploymentStatus: true, ReviewUserId: "D*****", JobType: "IT 엔지니어", UnixMilli: 1635292800000}, {Score: 30, Summary: "정규직에게는 좋은 회사", EmploymentStatus: true, ReviewUserId: "R*****", JobType: "IT 엔지니어", UnixMilli: 1614470400000}, {Score: 50, Summary: "아직도 대안이 없다", EmploymentStatus: true, ReviewUserId: "e*****", JobType: "해외영업·사업개발·마케팅", UnixMilli: 1704931200000}, {Score: 50, Summary: "내 인생에서 가장 중요한 것들을 배운곳", EmploymentStatus: false, ReviewUserId: "성*****", JobType: "IT 기획·매니지먼트 전문가", UnixMilli: 1622851200000}, {Score: 50, Summary: "마음먹은만큼 성장할수있는 회사", EmploymentStatus: true, ReviewUserId: "e*****", JobType: "경영전략·사업기획·분석전문가", UnixMilli: 1606262400000}, {Score: 30, Summary: "네임벨류, 워크 다 무난한 회사", EmploymentStatus: true, ReviewUserId: "b*****", JobType: "영업 전문가", UnixMilli: 1640736000000}, {Score: 30, Summary: "문돌이에겐 네임밸류 그이상 그이하도 아님", EmploymentStatus: true, ReviewUserId: "x*****", JobType: "마케팅·홍보 전문가", UnixMilli: 1634601600000}, {Score: 30, Summary: "개발자가 아닐경우 한계가 명확함", EmploymentStatus: true, ReviewUserId: "뜀****", JobType: "영업 전문가", UnixMilli: 1630972800000}, {Score: 50, Summary: "그래도 좋다", EmploymentStatus: true, ReviewUserId: "S*********", JobType: "해외영업·사업개발·마케팅", UnixMilli: 1678147200000}, {Score: 50, Summary: "최고의 회사", EmploymentStatus: true, ReviewUserId: "구******", JobType: "IT 엔지니어", UnixMilli: 1598227200000}, {Score: 40, Summary: "좋은 점 많아요 근데 더 좋아져야할 부분도 꽤 있습니다", EmploymentStatus: true, ReviewUserId: "e*****", JobType: "경영전략·사업기획·분석전문가", UnixMilli: 1648339200000}, {Score: 40, Summary: "자율성이 보장되는 회사", EmploymentStatus: true, ReviewUserId: "i*********", JobType: "IT 엔지니어", UnixMilli: 1597622400000},
}

var oneReview = &source.Review{
	Score:            50,
	Summary:          "성장중인 회사",
	EmploymentStatus: true,
	ReviewUserId:     "P*****",
	JobType:          "IT 엔지니어",
	UnixMilli:        1680652800000,
}

func TestParseScore(t *testing.T) {

	t.Run("return review score from html", func(t *testing.T) {
		scoreReq, err := blind.ParseScoreHtml(reviewPage)
		require.NoError(t, err)

		expected := &source.ReviewScore{
			Site:        "blind",
			AvgScore:    46,
			ReviewCount: 507,
			PageCount:   17,
			PageReviews: pageReviews,
		}

		// require.Equal(t, expected, scoreReq)
		require.Equal(t, expected.Site, scoreReq.Site)
		require.Equal(t, expected.AvgScore, scoreReq.AvgScore)
		require.Equal(t, expected.ReviewCount, scoreReq.ReviewCount)
		require.Equal(t, expected.PageCount, scoreReq.PageCount)
		require.Equal(t, len(expected.PageReviews), len(scoreReq.PageReviews))
		require.Equal(t, expected.PageReviews, scoreReq.PageReviews)
	})

	t.Run("return review score from one review html", func(t *testing.T) {
		scoreReq, err := blind.ParseScoreHtml(oneReviewPage)
		require.NoError(t, err)

		expected := &source.ReviewScore{
			Site:        "blind",
			AvgScore:    50,
			ReviewCount: 1,
			PageCount:   1,
			PageReviews: []*source.Review{oneReview},
		}

		// require.Equal(t, expected, scoreReq)
		require.Equal(t, expected.Site, scoreReq.Site)
		require.Equal(t, expected.AvgScore, scoreReq.AvgScore)
		require.Equal(t, expected.ReviewCount, scoreReq.ReviewCount)
		require.Equal(t, expected.PageCount, scoreReq.PageCount)
		require.Equal(t, len(expected.PageReviews), len(scoreReq.PageReviews))
		require.Equal(t, expected.PageReviews, scoreReq.PageReviews)
	})

	t.Run("return no review score from html", func(t *testing.T) {
		scoreReq, err := blind.ParseScoreHtml(noneReviewPage)
		require.NoError(t, err)

		expected := &source.ReviewScore{
			Site:        "blind",
			AvgScore:    0,
			ReviewCount: 0,
			PageCount:   0,
			PageReviews: []*source.Review{},
		}

		require.Equal(t, expected, scoreReq)
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
				score, err := blind.ParseReviewScore(tc.ScoreStr)
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
				_, err := blind.ParseReviewScore(tc.ScoreStr)
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
				count, err := blind.ParseReviewCount(tc.CountStr)
				require.NoError(t, err)
				require.Equal(t, tc.Expected, count)
			}
		})
	})
}

func TestParseReviews(t *testing.T) {

	t.Run("parse review user", func(t *testing.T) {
		type reviewUser struct {
			UserId  string
			JobType string
			Date    string
		}
		type testCase struct {
			ReviewUserStr string
			Expected      *reviewUser
		}

		testCases := []testCase{
			{
				ReviewUserStr: "· p**** · 해외영업·사업개발·마케팅 - 2023.05.09",
				Expected: &reviewUser{
					UserId:  "p****",
					JobType: "해외영업·사업개발·마케팅",
					Date:    "2023-05-09 00:00:00",
				},
			},
			{
				ReviewUserStr: "· 루** · 마케팅·홍보 전문가 - 2022.07.31",
				Expected: &reviewUser{
					UserId:  "루**",
					JobType: "마케팅·홍보 전문가",
					Date:    "2022-07-31 00:00:00",
				},
			},
			{
				ReviewUserStr: "· F**** · 영업 전문가 - 2021.09.06",
				Expected: &reviewUser{
					UserId:  "F****",
					JobType: "영업 전문가",
					Date:    "2021-09-06 00:00:00",
				},
			},
		}

		for _, tc := range testCases {
			userId, jobType, date, err := blind.ParseReviewUser(tc.ReviewUserStr)
			require.NoError(t, err)

			require.Equal(t, tc.Expected.UserId, userId)
			require.Equal(t, tc.Expected.JobType, jobType)
			require.Equal(t, tc.Expected.Date, date.Format("2006-01-02 15:04:05"))
		}
	})

	t.Run("parse reviews page", func(t *testing.T) {
		reviews, err := blind.ParseReviews(reviewPage)
		require.NoError(t, err)

		expected := pageReviews

		require.Equal(t, len(expected), len(reviews))
	})

	t.Run("parse one review page", func(t *testing.T) {
		reviews, err := blind.ParseReviews(oneReviewPage)
		require.NoError(t, err)

		require.Len(t, reviews, 1)
		require.Equal(t, oneReview, reviews[0])
	})

	t.Run("parse empty review page", func(t *testing.T) {
		reviews, err := blind.ParseReviews(noneReviewPage)
		require.NoError(t, err)

		require.Empty(t, reviews)
	})
}
