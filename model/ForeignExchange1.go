package model

// Provides the details of the foreign exchange.
type ForeignExchange1 struct {

	// Currency into which an amount is to be converted in a currency conversion.
	ForeignCurrency *ActiveOrHistoricCurrencyCode `xml:"FrgnCcy"`

	// Foreign exchange rate between the source and the foreign currency applicable to the first leg of the FX swap transaction. The foreign exchange spot rate will be reported as the number of foreign currency units per one unit of the source currency.
	ExchangeSpotRate *BaseOneRate `xml:"XchgSpotRate"`

	// Difference between the foreign exchange spot rate and the foreign exchange forward rate expressed in basis points quoted in accordance with the prevailing market conventions for the currency pair.
	// Usage:
	// This value can be either positive or negative.
	ExchangeForwardPoint *DecimalNumber `xml:"XchgFwdPt"`
}

func (f *ForeignExchange1) SetForeignCurrency(value string) {
	f.ForeignCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchange1) SetExchangeSpotRate(value string) {
	f.ExchangeSpotRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchange1) SetExchangeForwardPoint(value string) {
	f.ExchangeForwardPoint = (*DecimalNumber)(&value)
}
