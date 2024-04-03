package postgres

import (
	"context"
	"fmt"
	"sync"

	"github.com/championlong/go-backend-common/postgres/config"
	"github.com/go-pg/pg/v10"
)

var DatabaseShardInstance DatabaseShard

// DatabaseShard 数据库分片
type DatabaseShard map[string]*ShardDatabaseGroup

var pgOnce sync.Once

type ShardDatabaseGroup struct {
	Mod                    int
	shards                 []*DatabaseGroup
	numShard               int
	numLogicShardEachShard int
}

type dbOptions struct {
	Master  bool
	ShardID int64
}

type Option func(*dbOptions)

func Master() Option {
	return func(opts *dbOptions) {
		opts.Master = true
	}
}

func Shard(shardID int64) Option {
	return func(opts *dbOptions) {
		opts.ShardID = shardID
	}
}

// InitDatabaseShard DB分片初始化
func InitDatabaseShard(config *config.Database) {
	pgOnce.Do(func() {
		DatabaseShardInstance = newDatabaseShard(config)
	})
}

func newDatabaseShard(dbConfig *config.Database) DatabaseShard {
	res := make(DatabaseShard, len(dbConfig.Postgres))
	for name, conf := range dbConfig.Postgres {
		res[name] = newShardDatabaseGroup(conf)
	}
	return res
}

// GetDB 获取对应数据库
func GetDB(key string) *ShardDatabaseGroup {
	return DatabaseShardInstance[key]
}

func newShardDatabaseGroup(c config.PostgresCluster) *ShardDatabaseGroup {
	group := &ShardDatabaseGroup{
		Mod:      c.Mod,
		numShard: len(c.Shards),
		shards:   make([]*DatabaseGroup, len(c.Shards)),
	}
	if group.Mod%group.numShard != 0 {
		panic(fmt.Sprintf("Shard config not right %+v", c))
	}
	group.numLogicShardEachShard = group.Mod / group.numShard
	fromLogicalShardMod := 0
	toLogicalShardMod := group.numLogicShardEachShard - 1
	for i, shard := range c.Shards {
		if shard.FromLogicalShardMod != fromLogicalShardMod || shard.ToLogicalShardMod != toLogicalShardMod {
			panic(fmt.Sprintf("Shard config not correct %+v %d", c, i))
		}
		group.shards[i] = newDatabaseGroup(shard.Master, shard.Slaves)
		fromLogicalShardMod += group.numLogicShardEachShard
		toLogicalShardMod += group.numLogicShardEachShard
	}
	return group
}

func (g *ShardDatabaseGroup) Slave() *pg.DB {
	return g.shards[0].Slave()
}

func (g *ShardDatabaseGroup) Master() *pg.DB {
	return g.shards[0].Master()
}

func (g *ShardDatabaseGroup) ShardDB(ctx context.Context, opts ...Option) (db *pg.DB) {
	o := &dbOptions{
		ShardID: -1,
	}
	for _, opt := range opts {
		opt(o)
	}

	var shard dbGroup
	if o.ShardID >= 0 {
		idx := int(o.ShardID%int64(g.Mod)) / g.numLogicShardEachShard
		shard = g.shards[idx]
	} else {
		shard = g.shards[0]
	}
	if o.Master {
		db = shard.Master()
	} else {
		db = shard.Slave()
	}
	db = db.WithContext(ctx)
	return db
}
