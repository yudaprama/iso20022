package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice18 struct {

	// Specifies the unit price.
	UnitPrice *UnitOfMeasure3Choice `xml:"UnitPric"`

	// Price expressed as a currency and value.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Multiplication factor of measurement values. For example: goods that can be ordered by 36 pieces.
	Factor *Max15NumericText `xml:"Fctr,omitempty"`
}

func (u *UnitPrice18) AddUnitPrice() *UnitOfMeasure3Choice {
	u.UnitPrice = new(UnitOfMeasure3Choice)
	return u.UnitPrice
}

func (u *UnitPrice18) SetAmount(value, currency string) {
	u.Amount = NewCurrencyAndAmount(value, currency)
}

func (u *UnitPrice18) SetFactor(value string) {
	u.Factor = (*Max15NumericText)(&value)
}
