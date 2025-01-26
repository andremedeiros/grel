package grel

import "fmt"

// Value represents a value in a SQL statement. It is also used to track placeholder values.
type Value struct {
	Expression

	Value       interface{}
	Placeholder int
}

// NewValue creates a new value.
func NewValue(value interface{}) Value {
	return Value{Value: value}
}

// SQL returns the SQL representation of the value.
func (v Value) SQL() string {
	return fmt.Sprintf(`$%d`, v.Placeholder)
}
