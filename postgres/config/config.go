package config

import (
	"context"
	"net"
)

// Database : 数据库配置
type Database struct {
	Postgres map[string]PostgresCluster
}

// PostgresCluster 集群配置
type PostgresCluster struct {
	// 逻辑分片
	Mod int
	// 分片配置
	Shards []struct {
		// 逻辑分片开始id
		FromLogicalShardMod int
		// 逻辑分片结束id
		ToLogicalShardMod int
		Master            PostgresConfig
		Slaves            []PostgresConfig
	}
}

// PostgresConfig 数据库链接配置
type PostgresConfig struct {
	// 监听的地址
	Address string
	// 监听的端口
	Port string
	// 自定义 Dialer
	Dialer func(ctx context.Context, network, addr string) (net.Conn, error) `json:"-"`
	// 用户名
	User string
	// 密码
	Password string
	// 数据库名字
	Database string
	// 应用程序名称
	ApplicationName string
	// 连接池大小, default : 10 * runtime.NumCPU
	PoolSize int
	// 最小空闲连接的数量
	MinIdleConns int
	// 最大重试次数，默认不重试.
	MaxRetries int
	// 重试时，请求之间最小退避时间(单位毫秒), 默认 250ms; -1 表示关闭退避
	MinRetryBackoffMs int
	// 重试时，请求之间最大退避时间(单位毫秒), 默认 4s；-1 表示关闭退避
	MaxRetryBackoffMs int
	// PG的最大链接时间, 单位秒, pgv10版本有效
	MaxAge int
	// 从连接池获取连接超时时间,单位秒, 默认 30s
	PoolTimeout int
	// idle的超时时间, 单位秒, 默认 5min
	IdleTimeout int
	// idle检查频率, 单位秒, 默认 60s
	IdleCheckFrequency int
	// 连接超时, 默认 4s
	DialTimeout int
	// 读超时（tcp层面），会导致重建链接, 默认 4s
	ReadTimeout int
	// 写超时（tcp层面），会导致重建链接, 默认 4s
	WriteTimeout int
}
