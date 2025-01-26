package grel

import (
	"fmt"
	"strconv"
)

// Table represents a table in a database.
type Table struct {
	Expression

	Name  string
	Alias string
}

// NewTable creates a new table.
func NewTable(name string) Table {
	return Table{Name: name}
}

// NewTableWithAlias creates a new table with an alias.
func NewTableWithAlias(name, alias string) Table {
	return Table{Name: name, Alias: alias}
}

// SQL returns the SQL representation of the table.
func (t Table) SQL() string {
	if t.Alias == "" {
		return quoteTable(t.Name)
	}

	return fmt.Sprintf("%s AS %s", quoteTable(t.Name), quoteTable(t.Alias))
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
