package model

// Specifies the quantity of a product in a trade transaction.
type Quantity9 struct {

	// Specifies a unit of measure with a code or free text.
	UnitOfMeasure *UnitOfMeasure3Choice `xml:"UnitOfMeasr"`

	// Quantity of a product on a line specified by a number. For example, 100 (kgs), 50 (pieces).
	Value *DecimalNumber `xml:"Val"`

	// Multiplication factor of measurement values. For example: goods that can be ordered by 36 pieces.
	Factor *Max15NumericText `xml:"Fctr,omitempty"`
}

func (q *Quantity9) AddUnitOfMeasure() *UnitOfMeasure3Choice {
	q.UnitOfMeasure = new(UnitOfMeasure3Choice)
	return q.UnitOfMeasure
}

func (q *Quantity9) SetValue(value string) {
	q.Value = (*DecimalNumber)(&value)
}

func (q *Quantity9) SetFactor(value string) {
	q.Factor = (*Max15NumericText)(&value)
}
