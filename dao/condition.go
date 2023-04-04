package dao

import (
	"fmt"
	"strings"
)

type SqlCondition struct {
	Key      string
	Operator string
	Clause   string
	Values   []interface{}
}

var (
	IllegalChar = `\'"/#+*%^~`
)

func NewSingleSqlCondition(key, operator string, value interface{}) SqlCondition {
	return NewSqlCondition(key, operator, []interface{}{value})
}

func NewSqlCondition(key string, operator string, values []interface{}) SqlCondition {
	return SqlCondition{Key: key, Operator: operator, Values: values}
}

func NewClauseSqlCondition(clause string, values []interface{}) SqlCondition {
	return SqlCondition{Clause: clause, Values: values}
}

func BuildSingleWhereClause(SqlCondition *SqlCondition) (string, []interface{}) {
	var questionMarkStr string

	if SqlCondition.Operator == "" {
		SqlCondition.Operator = "="
	} else if SqlCondition.Operator == "BETWEEN" {
		questionMarkStr = "? AND ?"
	} else {
		questionMarkStr = GenerateQuestionMarks(SqlCondition.Values)
		if SqlCondition.Operator == "IN" || SqlCondition.Operator == "NOT IN" || len(SqlCondition.Values) > 1 {
			questionMarkStr = "(" + questionMarkStr + ")"
		}
	}
	if strings.ContainsAny(SqlCondition.Key, IllegalChar) || strings.ContainsAny(SqlCondition.Operator, IllegalChar) {
		SqlCondition.Clause = "( 1 = 0 )"
		SqlCondition.Values = []interface{}{}
		return SqlCondition.Clause, SqlCondition.Values
	}
	SqlCondition.Clause = fmt.Sprintf("(%s %s %s)", SqlCondition.Key, SqlCondition.Operator, questionMarkStr)
	return SqlCondition.Clause, SqlCondition.Values
}

func BuildWhereClause(SqlConditions []SqlCondition) (string, []interface{}) {
	if len(SqlConditions) == 0 {
		return "", []interface{}{}
	}
	var allValues = make([]interface{}, 0)
	var allQueries = make([]string, 0)
	for _, condition := range SqlConditions {
		if condition.Clause == "" {
			BuildSingleWhereClause(&condition)

		}
		allQueries = append(allQueries, condition.Clause)
		for _, v := range condition.Values {
			allValues = append(allValues, v)
		}
	}
	whereClause := strings.Join(allQueries, " AND ")
	return whereClause, allValues
}

func GenerateQuestionMarks(columns []interface{}) string {
	var questionMarks = make([]string, 0)
	for range columns {
		questionMarks = append(questionMarks, "?")
	}
	return strings.Join(questionMarks, ", ")
}
