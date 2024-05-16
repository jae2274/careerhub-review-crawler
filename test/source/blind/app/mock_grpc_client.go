package app

import (
	"context"

	"github.com/jae2274/careerhub-review-crawler/careerhub/review_crawler/crawler_grpc"
	"github.com/jae2274/goutils/llog"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockReviewGrpcClient struct {
	SetReviewScoreReqs     []*crawler_grpc.SetReviewScoreRequest
	SetNotExistReqs        []*crawler_grpc.SetNotExistRequest
	FinishCrawlingTaskReqs []*crawler_grpc.FinishCrawlingTaskRequest
	SaveCompanyReviewsReqs []*crawler_grpc.SaveCompanyReviewsRequest
}

func NewMockReviewGrpcClient() *MockReviewGrpcClient {
	return &MockReviewGrpcClient{
		SetReviewScoreReqs:     make([]*crawler_grpc.SetReviewScoreRequest, 0),
		SetNotExistReqs:        make([]*crawler_grpc.SetNotExistRequest, 0),
		FinishCrawlingTaskReqs: make([]*crawler_grpc.FinishCrawlingTaskRequest, 0),
		SaveCompanyReviewsReqs: make([]*crawler_grpc.SaveCompanyReviewsRequest, 0),
	}
}

func (m *MockReviewGrpcClient) GetCrawlingTasks(ctx context.Context, in *crawler_grpc.GetCrawlingTasksRequest, opts ...grpc.CallOption) (*crawler_grpc.GetCrawlingTasksResponse, error) {
	llog.Debug(ctx, "GetCrawlingTasks")
	if in.Site == mockSiteName {
		return &crawler_grpc.GetCrawlingTasksResponse{
			CompanyNames: []string{existedCompanyName, onePageReviewCompanyName, zeroReviewCompanyName, nonExistedCompanyName},
		}, nil
	}

	return &crawler_grpc.GetCrawlingTasksResponse{CompanyNames: []string{}}, nil
}

func (m *MockReviewGrpcClient) SetReviewScore(ctx context.Context, in *crawler_grpc.SetReviewScoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.SetReviewScoreReqs = append(m.SetReviewScoreReqs, in)
	return &emptypb.Empty{}, nil
}

func (m *MockReviewGrpcClient) SetNotExist(ctx context.Context, in *crawler_grpc.SetNotExistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.SetNotExistReqs = append(m.SetNotExistReqs, in)
	return &emptypb.Empty{}, nil
}
func (m *MockReviewGrpcClient) FinishCrawlingTask(ctx context.Context, in *crawler_grpc.FinishCrawlingTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.FinishCrawlingTaskReqs = append(m.FinishCrawlingTaskReqs, in)
	return &emptypb.Empty{}, nil
}

func (m *MockReviewGrpcClient) GetCrawlingTargets(ctx context.Context, in *crawler_grpc.GetCrawlingTargetsRequest, opts ...grpc.CallOption) (*crawler_grpc.GetCrawlingTargetsResponse, error) {
	finishedMap := make(map[string]bool)
	for _, req := range m.FinishCrawlingTaskReqs {
		finishedMap[req.Site+req.CompanyName] = true
	}

	targets := make([]*crawler_grpc.CrawlingTarget, 0)
	for _, req := range m.SetReviewScoreReqs {
		if _, ok := finishedMap[req.Site+req.CompanyName]; req.Site == in.Site && !ok {
			targets = append(targets, &crawler_grpc.CrawlingTarget{
				CompanyName:    req.CompanyName,
				TotalPageCount: req.TotalPageCount,
			})
		}
	}

	return &crawler_grpc.GetCrawlingTargetsResponse{Targets: targets}, nil
}

func (m *MockReviewGrpcClient) SaveCompanyReviews(ctx context.Context, in *crawler_grpc.SaveCompanyReviewsRequest, opts ...grpc.CallOption) (*crawler_grpc.SaveCompanyReviewsResponse, error) {
	m.SaveCompanyReviewsReqs = append(m.SaveCompanyReviewsReqs, in)
	return &crawler_grpc.SaveCompanyReviewsResponse{
		InsertedCount: int32(len(in.Reviews)),
	}, nil
}
