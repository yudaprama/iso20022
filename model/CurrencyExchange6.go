package model

// Describes the details of the currency exchange.
type CurrencyExchange6 struct {

	// Currency from which an amount is to be converted in a currency conversion.
	SourceCurrency *ActiveOrHistoricCurrencyCode `xml:"SrcCcy"`

	// Currency into which an amount is to be converted in a currency conversion.
	TargetCurrency *ActiveOrHistoricCurrencyCode `xml:"TrgtCcy"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Unique identification to unambiguously identify the foreign exchange contract.
	Description *Max40Text `xml:"Desc,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty"`

	// Further information on the exchange rate.
	Comments *Max70Text `xml:"Cmnts,omitempty"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`
}

func (c *CurrencyExchange6) SetSourceCurrency(value string) {
	c.SourceCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CurrencyExchange6) SetTargetCurrency(value string) {
	c.TargetCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CurrencyExchange6) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CurrencyExchange6) SetDescription(value string) {
	c.Description = (*Max40Text)(&value)
}

func (c *CurrencyExchange6) SetUnitCurrency(value string) {
	c.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CurrencyExchange6) SetComments(value string) {
	c.Comments = (*Max70Text)(&value)
}

func (c *CurrencyExchange6) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}
