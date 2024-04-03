package hbase

import (
	"context"
	"github.com/championlong/go-backend-common/hbase/gen-go/hbase"
	"github.com/championlong/go-backend-common/hbase/pool"
)

type Hbase struct {
	ctx  context.Context
	opt  *Options
	pool pool.Pooler
}

func Connect(opt *Options) Hbase {
	opt.init()
	return Hbase{
		ctx:  context.Background(),
		opt:  opt,
		pool: newConnPool(opt),
	}
}

// Exists Test for the existence of columns in the table, as specified in the TGet.
func (db *Hbase) Exists(ctx context.Context, table []byte, tget *hbase.TGet) (_r bool, _err error) {
	var exist bool
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		exist, err = cn.HbaseClient().Exists(ctx, table, tget)
		return err
	})
	return exist, err
}

// ExistsAll Test for the existence of columns in the table, as specified by the TGets.
func (db *Hbase) ExistsAll(ctx context.Context, table []byte, tget []*hbase.TGet) (_r []bool, _err error) {
	var exists []bool
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		exists, err = cn.HbaseClient().ExistsAll(ctx, table, tget)
		return err
	})
	return exists, err
}

// GetMultiple Method for getting multiple rows.
func (db *Hbase) GetMultiple(ctx context.Context, table []byte, contents []*hbase.TGet) ([]*hbase.TResult_, error) {
	var cells []*hbase.TResult_
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		cells, err = cn.HbaseClient().GetMultiple(ctx, table, contents)
		return err
	})
	return cells, err
}

// Put Commit a TPut to a table.
func (db *Hbase) Put(ctx context.Context, table []byte, contents *hbase.TPut) error {
	err := db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err := cn.HbaseClient().Put(ctx, table, contents)
		return err
	})
	return err
}

// PutMultiple Commit a List of Puts to the table.
func (db *Hbase) PutMultiple(ctx context.Context, table []byte, contents []*hbase.TPut) error {
	err := db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err := cn.HbaseClient().PutMultiple(ctx, table, contents)
		return err
	})
	return err
}

func (db *Hbase) Get(ctx context.Context, table []byte, contents *hbase.TGet) (*hbase.TResult_, error) {
	var cell *hbase.TResult_
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		cell, err = cn.HbaseClient().Get(ctx, table, contents)
		return err
	})
	return cell, err
}

// CheckAndPut Atomically checks if a row/family/qualifier value matches the expected value. If it does, it adds the TPut.
func (db *Hbase) CheckAndPut(ctx context.Context, table []byte, row []byte, family []byte, qualifier []byte, value []byte, tput *hbase.TPut) (_r bool, _err error) {
	var result bool
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().CheckAndPut(ctx, table, row, family, qualifier, value, tput)
		return err
	})
	return result, err
}

// DeleteSingle Deletes as specified by the TDelete.
func (db *Hbase) DeleteSingle(ctx context.Context, table []byte, tdelete *hbase.TDelete) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().DeleteSingle(ctx, table, tdelete)
		return err
	})
	return err
}

// DeleteMultiple Bulk commit a List of TDeletes to the table.
func (db *Hbase) DeleteMultiple(ctx context.Context, table []byte, tdeletes []*hbase.TDelete) (_r []*hbase.TDelete, _err error) {
	var cell []*hbase.TDelete
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		cell, err = cn.HbaseClient().DeleteMultiple(ctx, table, tdeletes)
		return err
	})
	return cell, err
}

// CheckAndDelete Atomically checks if a row/family/qualifier value matches the expected value. If it does, it adds the delete.
func (db *Hbase) CheckAndDelete(ctx context.Context, table []byte, row []byte, family []byte, qualifier []byte, value []byte, tdelete *hbase.TDelete) (_r bool, _err error) {
	var result bool
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().CheckAndDelete(ctx, table, row, family, qualifier, value, tdelete)
		return err
	})
	return result, err
}

func (db *Hbase) Increment(ctx context.Context, table []byte, tincrement *hbase.TIncrement) (_r *hbase.TResult_, _err error) {
	var result *hbase.TResult_
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().Increment(ctx, table, tincrement)
		return err
	})
	return result, err
}

func (db *Hbase) Append(ctx context.Context, table []byte, tappend *hbase.TAppend) (_r *hbase.TResult_, _err error) {
	var result *hbase.TResult_
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().Append(ctx, table, tappend)
		return err
	})
	return result, err
}

// OpenScanner a Scanner for the provided TScan object.
func (db *Hbase) OpenScanner(ctx context.Context, table []byte, tscan *hbase.TScan) (_r int32, _err error) {
	var result int32
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().OpenScanner(ctx, table, tscan)
		return err
	})
	return result, err
}

// GetScannerRows Grabs multiple rows from a Scanner.
func (db *Hbase) GetScannerRows(ctx context.Context, scannerId int32, numRows int32) (_r []*hbase.TResult_, _err error) {
	var result []*hbase.TResult_
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetScannerRows(ctx, scannerId, numRows)
		return err
	})
	return result, err
}

// CloseScanner Should be called to free server side resources timely.
func (db *Hbase) CloseScanner(ctx context.Context, scannerId int32) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().CloseScanner(ctx, scannerId)
		return err
	})
	return err
}

// MutateRow performs multiple mutations atomically on a single row.
func (db *Hbase) MutateRow(ctx context.Context, table []byte, trowMutations *hbase.TRowMutations) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().MutateRow(ctx, table, trowMutations)
		return err
	})
	return err
}

// GetScannerResults results for the provided TScan object.
func (db *Hbase) GetScannerResults(ctx context.Context, table []byte, tscan *hbase.TScan, numRows int32) (_r []*hbase.TResult_, _err error) {
	var result []*hbase.TResult_
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetScannerResults(ctx, table, tscan, numRows)
		return err
	})
	return result, err
}

// GetRegionLocation Given a table and a row get the location of the region that would contain the given row key.
func (db *Hbase) GetRegionLocation(ctx context.Context, table []byte, row []byte, reload bool) (_r *hbase.THRegionLocation, _err error) {
	var result *hbase.THRegionLocation
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetRegionLocation(ctx, table, row, reload)
		return err
	})
	return result, err
}

// GetAllRegionLocations all of the region locations for a given table.
func (db *Hbase) GetAllRegionLocations(ctx context.Context, table []byte) (_r []*hbase.THRegionLocation, _err error) {
	var result []*hbase.THRegionLocation
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetAllRegionLocations(ctx, table)
		return err
	})
	return result, err
}

// CheckAndMutate Atomically checks if a row/family/qualifier value matches the expected value. If it does, it mutates the row.
func (db *Hbase) CheckAndMutate(ctx context.Context, table []byte, row []byte, family []byte, qualifier []byte, compareOperator hbase.TCompareOperator, value []byte,
	rowMutations *hbase.TRowMutations) (_r bool, _err error) {
	var result bool
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().CheckAndMutate(ctx, table, row, family, qualifier, compareOperator, value, rowMutations)
		return err
	})
	return result, err
}

// GetTableDescriptor a table descriptor.
func (db *Hbase) GetTableDescriptor(ctx context.Context, table *hbase.TTableName) (_r *hbase.TTableDescriptor, _err error) {
	var result *hbase.TTableDescriptor
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetTableDescriptor(ctx, table)
		return err
	})
	return result, err
}

// GetTableDescriptors table descriptors of tables.
func (db *Hbase) GetTableDescriptors(ctx context.Context, tables []*hbase.TTableName) (_r []*hbase.TTableDescriptor, _err error) {
	var result []*hbase.TTableDescriptor
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetTableDescriptors(ctx, tables)
		return err
	})
	return result, err
}

func (db *Hbase) TableExists(ctx context.Context, tableName *hbase.TTableName) (_r bool, _err error) {
	var result bool
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().TableExists(ctx, tableName)
		return err
	})
	return result, err
}

// GetTableDescriptorsByPattern table descriptors of tables that match the given pattern
func (db *Hbase) GetTableDescriptorsByPattern(ctx context.Context, regex string, includeSysTables bool) (_r []*hbase.TTableDescriptor, _err error) {
	var result []*hbase.TTableDescriptor
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetTableDescriptorsByPattern(ctx, regex, includeSysTables)
		return err
	})
	return result, err
}

// GetTableDescriptorsByNamespace table descriptors of tables in the given namespace
func (db *Hbase) GetTableDescriptorsByNamespace(ctx context.Context, name string) (_r []*hbase.TTableDescriptor, _err error) {
	var result []*hbase.TTableDescriptor
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetTableDescriptorsByNamespace(ctx, name)
		return err
	})
	return result, err
}

// GetTableNamesByPattern table names of tables that match the given pattern
func (db *Hbase) GetTableNamesByPattern(ctx context.Context, regex string, includeSysTables bool) (_r []*hbase.TTableName, _err error) {
	var result []*hbase.TTableName
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetTableNamesByPattern(ctx, regex, includeSysTables)
		return err
	})
	return result, err
}

// GetTableNamesByNamespace table names of tables in the given namespace
func (db *Hbase) GetTableNamesByNamespace(ctx context.Context, name string) (_r []*hbase.TTableName, _err error) {
	var result []*hbase.TTableName
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetTableNamesByNamespace(ctx, name)
		return err
	})
	return result, err
}

// CreateTable Creates a new table with an initial set of empty regions defined by the specified split keys.
func (db *Hbase) CreateTable(ctx context.Context, desc *hbase.TTableDescriptor, splitKeys [][]byte) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().CreateTable(ctx, desc, splitKeys)
		return err
	})
	return err
}

// DeleteTable Deletes a table. Synchronous operation.
func (db *Hbase) DeleteTable(ctx context.Context, tableName *hbase.TTableName) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().DeleteTable(ctx, tableName)
		return err
	})
	return err
}

// TruncateTable Truncate a table. Synchronous operation.
func (db *Hbase) TruncateTable(ctx context.Context, tableName *hbase.TTableName, preserveSplits bool) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().TruncateTable(ctx, tableName, preserveSplits)
		return err
	})
	return err
}

// EnableTable Enalbe a table
func (db *Hbase) EnableTable(ctx context.Context, tableName *hbase.TTableName) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().EnableTable(ctx, tableName)
		return err
	})
	return err
}

// DisableTable Disable a table
func (db *Hbase) DisableTable(ctx context.Context, tableName *hbase.TTableName) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().DisableTable(ctx, tableName)
		return err
	})
	return err
}

func (db *Hbase) IsTableEnabled(ctx context.Context, tableName *hbase.TTableName) (_r bool, _err error) {
	var result bool
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().IsTableEnabled(ctx, tableName)
		return err
	})
	return result, err
}

func (db *Hbase) IsTableDisabled(ctx context.Context, tableName *hbase.TTableName) (_r bool, _err error) {
	var result bool
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().IsTableDisabled(ctx, tableName)
		return err
	})
	return result, err
}

func (db *Hbase) IsTableAvailable(ctx context.Context, tableName *hbase.TTableName) (_r bool, _err error) {
	var result bool
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().IsTableAvailable(ctx, tableName)
		return err
	})
	return result, err
}

func (db *Hbase) IsTableAvailableWithSplit(ctx context.Context, tableName *hbase.TTableName, splitKeys [][]byte) (_r bool, _err error) {
	var result bool
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().IsTableAvailableWithSplit(ctx, tableName, splitKeys)
		return err
	})
	return result, err
}

// AddColumnFamily Add a column family to an existing table. Synchronous operation.
func (db *Hbase) AddColumnFamily(ctx context.Context, tableName *hbase.TTableName, column *hbase.TColumnFamilyDescriptor) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().AddColumnFamily(ctx, tableName, column)
		return err
	})
	return err
}

// DeleteColumnFamily Delete a column family from a table. Synchronous operation.
func (db *Hbase) DeleteColumnFamily(ctx context.Context, tableName *hbase.TTableName, column []byte) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().DeleteColumnFamily(ctx, tableName, column)
		return err
	})
	return err
}

// ModifyColumnFamily Modify an existing column family on a table. Synchronous operation.
func (db *Hbase) ModifyColumnFamily(ctx context.Context, tableName *hbase.TTableName, column *hbase.TColumnFamilyDescriptor) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().ModifyColumnFamily(ctx, tableName, column)
		return err
	})
	return err
}

// ModifyTable Modify an existing table
func (db *Hbase) ModifyTable(ctx context.Context, desc *hbase.TTableDescriptor) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().ModifyTable(ctx, desc)
		return err
	})
	return err
}

// CreateNamespace Create a new namespace. Blocks until namespace has been successfully created or an exception is thrown
func (db *Hbase) CreateNamespace(ctx context.Context, namespaceDesc *hbase.TNamespaceDescriptor) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().CreateNamespace(ctx, namespaceDesc)
		return err
	})
	return err
}

// ModifyNamespace Modify an existing namespace.  Blocks until namespace has been successfully modified or an exception is thrown
func (db *Hbase) ModifyNamespace(ctx context.Context, namespaceDesc *hbase.TNamespaceDescriptor) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().ModifyNamespace(ctx, namespaceDesc)
		return err
	})
	return err
}

// DeleteNamespace Delete an existing namespace. Only empty namespaces (no tables) can be removed.
func (db *Hbase) DeleteNamespace(ctx context.Context, name string) (_err error) {
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		err = cn.HbaseClient().DeleteNamespace(ctx, name)
		return err
	})
	return err
}

// GetNamespaceDescriptor a namespace descriptor by name.
func (db *Hbase) GetNamespaceDescriptor(ctx context.Context, name string) (_r *hbase.TNamespaceDescriptor, _err error) {
	var result *hbase.TNamespaceDescriptor
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetNamespaceDescriptor(ctx, name)
		return err
	})
	return result, err
}

// ListNamespaceDescriptors @return all namespaces
func (db *Hbase) ListNamespaceDescriptors(ctx context.Context) (_r []*hbase.TNamespaceDescriptor, _err error) {
	var result []*hbase.TNamespaceDescriptor
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().ListNamespaceDescriptors(ctx)
		return err
	})
	return result, err
}

// ListNamespaces @return all namespace names
func (db *Hbase) ListNamespaces(ctx context.Context) (_r []string, _err error) {
	var result []string
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().ListNamespaces(ctx)
		return err
	})
	return result, err
}

// GetThriftServerType the type of this thrift server.
func (db *Hbase) GetThriftServerType(ctx context.Context) (_r hbase.TThriftServerType, _err error) {
	var result hbase.TThriftServerType
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetThriftServerType(ctx)
		return err
	})
	return result, err
}

// GetClusterId Returns the cluster ID for this cluster.
func (db *Hbase) GetClusterId(ctx context.Context) (_r string, _err error) {
	var result string
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetClusterId(ctx)
		return err
	})
	return result, err
}

// GetSlowLogResponses Retrieves online slow RPC logs from the provided list of RegionServers
func (db *Hbase) GetSlowLogResponses(ctx context.Context, serverNames []*hbase.TServerName, logQueryFilter *hbase.TLogQueryFilter) (_r []*hbase.TOnlineLogRecord, _err error) {
	var result []*hbase.TOnlineLogRecord
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().GetSlowLogResponses(ctx, serverNames, logQueryFilter)
		return err
	})
	return result, err
}

// ClearSlowLogResponses Clears online slow/large RPC logs from the provided list of RegionServers
func (db *Hbase) ClearSlowLogResponses(ctx context.Context, serverNames []*hbase.TServerName) (_r []bool, _err error) {
	var result []bool
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().ClearSlowLogResponses(ctx, serverNames)
		return err
	})
	return result, err
}

// Grant permissions in table or namespace level.
func (db *Hbase) Grant(ctx context.Context, info *hbase.TAccessControlEntity) (_r bool, _err error) {
	var result bool
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().Grant(ctx, info)
		return err
	})
	return result, err
}

// Revoke permissions in table or namespace level.
func (db *Hbase) Revoke(ctx context.Context, info *hbase.TAccessControlEntity) (_r bool, _err error) {
	var result bool
	var err error
	err = db.withConn(ctx, func(ctx context.Context, cn *pool.Conn) error {
		result, err = cn.HbaseClient().Revoke(ctx, info)
		return err
	})
	return result, err
}

func (db *Hbase) withConn(
	ctx context.Context, fn func(context.Context, *pool.Conn) error,
) error {
	cn, err := db.pool.Get(ctx)
	if err != nil {
		return err
	}
	if !cn.SocketConn().IsOpen() {
		err = cn.SocketConn().Open()
		if err != nil {
			return err
		}
	}
	var fnDone chan struct{}
	if ctx != nil && ctx.Done() != nil {
		fnDone = make(chan struct{})
		go func() {
			select {
			case <-fnDone: // fn has finished, skip cancel
			case <-ctx.Done():
				// Signal end of conn use.
				fnDone <- struct{}{}
			}
		}()
	}

	defer func() {
		if fnDone == nil {
			db.releaseConn(ctx, cn, err)
			return
		}

		select {
		case <-fnDone: // wait for cancel to finish request
			// Looks like the canceled connection must be always removed from the pool.
			db.pool.Remove(ctx, cn, err)
		case fnDone <- struct{}{}: // signal fn finish, skip cancel goroutine
			db.releaseConn(ctx, cn, err)
		}
	}()

	err = fn(ctx, cn)
	return err
}

func (db *Hbase) releaseConn(ctx context.Context, cn *pool.Conn, err error) {
	if bad := isBadConn(err); bad {
		db.pool.Remove(ctx, cn, err)
	} else {
		db.pool.Put(ctx, cn)
	}
}

func isBadConn(err error) bool {
	if err == nil {
		return false
	}
	return true
}

func terminateConn(conn *pool.Conn) error {
	return nil
}
