package model

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms7 struct {

	// Cash amount resulting from a foreign exchange trade.
	ToAmount *ActiveCurrencyAnd13DecimalAmount `xml:"ToAmt,omitempty"`

	// Cash amount for which a foreign exchange is required.
	FromAmount *ActiveCurrencyAndAmount `xml:"FrAmt,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveOrHistoricCurrencyCode `xml:"QtdCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a  foreign exchange rate.
	QuotingInstitution *PartyIdentification2Choice `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms7) SetToAmount(value, currency string) {
	f.ToAmount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (f *ForeignExchangeTerms7) SetFromAmount(value, currency string) {
	f.FromAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *ForeignExchangeTerms7) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms7) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms7) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms7) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms7) AddQuotingInstitution() *PartyIdentification2Choice {
	f.QuotingInstitution = new(PartyIdentification2Choice)
	return f.QuotingInstitution
}
