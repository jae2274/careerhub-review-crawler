package source

type ReviewScore struct {
	Site        string
	CompanyName string
	AvgScore    int32
	ReviewCount int32
	PageCount   int32
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
	GetReviewScore(companyName string) (*ReviewScore, bool, error)
	GetReviews(companyName string, page int) ([]*Review, error)
}
