package ready_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/atlaskerr/titan/http/ready"
	"github.com/atlaskerr/titan/http/ready/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
)

type serverTestComponents struct {
	server           *ready.Server
	core             *mock.Readiness
	undefinedHandler *mock.Handler
	responseWriter   *mock.ResponseWriter
	types            serverTestMatchers
}

type serverTestMatchers struct {
	request        gomock.Matcher
	responseWriter gomock.Matcher
}

func setupServerTestComponents(t *testing.T,
	ctrl *gomock.Controller) serverTestComponents {
	t.Helper()
	cmp := serverTestComponents{
		core:             mock.NewReadiness(ctrl),
		undefinedHandler: mock.NewHandler(ctrl),
		responseWriter:   mock.NewResponseWriter(ctrl),
		types: serverTestMatchers{
			request:        gomock.AssignableToTypeOf(new(http.Request)),
			responseWriter: gomock.Any(),
		},
	}
	opts := []ready.ServerOption{
		ready.OptionCore(cmp.core),
		ready.OptionMetricsCollector(metrics.NewCollector()),
		ready.OptionUndefinedHandler(cmp.undefinedHandler),
	}
	server, err := ready.NewServer(opts...)
	if err != nil {
		t.Fatal(err)
	}
	cmp.server = server
	return cmp
}

func TestServerUndefinedPath(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/unknown/path",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.undefinedHandler.EXPECT().ServeHTTP(
		cmp.types.responseWriter, cmp.types.request,
	)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerReadyOK(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	gomock.InOrder(
		cmp.core.EXPECT().Ready().Return(true),
		cmp.responseWriter.EXPECT().WriteHeader(200),
	)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerReadyNotOK(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	gomock.InOrder(
		cmp.core.EXPECT().Ready().Return(false),
		cmp.responseWriter.EXPECT().WriteHeader(404),
	)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}
