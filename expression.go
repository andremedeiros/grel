package grel

// Expression is the interface that wraps the basic SQL method.
type Expression interface {
	SQL() string
}
