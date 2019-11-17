package model

// Choice between a unit in decimal number, a face amount in currency and amount or a unit expressed as a code.
type UnitOrFaceAmountOrCodeChoice struct {

	// Quantity expressed as a number, eg, a number of shares.
	Unit *DecimalNumber `xml:"Unit"`

	// Quantity expressed as an amount representing the face amount, ie, the principal, of a debt instrument.
	FaceAmount *CurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as a code.
	Code *Quantity1Code `xml:"Cd"`
}

func (u *UnitOrFaceAmountOrCodeChoice) SetUnit(value string) {
	u.Unit = (*DecimalNumber)(&value)
}

func (u *UnitOrFaceAmountOrCodeChoice) SetFaceAmount(value, currency string) {
	u.FaceAmount = NewCurrencyAndAmount(value, currency)
}

func (u *UnitOrFaceAmountOrCodeChoice) SetCode(value string) {
	u.Code = (*Quantity1Code)(&value)
}
