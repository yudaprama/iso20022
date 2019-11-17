package model

// Choice between a unit in decimal number or a face amount in currency and amount.
type UnitOrFaceAmount1Choice struct {

	// Quantity expressed as a number, eg, a number of shares.
	Unit *DecimalNumber `xml:"Unit"`

	// Quantity expressed as an amount representing the face amount, ie, the principal, of a debt instrument.
	FaceAmount *ActiveCurrencyAndAmount `xml:"FaceAmt"`
}

func (u *UnitOrFaceAmount1Choice) SetUnit(value string) {
	u.Unit = (*DecimalNumber)(&value)
}

func (u *UnitOrFaceAmount1Choice) SetFaceAmount(value, currency string) {
	u.FaceAmount = NewActiveCurrencyAndAmount(value, currency)
}
