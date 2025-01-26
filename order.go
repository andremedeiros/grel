package grel

type OrderDirection string

const (
	Ascending  OrderDirection = "ASC"
	Descending OrderDirection = "DESC"
)

type Order struct {
	Expression

	Column    Expression
	Direction OrderDirection
}

func NewOrder(column Expression, direction OrderDirection) Order {
	return Order{Column: column, Direction: direction}
}

func (o Order) SQL() string {
	return o.Column.SQL() + " " + string(o.Direction)
}
