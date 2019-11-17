package model

// Choice between formats for the quantity of security.
type FinancialInstrumentQuantity18Choice struct {

	// Quantity expressed as a number, eg, a number of shares.
	Unit *DecimalNumber `xml:"Unit"`

	// Quantity expressed as an amount representing the face amount, ie, the principal, of a debt instrument.
	FaceAmount *ImpliedCurrencyAndAmount `xml:"FaceAmt"`
}

func (f *FinancialInstrumentQuantity18Choice) SetUnit(value string) {
	f.Unit = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentQuantity18Choice) SetFaceAmount(value, currency string) {
	f.FaceAmount = NewImpliedCurrencyAndAmount(value, currency)
}
