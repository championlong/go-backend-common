package hbase

import (
	"context"
	"fmt"
	"testing"

	"github.com/championlong/go-backend-common/hbase/gen-go/hbase"
	"github.com/stretchr/testify/assert"
)

var (
	hb = Connect(&Options{
		ThriftHost: "docker-hbase:19090",
	})
	table = []byte("test_table")
)

func Test_PutByRowKey(t *testing.T) {
	put := &hbase.TPut{Row: []byte("row2"), ColumnValues: []*hbase.TColumnValue{{
		Family:    []byte("f1"),
		Qualifier: []byte("q1"),
		Value:     []byte("value2")}}}
	err := hb.Put(context.Background(), table, put)
	assert.NoError(t, err)
}

func Test_PutsByRowKeys(t *testing.T) {
	puts := []*hbase.TPut{{Row: []byte("row1"), ColumnValues: []*hbase.TColumnValue{{
		Family:    []byte("f1"),
		Qualifier: []byte("q1"),
		Value:     []byte("value1")}}}, {Row: []byte("row1"), ColumnValues: []*hbase.TColumnValue{{
		Family:    []byte("f2"),
		Qualifier: []byte("q2"),
		Value:     []byte("value2")}}}}
	err := hb.PutMultiple(context.Background(), table, puts)
	assert.NoError(t, err)
}

func Test_GetByRowKeys(t *testing.T) {
	get := &hbase.TGet{Row: []byte("row2")}
	result, err := hb.Get(context.Background(), table, get)
	fmt.Println("GetMultiple result:")
	fmt.Println(result)
	assert.NoError(t, err)
	for _, column := range result.GetColumnValues() {
		if string(column.Family) == "f1" {
			fmt.Println("11111", string(column.Value))
		}
		if string(column.Family) == "f2" {
			fmt.Println("2222", string(column.Value))
		}
	}
}

func Test_GetsByRowKeys(t *testing.T) {
	gets := []*hbase.TGet{{Row: []byte("row1")}, {Row: []byte("row2")}}
	results, err := hb.GetMultiple(context.Background(), table, gets)
	fmt.Println("GetMultiple result:")
	fmt.Println(results)
	assert.NoError(t, err)
	for _, item := range results {
		for _, column := range item.GetColumnValues() {
			if string(column.Family) == "f1" {
				fmt.Println("11111", string(column.Value))
			}
			if string(column.Family) == "f2" {
				fmt.Println("2222", string(column.Value))
			}
		}
	}
}
