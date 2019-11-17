package model

// Choice between formats for the quantity of security.
type FinancialInstrumentQuantity15Choice struct {

	// Quantity expressed as a number, eg, a number of shares.
	Unit *RestrictedFINDecimalNumber `xml:"Unit"`

	// Quantity expressed as an amount representing the face amount, ie, the principal, of a debt instrument.
	FaceAmount *RestrictedFINImpliedCurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as an amount representing the current amortised face amount of a bond, for example, a periodic reduction/increase of a bond's principal amount.
	AmortisedValue *RestrictedFINImpliedCurrencyAndAmount `xml:"AmtsdVal"`
}

func (f *FinancialInstrumentQuantity15Choice) SetUnit(value string) {
	f.Unit = (*RestrictedFINDecimalNumber)(&value)
}

func (f *FinancialInstrumentQuantity15Choice) SetFaceAmount(value, currency string) {
	f.FaceAmount = NewRestrictedFINImpliedCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity15Choice) SetAmortisedValue(value, currency string) {
	f.AmortisedValue = NewRestrictedFINImpliedCurrencyAndAmount(value, currency)
}
