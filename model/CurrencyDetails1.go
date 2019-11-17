package model

// Details of a currency.
type CurrencyDetails1 struct {

	// Alpha currency code (ISO 4217, 3 alphanumeric characters).
	AlphaCode *CurrencyCode `xml:"AlphaCd"`

	// Numeric currency code (ISO 4217, 3 numeric characters).
	NumericCode *Exact3NumericText `xml:"NmrcCd"`

	// Maximal number of digits after the decimal separator for the currency.
	Decimal *Number `xml:"Dcml"`

	// Full name of the currency.
	Name *Max35Text `xml:"Nm,omitempty"`
}

func (c *CurrencyDetails1) SetAlphaCode(value string) {
	c.AlphaCode = (*CurrencyCode)(&value)
}

func (c *CurrencyDetails1) SetNumericCode(value string) {
	c.NumericCode = (*Exact3NumericText)(&value)
}

func (c *CurrencyDetails1) SetDecimal(value string) {
	c.Decimal = (*Number)(&value)
}

func (c *CurrencyDetails1) SetName(value string) {
	c.Name = (*Max35Text)(&value)
}
