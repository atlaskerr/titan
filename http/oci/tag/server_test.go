package tag_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/atlaskerr/titan/http/oci/tag"
	"github.com/atlaskerr/titan/http/oci/tag/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
)

type serverTestComponents struct {
	server           *tag.Server
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
		undefinedHandler: mock.NewHandler(ctrl),
		responseWriter:   mock.NewResponseWriter(ctrl),
		types: serverTestMatchers{
			request:        gomock.AssignableToTypeOf(new(http.Request)),
			responseWriter: gomock.Any(),
		},
	}
	opts := []tag.ServerOption{
		tag.OptionMetricsCollector(metrics.NewCollector()),
		tag.OptionUndefinedHandler(cmp.undefinedHandler),
	}
	server, err := tag.NewServer(opts...)
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
	cmp.undefinedHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}
