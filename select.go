package grel

// Select is a struct that represents a SELECT statement.
type Select struct {
	Statement

	Columns   []Expression
	Joins     []Expression
	Order     []Expression
	Predicate Expression

	columnCache map[string]Column
}

// NewSelect creates a new SELECT statement.
func NewSelect(table Table, columns ...Column) Select {
	s := Select{
		Statement: NewStatement(table),

		Columns: []Expression{},
		Joins:   []Expression{},
		Order:   []Expression{},

		columnCache: map[string]Column{},
	}
	for _, column := range columns {
		s = s.Column(column)
	}

	return s
}

// Column adds a column to the SELECT statement. It will automatically assign columns to the main table if they don't
// have a table reference.
func (s Select) Column(column Column) Select {
	if column.Table.Name == "" {
		column.Table = s.Table
	}

	cck := columnCacheKey(column)

	if _, ok := s.columnCache[cck]; !ok {
		s.Columns = append(s.Columns, column)
		s.columnCache[cck] = column
	}

	return s
}

// Join adds a JOIN clause to the SELECT statement.
func (s Select) Join(table Expression, joinType JoinType, on Expression) Select {
	on = s.parameterizePredicate(on)
	j := NewJoin(table, joinType, on)
	s.Joins = append(s.Joins, j)

	return s
}

// OrderBy adds an ORDER BY clause to the SELECT statement.
func (s Select) OrderBy(order ...Order) Select {
	for _, o := range order {
		s.Order = append(s.Order, o)
	}

	return s
}

// Where adds a WHERE clause to the SELECT statement.
func (s Select) Where(where Predicate) Select {
	s.Predicate = s.parameterizePredicate(where)

	return s
}

// SQL returns the SQL representation of the SELECT statement.
func (s Select) SQL() string {
	sql := "SELECT "

	if len(s.Columns) == 0 {
		sql += "*"
	} else {
		for i, column := range s.Columns {
			if i > 0 {
				sql += ", "
			}

			sql += column.SQL()
		}
	}

	sql += " FROM " + s.Table.SQL()

	if len(s.Joins) > 0 {
		for _, join := range s.Joins {
			sql += " " + join.SQL()
		}
	}

	if s.Predicate != nil {
		sql += " WHERE " + s.Predicate.SQL()
	}

	if len(s.Order) > 0 {
		sql += " ORDER BY "

		for i, order := range s.Order {
			if i > 0 {
				sql += ", "
			}

			sql += order.SQL()
		}
	}

	return sql
}
