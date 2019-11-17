package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice9 struct {

	// Specifies the unit of measurement. For example, kilo, tons.
	UnitOfMeasureCode *UnitOfMeasure4Code `xml:"UnitOfMeasrCd"`

	// Identifies the unit of measure not present in the code list.
	OtherUnitOfMeasure *Max35Text `xml:"OthrUnitOfMeasr"`

	// Price expressed as a currency and value.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Multiplication factor of measurement values. For example: goods that can be ordered by 36 pieces.
	Factor *Max15NumericText `xml:"Fctr,omitempty"`
}

func (u *UnitPrice9) SetUnitOfMeasureCode(value string) {
	u.UnitOfMeasureCode = (*UnitOfMeasure4Code)(&value)
}

func (u *UnitPrice9) SetOtherUnitOfMeasure(value string) {
	u.OtherUnitOfMeasure = (*Max35Text)(&value)
}

func (u *UnitPrice9) SetAmount(value, currency string) {
	u.Amount = NewCurrencyAndAmount(value, currency)
}

func (u *UnitPrice9) SetFactor(value string) {
	u.Factor = (*Max15NumericText)(&value)
}
