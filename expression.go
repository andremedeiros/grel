package grel

type Expression interface {
	SQL() string
}
