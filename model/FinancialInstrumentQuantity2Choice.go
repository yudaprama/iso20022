package model

// Choice between formats for the quantity of security.
type FinancialInstrumentQuantity2Choice struct {

	// Quantity expressed as a number, for example, a number of shares.
	Unit *DecimalNumber `xml:"Unit"`

	// Quantity expressed as an amount representing the face amount, that is, the principal, of a debt instrument.
	FaceAmount *ImpliedCurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as an amount representing the current amortised face amount of a bond, for example, a periodic reduction/increase of a bond's principal amount.
	AmortisedValue *ImpliedCurrencyAndAmount `xml:"AmtsdVal"`

	// Quantity expressed as a code.
	Code *Quantity2Code `xml:"Cd"`
}

func (f *FinancialInstrumentQuantity2Choice) SetUnit(value string) {
	f.Unit = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentQuantity2Choice) SetFaceAmount(value, currency string) {
	f.FaceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity2Choice) SetAmortisedValue(value, currency string) {
	f.AmortisedValue = NewImpliedCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity2Choice) SetCode(value string) {
	f.Code = (*Quantity2Code)(&value)
}
