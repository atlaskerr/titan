package server

import (
	"context"
	"fmt"

	"github.com/atlaskerr/titan/http/live"
	"github.com/atlaskerr/titan/http/metrics"
	"github.com/atlaskerr/titan/http/oci"
	"github.com/atlaskerr/titan/http/ready"
	"github.com/atlaskerr/titan/http/titan"
	"github.com/atlaskerr/titan/http/unknown"
	titanMetrics "github.com/atlaskerr/titan/metrics"
)

type component func(*service) error

type service struct {
	collector *titanMetrics.Collector
	handlers  handlers
}

type handlers struct {
	titan   *titan.Server
	oci     *oci.Server
	metrics *metrics.Server
	live    *live.Server
	ready   *ready.Server
	unknown *unknown.Server
}

func (s *service) shutdown(ctx context.Context) error {
	return nil
}

func newService() (*service, error) {
	s := &service{}
	// TODO(atlaskerr): Find a cleaner way to do this. Eventually there will be
	// need to start and stop components running in goroutines and whatnot.
	components := []component{
		collector,
		unknownHandler,
		liveHandler,
		readyHandler,
		ociHandler,
		titanHandler,
	}
	for _, component := range components {
		err := component(s)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize component: %s", err)
		}
	}
	return s, nil
}