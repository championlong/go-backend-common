package hbase

import (
	"github.com/championlong/go-backend-common/hbase/pool"
)

type baseHbase struct {
	opt  *Options
	pool pool.Pooler
}

func newHbase(opt Options) {
	opt.init()

	return
}

func terminateConn(conn *pool.Conn) error {

}
