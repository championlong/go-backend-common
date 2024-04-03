package hbase

import (
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/championlong/go-backend-common/hbase/pool"
	"runtime"
	"time"
)

// Options contains database connection options.
type Options struct {
	Dialer  func(context.Context, string, *thrift.TConfiguration) (thrift.TTransport, error)
	Factory func(context.Context, *thrift.TConfiguration) thrift.TProtocolFactory

	ThriftHost string `json:"thrift"`

	ThriftConfig *thrift.TConfiguration

	// Maximum number of socket connections.
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	PoolSize int
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns int
	// Connection age at which client retires (closes) the connection.
	// It is useful with proxies like PgBouncer and HAProxy.
	// Default is to not close aged connections.
	MaxConnAge time.Duration
	// Time for which client waits for free connection if all
	// connections are busy before returning an error.
	// Default is 30 seconds if ReadTimeOut is not defined, otherwise,
	// ReadTimeout + 1 second.
	PoolTimeout time.Duration
	// Amount of time after which client closes idle connections.
	// Should be less than server's timeout.
	// Default is 5 minutes. -1 disables idle timeout check.
	IdleTimeout time.Duration
	// Frequency of idle checks made by idle connections reaper.
	// Default is 1 minute. -1 disables idle connections reaper,
	// but idle connections are still discarded by the client
	// if IdleTimeout is set.
	IdleCheckFrequency time.Duration
}

func (opt *Options) init() {
	if opt.Dialer == nil {
		opt.Dialer = func(ctx context.Context, thriftHost string, config *thrift.TConfiguration) (thrift.TTransport, error) {
			if config == nil {
				config = &thrift.TConfiguration{
					ConnectTimeout: time.Second * 60,
					SocketTimeout:  time.Second * 60,
				}
			}
			tSocket := thrift.NewTSocketConf(thriftHost, config)
			return tSocket, nil
		}
	}

	if opt.Factory == nil {
		opt.Factory = func(ctx context.Context, config *thrift.TConfiguration) thrift.TProtocolFactory {
			if config == nil {
				config = &thrift.TConfiguration{
					TBinaryStrictRead:  thrift.BoolPtr(true),
					TBinaryStrictWrite: thrift.BoolPtr(true),
				}
			}
			return thrift.NewTBinaryProtocolFactoryConf(config)
		}
	}

	if opt.PoolSize == 0 {
		opt.PoolSize = 10 * runtime.NumCPU()
	}

	if opt.PoolTimeout == 0 {
		opt.PoolTimeout = 30 * time.Second
	}

	if opt.IdleTimeout == 0 {
		opt.IdleTimeout = 5 * time.Minute
	}
	if opt.IdleCheckFrequency == 0 {
		opt.IdleCheckFrequency = time.Minute
	}
}

func (opt *Options) getDialer() func(context.Context) (thrift.TTransport, error) {
	return func(ctx context.Context) (thrift.TTransport, error) {
		return opt.Dialer(ctx, opt.ThriftHost, opt.ThriftConfig)
	}
}

func (opt *Options) getFactory() func(context.Context) thrift.TProtocolFactory {
	return func(ctx context.Context) thrift.TProtocolFactory {
		return opt.Factory(ctx, opt.ThriftConfig)
	}
}

func newConnPool(opt *Options) *pool.ConnPool {
	return pool.NewConnPool(&pool.Options{
		Dialer:  opt.getDialer(),
		Factory: opt.getFactory(),
		OnClose: terminateConn,

		PoolSize:           opt.PoolSize,
		MinIdleConns:       opt.MinIdleConns,
		MaxConnAge:         opt.MaxConnAge,
		PoolTimeout:        opt.PoolTimeout,
		IdleTimeout:        opt.IdleTimeout,
		IdleCheckFrequency: opt.IdleCheckFrequency,
	})
}
