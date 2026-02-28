package lib_ws

import (
	"time"

	"github.com/hdget/sdk/common/types"
)

type Option func(s *baseServer)

var (
	defaultGracefulShutdownWaitTime = 10 * time.Second
)

func WithGracefulShutdownWaitTime(waitTime time.Duration) Option {
	return func(s *baseServer) {
		s.gracefulShutdownWaitTime = waitTime
	}
}

func WithProviders(providers ...types.Provider) Option {
	return func(s *baseServer) {
		for _, p := range providers {
			s.providers[p.GetCapability().Category] = p
		}
	}
}
