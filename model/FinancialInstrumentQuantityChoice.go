package model

// Choice between formats for the quantity of security.
type FinancialInstrumentQuantityChoice struct {

	// Quantity expressed as a number, eg, a number of shares.
	Unit *DecimalNumber `xml:"Unit"`

	// Quantity expressed as an amount representing the face amount, ie, the principal, of a debt instrument.
	FaceAmount *ImpliedCurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as an amount representing the current amortised face amount of a bond, eg, a periodic reduction of a bond's principal amount.
	AmortisedValue *ImpliedCurrencyAndAmount `xml:"AmtsdVal"`
}

func (f *FinancialInstrumentQuantityChoice) SetUnit(value string) {
	f.Unit = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentQuantityChoice) SetFaceAmount(value, currency string) {
	f.FaceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantityChoice) SetAmortisedValue(value, currency string) {
	f.AmortisedValue = NewImpliedCurrencyAndAmount(value, currency)
}
