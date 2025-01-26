package grel

import "fmt"

type Predicate interface {
	Expression
}

type BinaryPredicateType string

const (
	AndOperator BinaryPredicateType = "AND"
	OrOperator  BinaryPredicateType = "OR"
	EqOperator  BinaryPredicateType = "="
	NeqOperator BinaryPredicateType = "!="
	GtOperator  BinaryPredicateType = ">"
	GoeOperator BinaryPredicateType = ">="
	LtOperator  BinaryPredicateType = "<"
	LoeOperator BinaryPredicateType = "<="
)

type BinaryPredicate struct {
	Predicate

	Operator BinaryPredicateType
	Left     Expression
	Right    Expression
}

func (bo BinaryPredicate) SQL() string {
	switch bo.Operator {
	case AndOperator, OrOperator:
		return fmt.Sprintf(`(%s %s %s)`, bo.Left.SQL(), bo.Operator, bo.Right.SQL())
	case EqOperator, NeqOperator, GtOperator, GoeOperator, LtOperator, LoeOperator:
		return fmt.Sprintf(`%s %s %s`, bo.Left.SQL(), bo.Operator, bo.Right.SQL())
	}

	panic("unknown operator")
}

func Eq(left, right Expression) BinaryPredicate {
	return BinaryPredicate{Operator: EqOperator, Left: left, Right: right}
}

func Neq(left, right Expression) BinaryPredicate {
	return BinaryPredicate{Operator: NeqOperator, Left: left, Right: right}
}

func Gt(left, right Expression) BinaryPredicate {
	return BinaryPredicate{Operator: GtOperator, Left: left, Right: right}
}

func Goe(left, right Expression) BinaryPredicate {
	return BinaryPredicate{Operator: GoeOperator, Left: left, Right: right}
}

func Lt(left, right Expression) BinaryPredicate {
	return BinaryPredicate{Operator: LtOperator, Left: left, Right: right}
}

func Loe(left, right Expression) BinaryPredicate {
	return BinaryPredicate{Operator: LoeOperator, Left: left, Right: right}
}

func And(left, right Predicate) BinaryPredicate {
	return BinaryPredicate{Operator: AndOperator, Left: left, Right: right}
}

func Or(left, right Predicate) BinaryPredicate {
	return BinaryPredicate{Operator: OrOperator, Left: left, Right: right}
}
