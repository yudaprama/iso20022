package model

// Quantity expressed as a number and its details.
type Unit4 struct {

	// Quantity expressed as a number, for example, a number of shares.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitDetails []*Unit5 `xml:"UnitDtls,omitempty"`
}

func (u *Unit4) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	u.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return u.TotalUnitsNumber
}

func (u *Unit4) AddUnitDetails() *Unit5 {
	newValue := new(Unit5)
	u.UnitDetails = append(u.UnitDetails, newValue)
	return newValue
}
