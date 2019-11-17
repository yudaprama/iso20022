package model

// Set of elements used to provide details of the currency exchange.
type CurrencyExchange5 struct {

	// Currency from which an amount is to be converted in a currency conversion.
	SourceCurrency *ActiveOrHistoricCurrencyCode `xml:"SrcCcy"`

	// Currency into which an amount is to be converted in a currency conversion.
	TargetCurrency *ActiveOrHistoricCurrencyCode `xml:"TrgtCcy,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	//
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Unique identification to unambiguously identify the foreign exchange contract.
	ContractIdentification *Max35Text `xml:"CtrctId,omitempty"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`
}

func (c *CurrencyExchange5) SetSourceCurrency(value string) {
	c.SourceCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CurrencyExchange5) SetTargetCurrency(value string) {
	c.TargetCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CurrencyExchange5) SetUnitCurrency(value string) {
	c.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CurrencyExchange5) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CurrencyExchange5) SetContractIdentification(value string) {
	c.ContractIdentification = (*Max35Text)(&value)
}

func (c *CurrencyExchange5) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}
