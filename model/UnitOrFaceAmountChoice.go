package model

// Choice between a unit in decimal number or a face amount in currency and amount.
type UnitOrFaceAmountChoice struct {

	// Quantity expressed as a number, eg, a number of shares.
	Unit *DecimalNumber `xml:"Unit"`

	// Quantity expressed as an amount representing the face amount, ie, the principal, of a debt instrument.
	FaceAmount *CurrencyAndAmount `xml:"FaceAmt"`
}

func (u *UnitOrFaceAmountChoice) SetUnit(value string) {
	u.Unit = (*DecimalNumber)(&value)
}

func (u *UnitOrFaceAmountChoice) SetFaceAmount(value, currency string) {
	u.FaceAmount = NewCurrencyAndAmount(value, currency)
}
