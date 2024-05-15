package source

type ReviewScoreResult struct {
	IsExist     bool
	CompanyName string
	ReviewScore *ReviewScore
}

type ReviewScore struct {
	Site        string
	AvgScore    int32
	ReviewCount int32
	PageCount   int32
}

type ReviewList struct {
	CompanyName string
	Page        int
	Reviews     []*Review
}

type Review struct {
	Score            int32
	Summary          string
	EmploymentStatus bool
	ReviewUserId     string
	JobType          string
	UnixMilli        int64
}

type Source interface {
	GetReviewScore(companyName string) (*ReviewScoreResult, error)
	GetReviews(companyName string, page int) (*ReviewList, error)
}
