package grel

import "fmt"

type JoinType string

const (
	InnerJoin JoinType = "INNER JOIN"
	LeftJoin  JoinType = "LEFT JOIN"
	OuterJoin JoinType = "OUTER JOIN"
)

type Join struct {
	Expression

	Table     Expression
	Type      JoinType
	Predicate Predicate
}

func NewJoin(table Expression, joinType JoinType, on Predicate) Join {
	return Join{Table: table, Type: joinType, Predicate: on}
}

func (j Join) SQL() string {
	return fmt.Sprintf("%s %s ON %s", string(j.Type), j.Table.SQL(), j.Predicate.SQL())
}
