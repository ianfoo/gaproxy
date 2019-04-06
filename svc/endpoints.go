// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: feec04999c
// Version Date: Sat Mar 16 17:27:47 UTC 2019

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	pb "github.com/ianfoo/gaproxy"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	LoginEndpoint        endpoint.Endpoint
	LogoutEndpoint       endpoint.Endpoint
	CheckSessionEndpoint endpoint.Endpoint
	QueryEndpoint        endpoint.Endpoint
	PingEndpoint         endpoint.Endpoint
}

// Endpoints

func (e Endpoints) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	response, err := e.LoginEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.LoginResponse), nil
}

func (e Endpoints) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	response, err := e.LogoutEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.LogoutResponse), nil
}

func (e Endpoints) CheckSession(ctx context.Context, in *pb.CheckSessionRequest) (*pb.CheckSessionResponse, error) {
	response, err := e.CheckSessionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.CheckSessionResponse), nil
}

func (e Endpoints) Query(ctx context.Context, in *pb.QueryRequest) (*pb.QueryResponse, error) {
	response, err := e.QueryEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.QueryResponse), nil
}

func (e Endpoints) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	response, err := e.PingEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.PingResponse), nil
}

// Make Endpoints

func MakeLoginEndpoint(s pb.GAProxyServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.LoginRequest)
		v, err := s.Login(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeLogoutEndpoint(s pb.GAProxyServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.LogoutRequest)
		v, err := s.Logout(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeCheckSessionEndpoint(s pb.GAProxyServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.CheckSessionRequest)
		v, err := s.CheckSession(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeQueryEndpoint(s pb.GAProxyServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.QueryRequest)
		v, err := s.Query(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakePingEndpoint(s pb.GAProxyServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.PingRequest)
		v, err := s.Ping(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"Login":        struct{}{},
		"Logout":       struct{}{},
		"CheckSession": struct{}{},
		"Query":        struct{}{},
		"Ping":         struct{}{},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc, _ := range included {
		if inc == "Login" {
			e.LoginEndpoint = middleware(e.LoginEndpoint)
		}
		if inc == "Logout" {
			e.LogoutEndpoint = middleware(e.LogoutEndpoint)
		}
		if inc == "CheckSession" {
			e.CheckSessionEndpoint = middleware(e.CheckSessionEndpoint)
		}
		if inc == "Query" {
			e.QueryEndpoint = middleware(e.QueryEndpoint)
		}
		if inc == "Ping" {
			e.PingEndpoint = middleware(e.PingEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"Login":        struct{}{},
		"Logout":       struct{}{},
		"CheckSession": struct{}{},
		"Query":        struct{}{},
		"Ping":         struct{}{},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc, _ := range included {
		if inc == "Login" {
			e.LoginEndpoint = middleware("Login", e.LoginEndpoint)
		}
		if inc == "Logout" {
			e.LogoutEndpoint = middleware("Logout", e.LogoutEndpoint)
		}
		if inc == "CheckSession" {
			e.CheckSessionEndpoint = middleware("CheckSession", e.CheckSessionEndpoint)
		}
		if inc == "Query" {
			e.QueryEndpoint = middleware("Query", e.QueryEndpoint)
		}
		if inc == "Ping" {
			e.PingEndpoint = middleware("Ping", e.PingEndpoint)
		}
	}
}
