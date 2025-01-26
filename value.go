package grel

import "fmt"

type Value struct {
	Expression

	Value       interface{}
	Placeholder int
}

func NewValue(value interface{}) Value {
	return Value{Value: value}
}

func (v Value) SQL() string {
	return fmt.Sprintf(`$%d`, v.Placeholder)
}
