package model

// Details of a currency.
type CurrencyDetails2 struct {

	// Alpha currency code (ISO 4217, 3 alphanumeric characters).
	AlphaCode *ActiveCurrencyCode `xml:"AlphaCd,omitempty"`

	// Numeric currency code (ISO 4217, 3 numeric characters).
	NumericCode *Exact3NumericText `xml:"NmrcCd,omitempty"`

	// Maximal number of digits after the decimal separator for the currency.
	Decimal *Number `xml:"Dcml,omitempty"`

	// Full name of the currency.
	Name *Max35Text `xml:"Nm,omitempty"`
}

func (c *CurrencyDetails2) SetAlphaCode(value string) {
	c.AlphaCode = (*ActiveCurrencyCode)(&value)
}

func (c *CurrencyDetails2) SetNumericCode(value string) {
	c.NumericCode = (*Exact3NumericText)(&value)
}

func (c *CurrencyDetails2) SetDecimal(value string) {
	c.Decimal = (*Number)(&value)
}

func (c *CurrencyDetails2) SetName(value string) {
	c.Name = (*Max35Text)(&value)
}
