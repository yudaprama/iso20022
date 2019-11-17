package model

// Choice between a unit in decimal number, a face amount in currency and amount or a unit expressed as a code.
type UnitOrFaceAmountOrCode1Choice struct {

	// Quantity expressed as a number, eg, a number of shares.
	Unit *DecimalNumber `xml:"Unit"`

	// Quantity expressed as an amount representing the face amount, ie, the principal, of a debt instrument.
	FaceAmount *ActiveCurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as a code.
	Code *Quantity1Code `xml:"Cd"`
}

func (u *UnitOrFaceAmountOrCode1Choice) SetUnit(value string) {
	u.Unit = (*DecimalNumber)(&value)
}

func (u *UnitOrFaceAmountOrCode1Choice) SetFaceAmount(value, currency string) {
	u.FaceAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (u *UnitOrFaceAmountOrCode1Choice) SetCode(value string) {
	u.Code = (*Quantity1Code)(&value)
}
