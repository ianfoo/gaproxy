package handlers

import (
	"context"

	pb "github.com/ianfoo/gaproxy"
	"github.com/ianfoo/gaproxy/service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.GAProxyServer {
	return gaproxyService{
		Service: service.NewService(),
	}
}

type gaproxyService struct {
	// Actual functionality is implemented in service.Service to keep this
	// generated file simple.
	service.Service
}

// Login implements Service.
func (s gaproxyService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return s.Service.Login(ctx, in)
}

// Logout implements Service.
func (s gaproxyService) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return s.Service.Logout(ctx, in)
}

// CheckSession implements Service.
func (s gaproxyService) CheckSession(ctx context.Context, in *pb.CheckSessionRequest) (*pb.CheckSessionResponse, error) {
	return s.Service.CheckSession(ctx, in)
}

// Query implements Service.
func (s gaproxyService) Query(ctx context.Context, in *pb.QueryRequest) (*pb.QueryResponse, error) {
	var resp pb.QueryResponse
	resp = pb.QueryResponse{
		QueryResults: "Not implemented yet",
	}
	return &resp, nil
}

// Ping implements Service.
func (s gaproxyService) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	return s.Service.Ping(ctx, in)
}
