package dao

import (
	"fmt"
	"testing"
)

func Test_sqlCondition(t *testing.T) {
	testSql := "select id from test where %s"
	var conditions []SqlCondition
	conditions = append(conditions, NewClauseSqlCondition("query1 = ?", []interface{}{"1"}))
	conditions = append(conditions, NewSingleSqlCondition("query2", "=", "2"))
	whereClause, value := BuildWhereClause(conditions)
	fmt.Println(fmt.Sprintf(testSql, whereClause))
	fmt.Println(value)
}
