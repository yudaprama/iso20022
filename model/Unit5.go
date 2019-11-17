package model

// Quantity expressed as a number and its details.
type Unit5 struct {

	// Quantity expressed as a number, for example, a number of shares.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units"`
}

func (u *Unit5) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	u.UnitsNumber = new(FinancialInstrumentQuantity1)
	return u.UnitsNumber
}

func (u *Unit5) SetGroup1Or2Units(value string) {
	u.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}
