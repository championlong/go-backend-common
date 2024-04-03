package pool

import (
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	hbase_thrift "github.com/championlong/go-backend-common/hbase/gen-go/hbase"
	"sync/atomic"
	"time"
)

var noDeadline = time.Time{}

type Conn struct {
	socketConn thrift.TTransport
	clientConn hbase_thrift.THBaseService
	createdAt  time.Time
	usedAt     uint32 // atomic
	pooled     bool
}

func NewConn(socketConn thrift.TTransport, factory thrift.TProtocolFactory) *Conn {
	cn := &Conn{
		createdAt: time.Now(),
	}
	cn.SetSocketConn(socketConn)
	cn.SetUsedAt(time.Now())
	cn.SetClientConn(socketConn, factory)
	return cn
}

func (cn *Conn) UsedAt() time.Time {
	unix := atomic.LoadUint32(&cn.usedAt)
	return time.Unix(int64(unix), 0)
}

func (cn *Conn) SetUsedAt(tm time.Time) {
	atomic.StoreUint32(&cn.usedAt, uint32(tm.Unix()))
}

func (cn *Conn) SetSocketConn(socketConn thrift.TTransport) {
	cn.socketConn = socketConn
}

func (cn *Conn) SetClientConn(socketConn thrift.TTransport, factory thrift.TProtocolFactory) {
	cn.clientConn = hbase_thrift.NewTHBaseServiceClientFactory(socketConn, factory)
}

func (cn *Conn) SocketConn() thrift.TTransport {
	return cn.socketConn
}

func (cn *Conn) HbaseClient() hbase_thrift.THBaseService {
	return cn.clientConn
}

func (cn *Conn) Close() error {
	return cn.socketConn.Close()
}

func (cn *Conn) deadline(ctx context.Context, timeout time.Duration) time.Time {
	tm := time.Now()
	cn.SetUsedAt(tm)

	if timeout > 0 {
		tm = tm.Add(timeout)
	}

	if ctx != nil {
		deadline, ok := ctx.Deadline()
		if ok {
			if timeout == 0 {
				return deadline
			}
			if deadline.Before(tm) {
				return deadline
			}
			return tm
		}
	}

	if timeout > 0 {
		return tm
	}

	return noDeadline
}
