// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: feec04999c
// Version Date: Sat Mar 16 17:27:47 UTC 2019

package svc

// This file provides server-side bindings for the HTTP transport.
// It utilizes the transport/http.Server.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"context"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	// This service
	pb "github.com/ianfoo/gaproxy"
)

const contentType = "application/json; charset=utf-8"

var (
	_ = fmt.Sprint
	_ = bytes.Compare
	_ = strconv.Atoi
	_ = httptransport.NewServer
	_ = ioutil.NopCloser
	_ = pb.NewGAProxyClient
	_ = io.Copy
	_ = errors.Wrap
)

// MakeHTTPHandler returns a handler that makes a set of endpoints available
// on predefined paths.
func MakeHTTPHandler(endpoints Endpoints) http.Handler {
	serverOptions := []httptransport.ServerOption{
		httptransport.ServerBefore(headersToContext),
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerAfter(httptransport.SetContentType(contentType)),
	}
	m := mux.NewRouter()

	m.Methods("POST").Path("/session/login").Handler(httptransport.NewServer(
		endpoints.LoginEndpoint,
		DecodeHTTPLoginZeroRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))

	m.Methods("GET").Path("/session/logout/{session_id}").Handler(httptransport.NewServer(
		endpoints.LogoutEndpoint,
		DecodeHTTPLogoutZeroRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))

	m.Methods("GET").Path("/session/check/{session_id}").Handler(httptransport.NewServer(
		endpoints.CheckSessionEndpoint,
		DecodeHTTPCheckSessionZeroRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))

	m.Methods("POST").Path("/query-ga").Handler(httptransport.NewServer(
		endpoints.QueryEndpoint,
		DecodeHTTPQueryZeroRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))

	m.Methods("GET").Path("/ping").Handler(httptransport.NewServer(
		endpoints.PingEndpoint,
		DecodeHTTPPingZeroRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	))
	return m
}

// ErrorEncoder writes the error to the ResponseWriter, by default a content
// type of application/json, a body of json with key "error" and the value
// error.Error(), and a status code of 500. If the error implements Headerer,
// the provided headers will be applied to the response. If the error
// implements json.Marshaler, and the marshaling succeeds, the JSON encoded
// form of the error will be used. If the error implements StatusCoder, the
// provided StatusCode will be used instead of 500.
func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	body, _ := json.Marshal(errorWrapper{Error: err.Error()})
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			body = jsonBody
		}
	}
	w.Header().Set("Content-Type", contentType)
	if headerer, ok := err.(httptransport.Headerer); ok {
		for k := range headerer.Headers() {
			w.Header().Set(k, headerer.Headers().Get(k))
		}
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(httptransport.StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	w.Write(body)
}

type errorWrapper struct {
	Error string `json:"error"`
}

// httpError satisfies the Headerer and StatusCoder interfaces in
// package github.com/go-kit/kit/transport/http.
type httpError struct {
	error
	statusCode int
	headers    map[string][]string
}

func (h httpError) StatusCode() int {
	return h.statusCode
}

func (h httpError) Headers() http.Header {
	return h.headers
}

// Server Decode

// DecodeHTTPLoginZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded login request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPLoginZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req pb.LoginRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.Unmarshal(buf, &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{fmt.Errorf("request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if IdentityLoginStrArr, ok := queryParams["identity"]; ok {
		IdentityLoginStr := IdentityLoginStrArr[0]
		IdentityLogin := IdentityLoginStr
		req.Identity = IdentityLogin
	}

	return &req, err
}

// DecodeHTTPLogoutZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded logout request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPLogoutZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req pb.LogoutRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.Unmarshal(buf, &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{fmt.Errorf("request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	SessionIdLogoutStr := pathParams["session_id"]
	SessionIdLogout := SessionIdLogoutStr
	req.SessionId = SessionIdLogout

	return &req, err
}

// DecodeHTTPCheckSessionZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded checksession request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPCheckSessionZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req pb.CheckSessionRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.Unmarshal(buf, &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{fmt.Errorf("request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	SessionIdCheckSessionStr := pathParams["session_id"]
	SessionIdCheckSession := SessionIdCheckSessionStr
	req.SessionId = SessionIdCheckSession

	return &req, err
}

// DecodeHTTPQueryZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded query request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPQueryZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req pb.QueryRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.Unmarshal(buf, &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{fmt.Errorf("request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if SessionIdQueryStrArr, ok := queryParams["session_id"]; ok {
		SessionIdQueryStr := SessionIdQueryStrArr[0]
		SessionIdQuery := SessionIdQueryStr
		req.SessionId = SessionIdQuery
	}

	if StartDateQueryStrArr, ok := queryParams["start_date"]; ok {
		StartDateQueryStr := StartDateQueryStrArr[0]
		StartDateQuery := StartDateQueryStr
		req.StartDate = StartDateQuery
	}

	if EndDateQueryStrArr, ok := queryParams["end_date"]; ok {
		EndDateQueryStr := EndDateQueryStrArr[0]
		EndDateQuery := EndDateQueryStr
		req.EndDate = EndDateQuery
	}

	if MetricsQueryStrArr, ok := queryParams["metrics"]; ok {
		MetricsQueryStr := MetricsQueryStrArr[0]

		var MetricsQuery []string
		err = json.Unmarshal([]byte(MetricsQueryStr), &MetricsQuery)
		if err != nil {
			MetricsQueryStr = "[" + MetricsQueryStr + "]"
		}
		err = json.Unmarshal([]byte(MetricsQueryStr), &MetricsQuery)
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't decode MetricsQuery from %v", MetricsQueryStr)
		}
		req.Metrics = MetricsQuery
	}

	if DimensionsQueryStrArr, ok := queryParams["dimensions"]; ok {
		DimensionsQueryStr := DimensionsQueryStrArr[0]

		var DimensionsQuery []string
		err = json.Unmarshal([]byte(DimensionsQueryStr), &DimensionsQuery)
		if err != nil {
			DimensionsQueryStr = "[" + DimensionsQueryStr + "]"
		}
		err = json.Unmarshal([]byte(DimensionsQueryStr), &DimensionsQuery)
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't decode DimensionsQuery from %v", DimensionsQueryStr)
		}
		req.Dimensions = DimensionsQuery
	}

	return &req, err
}

// DecodeHTTPPingZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded ping request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPPingZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req pb.PingRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.Unmarshal(buf, &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{fmt.Errorf("request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := mux.Vars(r)
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	return &req, err
}

// EncodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeHTTPGenericResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(response)
}

// Helper functions

func headersToContext(ctx context.Context, r *http.Request) context.Context {
	for k, _ := range r.Header {
		// The key is added both in http format (k) which has had
		// http.CanonicalHeaderKey called on it in transport as well as the
		// strings.ToLower which is the grpc metadata format of the key so
		// that it can be accessed in either format
		ctx = context.WithValue(ctx, k, r.Header.Get(k))
		ctx = context.WithValue(ctx, strings.ToLower(k), r.Header.Get(k))
	}

	// Tune specific change.
	// also add the request url
	ctx = context.WithValue(ctx, "request-url", r.URL.Path)
	ctx = context.WithValue(ctx, "transport", "HTTPJSON")

	return ctx
}
