// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: careerhub/review_crawler/crawler_grpc/review.proto

package crawler_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReviewGrpcClient is the client API for ReviewGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewGrpcClient interface {
	GetCrawlingTasks(ctx context.Context, in *GetCrawlingTasksRequest, opts ...grpc.CallOption) (*GetCrawlingTasksResponse, error)
	SetReviewScore(ctx context.Context, in *SetReviewScoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetNotExist(ctx context.Context, in *SetNotExistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCrawlingTargets(ctx context.Context, in *GetCrawlingTargetsRequest, opts ...grpc.CallOption) (*GetCrawlingTargetsResponse, error)
	SaveCompanyReviews(ctx context.Context, in *SaveCompanyReviewsRequest, opts ...grpc.CallOption) (*SaveCompanyReviewsResponse, error)
	FinishCrawlingTask(ctx context.Context, in *FinishCrawlingTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type reviewGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewGrpcClient(cc grpc.ClientConnInterface) ReviewGrpcClient {
	return &reviewGrpcClient{cc}
}

func (c *reviewGrpcClient) GetCrawlingTasks(ctx context.Context, in *GetCrawlingTasksRequest, opts ...grpc.CallOption) (*GetCrawlingTasksResponse, error) {
	out := new(GetCrawlingTasksResponse)
	err := c.cc.Invoke(ctx, "/careerhub.review_service.crawler_grpc.ReviewGrpc/getCrawlingTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewGrpcClient) SetReviewScore(ctx context.Context, in *SetReviewScoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/careerhub.review_service.crawler_grpc.ReviewGrpc/setReviewScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewGrpcClient) SetNotExist(ctx context.Context, in *SetNotExistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/careerhub.review_service.crawler_grpc.ReviewGrpc/setNotExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewGrpcClient) GetCrawlingTargets(ctx context.Context, in *GetCrawlingTargetsRequest, opts ...grpc.CallOption) (*GetCrawlingTargetsResponse, error) {
	out := new(GetCrawlingTargetsResponse)
	err := c.cc.Invoke(ctx, "/careerhub.review_service.crawler_grpc.ReviewGrpc/getCrawlingTargets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewGrpcClient) SaveCompanyReviews(ctx context.Context, in *SaveCompanyReviewsRequest, opts ...grpc.CallOption) (*SaveCompanyReviewsResponse, error) {
	out := new(SaveCompanyReviewsResponse)
	err := c.cc.Invoke(ctx, "/careerhub.review_service.crawler_grpc.ReviewGrpc/saveCompanyReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewGrpcClient) FinishCrawlingTask(ctx context.Context, in *FinishCrawlingTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/careerhub.review_service.crawler_grpc.ReviewGrpc/finishCrawlingTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewGrpcServer is the server API for ReviewGrpc service.
// All implementations must embed UnimplementedReviewGrpcServer
// for forward compatibility
type ReviewGrpcServer interface {
	GetCrawlingTasks(context.Context, *GetCrawlingTasksRequest) (*GetCrawlingTasksResponse, error)
	SetReviewScore(context.Context, *SetReviewScoreRequest) (*emptypb.Empty, error)
	SetNotExist(context.Context, *SetNotExistRequest) (*emptypb.Empty, error)
	GetCrawlingTargets(context.Context, *GetCrawlingTargetsRequest) (*GetCrawlingTargetsResponse, error)
	SaveCompanyReviews(context.Context, *SaveCompanyReviewsRequest) (*SaveCompanyReviewsResponse, error)
	FinishCrawlingTask(context.Context, *FinishCrawlingTaskRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedReviewGrpcServer()
}

// UnimplementedReviewGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedReviewGrpcServer struct {
}

func (UnimplementedReviewGrpcServer) GetCrawlingTasks(context.Context, *GetCrawlingTasksRequest) (*GetCrawlingTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCrawlingTasks not implemented")
}
func (UnimplementedReviewGrpcServer) SetReviewScore(context.Context, *SetReviewScoreRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetReviewScore not implemented")
}
func (UnimplementedReviewGrpcServer) SetNotExist(context.Context, *SetNotExistRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNotExist not implemented")
}
func (UnimplementedReviewGrpcServer) GetCrawlingTargets(context.Context, *GetCrawlingTargetsRequest) (*GetCrawlingTargetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCrawlingTargets not implemented")
}
func (UnimplementedReviewGrpcServer) SaveCompanyReviews(context.Context, *SaveCompanyReviewsRequest) (*SaveCompanyReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveCompanyReviews not implemented")
}
func (UnimplementedReviewGrpcServer) FinishCrawlingTask(context.Context, *FinishCrawlingTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishCrawlingTask not implemented")
}
func (UnimplementedReviewGrpcServer) mustEmbedUnimplementedReviewGrpcServer() {}

// UnsafeReviewGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewGrpcServer will
// result in compilation errors.
type UnsafeReviewGrpcServer interface {
	mustEmbedUnimplementedReviewGrpcServer()
}

func RegisterReviewGrpcServer(s grpc.ServiceRegistrar, srv ReviewGrpcServer) {
	s.RegisterService(&ReviewGrpc_ServiceDesc, srv)
}

func _ReviewGrpc_GetCrawlingTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCrawlingTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewGrpcServer).GetCrawlingTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.review_service.crawler_grpc.ReviewGrpc/getCrawlingTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewGrpcServer).GetCrawlingTasks(ctx, req.(*GetCrawlingTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewGrpc_SetReviewScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReviewScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewGrpcServer).SetReviewScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.review_service.crawler_grpc.ReviewGrpc/setReviewScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewGrpcServer).SetReviewScore(ctx, req.(*SetReviewScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewGrpc_SetNotExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNotExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewGrpcServer).SetNotExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.review_service.crawler_grpc.ReviewGrpc/setNotExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewGrpcServer).SetNotExist(ctx, req.(*SetNotExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewGrpc_GetCrawlingTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCrawlingTargetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewGrpcServer).GetCrawlingTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.review_service.crawler_grpc.ReviewGrpc/getCrawlingTargets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewGrpcServer).GetCrawlingTargets(ctx, req.(*GetCrawlingTargetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewGrpc_SaveCompanyReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveCompanyReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewGrpcServer).SaveCompanyReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.review_service.crawler_grpc.ReviewGrpc/saveCompanyReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewGrpcServer).SaveCompanyReviews(ctx, req.(*SaveCompanyReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewGrpc_FinishCrawlingTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishCrawlingTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewGrpcServer).FinishCrawlingTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.review_service.crawler_grpc.ReviewGrpc/finishCrawlingTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewGrpcServer).FinishCrawlingTask(ctx, req.(*FinishCrawlingTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReviewGrpc_ServiceDesc is the grpc.ServiceDesc for ReviewGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReviewGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "careerhub.review_service.crawler_grpc.ReviewGrpc",
	HandlerType: (*ReviewGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getCrawlingTasks",
			Handler:    _ReviewGrpc_GetCrawlingTasks_Handler,
		},
		{
			MethodName: "setReviewScore",
			Handler:    _ReviewGrpc_SetReviewScore_Handler,
		},
		{
			MethodName: "setNotExist",
			Handler:    _ReviewGrpc_SetNotExist_Handler,
		},
		{
			MethodName: "getCrawlingTargets",
			Handler:    _ReviewGrpc_GetCrawlingTargets_Handler,
		},
		{
			MethodName: "saveCompanyReviews",
			Handler:    _ReviewGrpc_SaveCompanyReviews_Handler,
		},
		{
			MethodName: "finishCrawlingTask",
			Handler:    _ReviewGrpc_FinishCrawlingTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "careerhub/review_crawler/crawler_grpc/review.proto",
}
