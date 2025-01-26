package grel

import "fmt"

// Statement represents a SQL statement.
type Statement struct {
	Expression

	Table Table

	Parameters []interface{}
}

// SQL returns the SQL representation of the statement.
func NewStatement(table Table) Statement {
	return Statement{Table: table}
}

// SQL returns the SQL representation of the statement.
func (s *Statement) parameterizePredicate(predicate Predicate) Predicate {
	switch predicate.(type) {
	case BinaryPredicate:
		bp, _ := predicate.(BinaryPredicate)
		if lv, ok := bp.Left.(Value); ok {
			s.Parameters = append(s.Parameters, lv.Value)
			lv.Placeholder = len(s.Parameters)
			bp.Left = lv
		}

		if rv, ok := bp.Right.(Value); ok {
			s.Parameters = append(s.Parameters, rv.Value)
			rv.Placeholder = len(s.Parameters)
			bp.Right = rv
		}

		predicate = bp
	default:
		fmt.Printf("Unknown type: %T\n", predicate)
	}

	return predicate
}
