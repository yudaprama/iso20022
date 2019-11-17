package model

// Further detailed information on the exchange rate that has been used in the payment transaction.
type ExchangeRate1 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty"`

	// The factor used for conversion of an amount from one currency to another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies the type used to complete the currency exchange.
	RateType *ExchangeRateType1Code `xml:"RateTp,omitempty"`

	// Unique and unambiguous reference to the foreign exchange contract agreed between the initiating party/creditor and the debtor agent.
	ContractIdentification *Max35Text `xml:"CtrctId,omitempty"`
}

func (e *ExchangeRate1) SetUnitCurrency(value string) {
	e.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (e *ExchangeRate1) SetExchangeRate(value string) {
	e.ExchangeRate = (*BaseOneRate)(&value)
}

func (e *ExchangeRate1) SetRateType(value string) {
	e.RateType = (*ExchangeRateType1Code)(&value)
}

func (e *ExchangeRate1) SetContractIdentification(value string) {
	e.ContractIdentification = (*Max35Text)(&value)
}
