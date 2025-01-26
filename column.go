package grel

import (
	"fmt"
	"strconv"
	"strings"
)

// Column represents a column in a table.
type Column struct {
	Expression

	Table Table
	Name  string
	Alias string
}

// NewColumn creates a new column.
func NewColumn(name string) Column {
	// Check if name is a table.column reference.
	// If it is, split it into table and column.
	// If it's not, just use the name as the column name.
	table, column := splitTableColumn(name)

	if table != "" {
		return Column{Table: NewTable(table), Name: column}
	}

	return Column{Name: column}
}

// NewColumnWithAlias creates a new column with an alias.
func NewColumnWithAlias(name, alias string) Column {
	return Column{Name: name, Alias: alias}
}

// SQL returns the SQL representation of the column.
func (c Column) SQL() string {
	if c.Table.Name != "" {
		return fmt.Sprintf("%s.%s", c.Table.SQL(), quoteColumn(c.Name))
	}

	return quoteColumn(c.Name)
}

func quoteColumn(s string) string {
	return strconv.Quote(s)
}

func columnCacheKey(c Column) string {
	if c.Alias != "" {
		return fmt.Sprintf("%s:%s:%s", tableCacheKey(c.Table), c.Name, c.Alias)
	} else {
		return fmt.Sprintf("%s:%s", tableCacheKey(c.Table), c.Name)
	}
}

func splitTableColumn(name string) (string, string) {
	parts := strings.SplitN(name, ".", 2) //nolint:gomnd
	if len(parts) == 1 {
		return "", parts[0]
	}

	return parts[0], parts[1]
}
