package handlers

import (
	"context"
	"time"

	pb "github.com/ianfoo/gaproxy"
	"github.com/ianfoo/gaproxy/session"
	"github.com/sirupsen/logrus"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.GAProxyServer {
	return gaproxyService{
		log:      logrus.New(),
		sessions: session.NewManager(),
	}
}

type gaproxyService struct {
	log      *logrus.Logger
	sessions session.Manager
}

// Login implements Service.
func (s gaproxyService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.log.WithFields(logrus.Fields{
		"identity": in.Identity,
	}).Info("login request")

	sn, err := s.sessions.New(session.SessionDefaultLifetime, in.Identity)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"identity": in.Identity,
			"err":      err,
		}).Error("login error")
		return nil, err
	}
	s.log.WithFields(logrus.Fields{
		"session_id":        sn.ID(),
		"expires_at":        sn.ExpiresAt().Format(time.RFC3339),
		"google_auth_token": sn.GoogleAuthToken(),
	}).Info("session created")

	resp := pb.LoginResponse{
		SessionId: sn.ID(),
		ExpiresAt: sn.ExpiresAt().Unix(),
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
