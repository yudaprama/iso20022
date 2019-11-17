package model

// Information needed to process a currency exchange or conversion.
type CurrencyExchange3 struct {

	// Currency of the amount to be converted in a currency conversion.
	SourceCurrency *CurrencyCode `xml:"SrcCcy"`

	// Currency into which an amount is to be converted in a currency conversion.
	TargetCurrency *CurrencyCode `xml:"TrgtCcy,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *CurrencyCode `xml:"UnitCcy,omitempty"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	//
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Unique and unambiguous identifier of the foreign exchange contract.
	ContractIdentification *Max35Text `xml:"CtrctId,omitempty"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`
}

func (c *CurrencyExchange3) SetSourceCurrency(value string) {
	c.SourceCurrency = (*CurrencyCode)(&value)
}

func (c *CurrencyExchange3) SetTargetCurrency(value string) {
	c.TargetCurrency = (*CurrencyCode)(&value)
}

func (c *CurrencyExchange3) SetUnitCurrency(value string) {
	c.UnitCurrency = (*CurrencyCode)(&value)
}

func (c *CurrencyExchange3) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CurrencyExchange3) SetContractIdentification(value string) {
	c.ContractIdentification = (*Max35Text)(&value)
}

func (c *CurrencyExchange3) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}
