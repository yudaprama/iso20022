package model

// Specifies the process of a currency exchange or conversion.
type CurrencyReference2 struct {

	// Currency into which an amount is to be converted in a currency conversion.
	TargetCurrency *CurrencyCode `xml:"TrgtCcy"`

	// Currency of the amount to be converted in a currency conversion.
	SourceCurrency *CurrencyCode `xml:"SrcCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRateInformation []*ExchangeRateInformation1 `xml:"XchgRateInf,omitempty"`
}

func (c *CurrencyReference2) SetTargetCurrency(value string) {
	c.TargetCurrency = (*CurrencyCode)(&value)
}

func (c *CurrencyReference2) SetSourceCurrency(value string) {
	c.SourceCurrency = (*CurrencyCode)(&value)
}

func (c *CurrencyReference2) AddExchangeRateInformation() *ExchangeRateInformation1 {
	newValue := new(ExchangeRateInformation1)
	c.ExchangeRateInformation = append(c.ExchangeRateInformation, newValue)
	return newValue
}
