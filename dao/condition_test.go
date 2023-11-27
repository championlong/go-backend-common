

import (
	"fmt"
	"testing"
)

func Test_sqlCondition(t *testing.T) {
	testSql := "select id from test where %s"
	var conditions []SqlCondition
	conditions = append(conditions, NewClauseSqlCondition("query1 = ?", []interface{}{"1"}))
	conditions = append(conditions, NewSingleSqlCondition("query2", "=", "2"))
	conditions = append(conditions, NewSqlCondition("query3", "IN", []interface{}{"3"}))
	conditions = append(conditions, SqlCondition{
		Clause: " ( query4 = ? OR query5 = ? ) ",
		Values: []interface{}{"4", "5"},
	})
	whereClause, value := BuildWhereClause(conditions)
	fmt.Println(fmt.Sprintf(testSql, whereClause))
	fmt.Println(value)
}
