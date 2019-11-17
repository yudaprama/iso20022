package model

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms4 struct {

	// Currency and amount bought in a foreign exchange trade. The buy amount is received by the buyer.
	BuyAmount *ActiveCurrencyAnd13DecimalAmount `xml:"BuyAmt,omitempty"`

	// Currency and amount sold in a foreign exchange trade. The sold amount is delivered by the buyer.
	SellAmount *ActiveCurrencyAndAmount `xml:"SellAmt,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *CurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *CurrencyCode `xml:"QtdCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a  foreign exchange rate.
	QuotingInstitution *PartyIdentification2Choice `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms4) SetBuyAmount(value, currency string) {
	f.BuyAmount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (f *ForeignExchangeTerms4) SetSellAmount(value, currency string) {
	f.SellAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *ForeignExchangeTerms4) SetUnitCurrency(value string) {
	f.UnitCurrency = (*CurrencyCode)(&value)
}

func (f *ForeignExchangeTerms4) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*CurrencyCode)(&value)
}

func (f *ForeignExchangeTerms4) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms4) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms4) AddQuotingInstitution() *PartyIdentification2Choice {
	f.QuotingInstitution = new(PartyIdentification2Choice)
	return f.QuotingInstitution
}
