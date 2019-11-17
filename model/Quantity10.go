package model

// Specifies the quantity of a product in a trade transaction.
type Quantity10 struct {

	// Specifies a unit of measure with a code or free text.
	UnitOfMeasure *UnitOfMeasure3Choice `xml:"UnitOfMeasr"`

	// Quantity of a product on a line specified by a number. For example, 100 (kgs), 50 (pieces).
	Value *DecimalNumber `xml:"Val"`
}

func (q *Quantity10) AddUnitOfMeasure() *UnitOfMeasure3Choice {
	q.UnitOfMeasure = new(UnitOfMeasure3Choice)
	return q.UnitOfMeasure
}

func (q *Quantity10) SetValue(value string) {
	q.Value = (*DecimalNumber)(&value)
}
