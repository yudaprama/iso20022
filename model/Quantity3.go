package model

// Specifies the quantity of a product in a trade transaction.
type Quantity3 struct {

	// Specifies the unit of measurement. For example, kilo, tons.
	UnitOfMeasureCode *UnitOfMeasure4Code `xml:"UnitOfMeasrCd"`

	// Identifies the unit of measure not present in the code list.
	OtherUnitOfMeasure *Max35Text `xml:"OthrUnitOfMeasr"`

	// Quantity of a product on a line specified by a number. For example, 100 (kgs), 50 (pieces).
	Value *DecimalNumber `xml:"Val"`
}

func (q *Quantity3) SetUnitOfMeasureCode(value string) {
	q.UnitOfMeasureCode = (*UnitOfMeasure4Code)(&value)
}

func (q *Quantity3) SetOtherUnitOfMeasure(value string) {
	q.OtherUnitOfMeasure = (*Max35Text)(&value)
}

func (q *Quantity3) SetValue(value string) {
	q.Value = (*DecimalNumber)(&value)
}
