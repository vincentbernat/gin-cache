package cache

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	logger Logger

	getCacheStrategyByRequest GetCacheStrategyByRequest

	hitCacheCallback OnHitCacheCallback

	singleFlightForgetTimeout time.Duration
}

type Option func(c *Config)

// WithLogger user can record logs by the logger
func WithLogger(l Logger) Option {
	return func(c *Config) {
		if l != nil {
			c.logger = l
		}
	}
}

type Logger interface {
	Errorf(string, ...interface{})
}

// Discard the default logger that discard all logs of gin-cache
type Discard struct {
}

func (l Discard) Errorf(string, ...interface{}) {

}

func WithCacheStrategyByRequest(getGetCacheStrategyByRequest GetCacheStrategyByRequest) Option {
	return func(c *Config) {
		if getGetCacheStrategyByRequest != nil {
			c.getCacheStrategyByRequest = getGetCacheStrategyByRequest
		}
	}
}

type OnHitCacheCallback func(c *gin.Context)

// WithOnHitCache will be called when cache hit.
func WithOnHitCache(cb OnHitCacheCallback) Option {
	return func(c *Config) {
		if cb != nil {
			c.hitCacheCallback = cb
		}
	}
}

var defaultHitCacheCallback = func(c *gin.Context) {}

// WithSingleFlightForgetTimeout to reduce the impact of long tail requests. when request in the singleflight,
// after the forget timeout, singleflight.Forget will be called
func WithSingleFlightForgetTimeout(forgetTimeout time.Duration) Option {
	return func(c *Config) {
		if forgetTimeout > 0 {
			c.singleFlightForgetTimeout = forgetTimeout
		}
	}
}
