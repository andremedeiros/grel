package grel

import "fmt"

// JoinType represents the type of join to be performed.
type JoinType string

const (
	// InnerJoin represents an inner join.
	InnerJoin JoinType = "INNER JOIN"

	// LeftJoin represents a left join.
	LeftJoin JoinType = "LEFT JOIN"

	// RightJoin represents a right join.
	RightJoin JoinType = "RIGHT JOIN"

	// OuterJoin represents an outer join.
	OuterJoin JoinType = "OUTER JOIN"
)

// Join represents a join operation between two tables.
type Join struct {
	Expression

	Table     Expression
	Type      JoinType
	Predicate Predicate
}

// NewJoin creates a new join expression.
func NewJoin(table Expression, joinType JoinType, on Predicate) Join {
	return Join{Table: table, Type: joinType, Predicate: on}
}

// SQL returns the SQL representation of the join expression.
func (j Join) SQL() string {
	return fmt.Sprintf("%s %s ON %s", string(j.Type), j.Table.SQL(), j.Predicate.SQL())
}
