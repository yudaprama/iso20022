package model

// Quantity of a security.
type FinancialInstrumentQuantity1 struct {

	// Quantity expressed as a number, eg, a number of shares.
	Unit *DecimalNumber `xml:"Unit"`
}

func (f *FinancialInstrumentQuantity1) SetUnit(value string) {
	f.Unit = (*DecimalNumber)(&value)
}
