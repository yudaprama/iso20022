package model

// Exchange rate and calculated amount to be presented to the customer when the dispense currency or the deposit currency (target currency) is different to account currency (source currency).
type CurrencyConversion5 struct {

	// Currency from which the amount is converted (ISO 4217, 3 alphanumeric characters).
	SourceCurrency *ActiveCurrencyCode `xml:"SrcCcy"`

	// Currency from which the amount is converted (ISO 4217, 3 numeric characters).
	SourceCurrencyNumeric *ActiveCurrencyCode `xml:"SrcCcyNmrc"`

	// Currency into which the amount is converted (ISO 4217, 3 alphanumeric characters).
	TargetCurrency *ActiveCurrencyCode `xml:"TrgtCcy"`

	// Currency into which the amount is converted (ISO 4217, 3 numeric characters).
	TargetCurrencyNumeric *Exact3NumericText `xml:"TrgtCcyNmrc"`

	// Currency exchange rate.
	Rate *BaseOneRate `xml:"Rate"`

	// Resulting calculated amount is in target currency.
	CalculatedAmount *ImpliedCurrencyAndAmount `xml:"ClctdAmt"`
}

func (c *CurrencyConversion5) SetSourceCurrency(value string) {
	c.SourceCurrency = (*ActiveCurrencyCode)(&value)
}

func (c *CurrencyConversion5) SetSourceCurrencyNumeric(value string) {
	c.SourceCurrencyNumeric = (*ActiveCurrencyCode)(&value)
}

func (c *CurrencyConversion5) SetTargetCurrency(value string) {
	c.TargetCurrency = (*ActiveCurrencyCode)(&value)
}

func (c *CurrencyConversion5) SetTargetCurrencyNumeric(value string) {
	c.TargetCurrencyNumeric = (*Exact3NumericText)(&value)
}

func (c *CurrencyConversion5) SetRate(value string) {
	c.Rate = (*BaseOneRate)(&value)
}

func (c *CurrencyConversion5) SetCalculatedAmount(value, currency string) {
	c.CalculatedAmount = NewImpliedCurrencyAndAmount(value, currency)
}
