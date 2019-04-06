package service

import (
	"context"
	"time"

	pb "github.com/ianfoo/gaproxy"
	"github.com/ianfoo/gaproxy/session"
	"github.com/sirupsen/logrus"
	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
)

type Service struct {
	log      *logrus.Logger
	sessions session.Manager
}

func NewService() Service {
	return Service{
		log:      logrus.New(),
		sessions: session.NewManager(),
	}
}

func (s Service) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.log.WithFields(logrus.Fields{
		"identity": in.Identity,
	}).Info("login request")

	sn, err := s.sessions.New(session.SessionDefaultLifetime, in.Identity)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"identity": in.Identity,
			"err":      err,
		}).Error("login error")
		return nil, grpcstatus.Error(grpccodes.Unauthenticated, "login error")
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

// Logout a session.
func (s Service) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	if err := s.sessions.Invalidate(in.SessionId); err != nil {
		s.log.WithFields(logrus.Fields{
			"session_id": in.SessionId,
			"err":        err,
		}).Error("logout error")
		return nil, grpcstatus.Error(grpccodes.NotFound, err.Error())
	}
	s.log.WithFields(logrus.Fields{
		"session_id": in.SessionId,
	}).Info("logged out")
	resp := pb.LogoutResponse{
		SessionId: in.SessionId,
	}
	return &resp, nil
}

// CheckSession implements Service.
func (s Service) CheckSession(ctx context.Context, in *pb.CheckSessionRequest) (*pb.CheckSessionResponse, error) {
	session, err := s.sessions.Get(in.SessionId)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"session_id": in.SessionId,
			"err":        err,
		}).Info("checking session")
		return nil, grpcstatus.Error(grpccodes.NotFound, err.Error())
	}
	resp := pb.CheckSessionResponse{
		SessionId: session.ID(),
		IsValid:   session.Valid(),
		ExpiresAt: session.ExpiresAt().Unix(),
	}
	return &resp, nil
}

// Ping implements Service.
func (s Service) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	var resp pb.PingResponse
	resp = pb.PingResponse{
		Status: pb.PingResponse_OK,
	}
	return &resp, nil
}
