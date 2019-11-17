package model

// Specifies the process of a currency exchange or conversion.
type CurrencyReference3 struct {

	// Currency into which an amount is to be converted in a currency conversion.
	TargetCurrency *ActiveCurrencyCode `xml:"TrgtCcy"`

	// Currency of the amount to be converted in a currency conversion.
	SourceCurrency *ActiveCurrencyCode `xml:"SrcCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRateInformation []*ExchangeRateInformation1 `xml:"XchgRateInf,omitempty"`
}

func (c *CurrencyReference3) SetTargetCurrency(value string) {
	c.TargetCurrency = (*ActiveCurrencyCode)(&value)
}

func (c *CurrencyReference3) SetSourceCurrency(value string) {
	c.SourceCurrency = (*ActiveCurrencyCode)(&value)
}

func (c *CurrencyReference3) AddExchangeRateInformation() *ExchangeRateInformation1 {
	newValue := new(ExchangeRateInformation1)
	c.ExchangeRateInformation = append(c.ExchangeRateInformation, newValue)
	return newValue
}
