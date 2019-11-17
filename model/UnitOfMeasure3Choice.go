package model

// Specifies a unit of measure with a code or free text.
type UnitOfMeasure3Choice struct {

	// Quantity of a product on a line specified by a number. For example, 100 (kgs), 50 (pieces).
	UnitOfMeasureCode *UnitOfMeasure4Code `xml:"UnitOfMeasrCd"`

	// Identifies the unit of measure not present in the code list.
	OtherUnitOfMeasure *Max35Text `xml:"OthrUnitOfMeasr"`
}

func (u *UnitOfMeasure3Choice) SetUnitOfMeasureCode(value string) {
	u.UnitOfMeasureCode = (*UnitOfMeasure4Code)(&value)
}

func (u *UnitOfMeasure3Choice) SetOtherUnitOfMeasure(value string) {
	u.OtherUnitOfMeasure = (*Max35Text)(&value)
}
