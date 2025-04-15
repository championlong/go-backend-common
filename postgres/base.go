package postgres

import (
	"fmt"
	"time"

	"github.com/championlong/go-backend-common/postgres/config"
	"github.com/go-pg/pg/v10"
)

type DatabaseGroup struct {
	master    *pg.DB
	slaves    []*pg.DB
	numSlaves int
	idxRobin  int
}

type dbGroup interface {
	Slave() *pg.DB
	Master() *pg.DB
}

//Slave 返回Slave
func (g *DatabaseGroup) Slave() *pg.DB {
	s := g.slaves[g.idxRobin]
	g.idxRobin = (g.idxRobin + 1) % g.numSlaves
	return s
}

// Master 返回Master
func (g *DatabaseGroup) Master() *pg.DB {
	return g.master
}

func newDatabaseGroup(master config.PostgresConfig, slaves []config.PostgresConfig) *DatabaseGroup {
	e := &DatabaseGroup{
		master: newDB(master),
	}
	e.slaves = make([]*pg.DB, len(slaves))
	for i, slave := range slaves {
		e.slaves[i] = newDB(slave)
	}
	e.numSlaves = len(e.slaves)
	return e
}

func newDB(pgConfig config.PostgresConfig) *pg.DB {
	opts := configToPGOptions(pgConfig)
	db := pg.Connect(opts)
	//go func() {
	//	for {
	//		stats := db.PoolStats()
	//	}
	//}()
	return db
}

func configToPGOptions(pgConfig config.PostgresConfig) *pg.Options {
	readTimeout := 4 * time.Second
	if pgConfig.ReadTimeout > 0 {
		readTimeout = time.Duration(pgConfig.ReadTimeout) * time.Second
	}

	writeTimeout := 4 * time.Second
	if pgConfig.WriteTimeout > 0 {
		writeTimeout = time.Duration(pgConfig.WriteTimeout) * time.Second
	}

	minRetryBackoff := time.Duration(-1)
	maxRetryBackoff := time.Duration(-1)
	if pgConfig.MaxRetries > 0 {
		if pgConfig.MinRetryBackoffMs > 0 {
			minRetryBackoff = time.Duration(pgConfig.MinRetryBackoffMs) * time.Millisecond
		}
		if pgConfig.MaxRetryBackoffMs > 0 {
			minRetryBackoff = time.Duration(pgConfig.MaxRetryBackoffMs) * time.Millisecond
		}
	}

	opts := &pg.Options{
		Addr:                  fmt.Sprintf("%s:%s", pgConfig.Address, pgConfig.Port),
		Dialer:                pgConfig.Dialer,
		User:                  pgConfig.User,
		Password:              pgConfig.Password,
		Database:              pgConfig.Database,
		ApplicationName:       pgConfig.ApplicationName,
		PoolSize:              pgConfig.PoolSize,
		MinIdleConns:          pgConfig.MinIdleConns,
		MaxRetries:            pgConfig.MaxRetries,
		MinRetryBackoff:       minRetryBackoff,
		MaxRetryBackoff:       maxRetryBackoff,
		MaxConnAge:            time.Duration(pgConfig.MaxAge) * time.Second,
		PoolTimeout:           time.Duration(pgConfig.PoolTimeout) * time.Second,
		ReadTimeout:           readTimeout,
		WriteTimeout:          writeTimeout,
		DialTimeout:           time.Duration(pgConfig.DialTimeout),
		IdleTimeout:           time.Duration(pgConfig.IdleTimeout) * time.Second,
		IdleCheckFrequency:    time.Duration(pgConfig.IdleCheckFrequency) * time.Second,
		RetryStatementTimeout: false,
	}
	return opts
}
