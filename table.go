package grel

import (
	"fmt"
	"strconv"
)

type Table struct {
	Expression

	Name  string
	Alias string
}

func NewTable(name string) Table {
	return Table{Name: name}
}

func NewTableWithAlias(name, alias string) Table {
	return Table{Name: name, Alias: alias}
}

func (t Table) SQL() string {
	sql := quoteTable(t.Name)
	if t.Alias != "" {
		sql += " AS " + quoteTable(t.Alias)
	}

	return sql
}

func quoteTable(s string) string {
	return strconv.Quote(s)
}

func tableCacheKey(t Table) string {
	if t.Alias != "" {
		return fmt.Sprintf("%s:%s", t.Name, t.Alias)
	} else {
		return t.Name
	}
}
