package handlers

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	pb "github.com/ianfoo/gaproxy"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.GAProxyServer {
	return gaproxyService{}
}

type gaproxyService struct{}

// Login implements Service.
func (s gaproxyService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	var resp pb.LoginResponse
	log.Printf("login request for identity: %s", in.Identity)

	sessionID := fmt.Sprintf("%d-%d", rand.Intn(100000), rand.Intn(100))
	resp = pb.LoginResponse{
		SessionId: sessionID,
		ExpiresAt: time.Now().Add(time.Duration(30*24) * time.Hour).Unix(),
	}
	return &resp, nil
}

// Query implements Service.
func (s gaproxyService) Query(ctx context.Context, in *pb.QueryRequest) (*pb.QueryResponse, error) {
	var resp pb.QueryResponse
	resp = pb.QueryResponse{
		QueryResults: "Not implemented yet",
	}
	return &resp, nil
}
