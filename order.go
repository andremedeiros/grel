package grel

// OrderDirection is an enum type for ordering.
type OrderDirection string

const (
	// Ascending is the ascending order.
	Ascending OrderDirection = "ASC"

	// Descending is the descending order.
	Descending OrderDirection = "DESC"
)

// Order is a struct that represents an order expression.
type Order struct {
	Expression

	Column    Expression
	Direction OrderDirection
}

// NewOrder creates a new Order expression.
func NewOrder(column Expression, direction OrderDirection) Order {
	return Order{Column: column, Direction: direction}
}

// SQL returns the SQL representation of the order expression.
func (o Order) SQL() string {
	return o.Column.SQL() + " " + string(o.Direction)
}
