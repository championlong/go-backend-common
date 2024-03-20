package main

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"postgres"
	"postgres/config"
)

func main() {
	postgres.InitDatabaseShard(&config.Database{
		Postgres: map[string]config.PostgresCluster{
			"postgres": {
				Mod: 2,
				Shards: []struct {
					FromLogicalShardMod int
					ToLogicalShardMod   int
					Master              config.PostgresConfig
					Slaves              []config.PostgresConfig
				}{
					{
						FromLogicalShardMod: 0,
						ToLogicalShardMod:   0,
						Master: config.PostgresConfig{
							Address:  "127.0.0.1",
							Port:     "5432",
							User:     "postgres",
							Database: "postgres",
						},
						Slaves: []config.PostgresConfig{
							{
								Address:  "127.0.0.1",
								Port:     "5432",
								User:     "postgres",
								Database: "postgres",
							},
						},
					},
					{
						FromLogicalShardMod: 1,
						ToLogicalShardMod:   1,
						Master: config.PostgresConfig{
							Address:  "127.0.0.1",
							Port:     "5432",
							User:     "postgres",
							Database: "postgres",
						},
						Slaves: []config.PostgresConfig{
							{
								Address:  "127.0.0.1",
								Port:     "5432",
								User:     "postgres",
								Database: "postgres",
							},
						},
					},
				},
			},
		},
	})
	var n int
	//_, err := postgres.GetDB("postgres").Slave().QueryOne(pg.Scan(&n), "SELECT 1")
	_, err := postgres.GetDB("postgres").ShardDB(context.Background(), postgres.Shard(1)).QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
