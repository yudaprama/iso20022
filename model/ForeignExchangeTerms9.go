package model

// Provides information about the terms of the foreign exchange transaction.
type ForeignExchangeTerms9 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	//
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Amount of money resulting from a foreign exchange transaction.
	ResultingAmount *ActiveCurrencyAndAmount `xml:"RsltgAmt"`

	// Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`
}

func (f *ForeignExchangeTerms9) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms9) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms9) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms9) SetResultingAmount(value, currency string) {
	f.ResultingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *ForeignExchangeTerms9) SetOriginalAmount(value, currency string) {
	f.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}
